package common

import (
	"context"
	"net/http"
	"sync"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/sdk-go/core/ops"
	"github.com/databricks/sdk-go/options/call"
	"github.com/databricks/sdk-go/options/client"
	sdkvolumes "github.com/databricks/sdk-go/uc/volumes/v1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type sdkGoContextKey struct{}

// UsesSdkGo reports whether the current operation belongs to a resource selected by
// USE_SDK_GO_RESOURCES.
func UsesSdkGo(ctx context.Context) bool {
	enabled, _ := ctx.Value(sdkGoContextKey{}).(bool)
	return enabled
}

// EnableSdkGo makes UsesSdkGo true for the operations and the importer of r.
func EnableSdkGo(r *schema.Resource) {
	r.CreateContext = withSdkGo(r.CreateContext)
	r.ReadContext = withSdkGo(r.ReadContext)
	r.UpdateContext = withSdkGo(r.UpdateContext)
	r.DeleteContext = withSdkGo(r.DeleteContext)
	if r.Importer != nil && r.Importer.StateContext != nil {
		importer := *r.Importer
		stateContext := importer.StateContext
		importer.StateContext = func(ctx context.Context, d *schema.ResourceData,
			m any) ([]*schema.ResourceData, error) {
			return stateContext(context.WithValue(ctx, sdkGoContextKey{}, true), d, m)
		}
		r.Importer = &importer
	}
}

type resourceOperation = func(context.Context, *schema.ResourceData, any) diag.Diagnostics

func withSdkGo[F ~resourceOperation](f F) F {
	if f == nil {
		return nil
	}
	return func(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
		return f(context.WithValue(ctx, sdkGoContextKey{}, true), d, m)
	}
}

// SdkGoVolumes is a github.com/databricks/sdk-go/uc/volumes/v1 client that sends
// requests the way the workspace client it was built from does.
type SdkGoVolumes struct {
	*sdkvolumes.Client
	bridge *sdkGoBridge
}

// SdkGoVolumesClient builds a uc/volumes/v1 client from w, the workspace client already
// resolved for the resource, for example by WorkspaceClientUnifiedProvider.
func SdkGoVolumesClient(ctx context.Context,
	w *databricks.WorkspaceClient) (*SdkGoVolumes, error) {
	api, err := workspaceApiClient(w)
	if err != nil {
		return nil, err
	}
	b, err := sdkGoBridgeFor(w.Config, api)
	if err != nil {
		return nil, err
	}
	opts := append(b.clientOptions(w.Config),
		client.WithWorkspaceID(w.Config.WorkspaceID))
	vc, err := sdkvolumes.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &SdkGoVolumes{Client: vc, bridge: b}, nil
}

// clientOptions returns the options that make an sdk-go client send requests to cfg's
// host through the bridge.
func (b *sdkGoBridge) clientOptions(cfg *config.Config) []client.Option {
	return []client.Option{
		client.WithHost(cfg.Host),
		client.WithCredentials(classicCredentials{cfg, b.httpTimeout}),
		client.WithHTTPClient(b.httpClient),
		// retry_timeout_seconds and http_timeout_seconds bound calls instead.
		client.WithTimeout(0),
		client.WithoutConfigFile(),
		client.WithoutEnv(),
	}
}

// CallSdkGoVolumes runs one method of vc with the classic client's retries, rate limit
// and response decoding, and returns the error the classic client would return for the
// same response.
func CallSdkGoVolumes[T any](ctx context.Context, vc *SdkGoVolumes,
	method func(context.Context, ...call.Option) (T, error)) (T, error) {
	return callSdkGo(ctx, vc.bridge, classicVolumesResponse[T](), method)
}

// classicVolumesResponse returns a constructor of the classic type that the classic
// client decodes a T response into, or nil if it decodes none, as for a delete.
func classicVolumesResponse[T any]() func() any {
	if _, ok := any(*new(T)).(*sdkvolumes.VolumeInfo); ok {
		return func() any { return &catalog.VolumeInfo{} }
	}
	return nil
}

func callSdkGo[T any](ctx context.Context, b *sdkGoBridge, classicResponse func() any,
	method func(context.Context, ...call.Option) (T, error)) (T, error) {
	attempt := &sdkGoAttempt{classicResponse: classicResponse}
	ctx = context.WithValue(ctx, sdkGoAttemptKey{}, attempt)
	retrier := b.newRetrier(ctx, attempt)
	result, err := method(ctx,
		call.WithRetrier(func() ops.Retrier { return retrier }),
		call.WithLimiter(classicLimiter{b.limiter}))
	attempt.endTimeout()
	if err == nil {
		return result, nil
	}
	var zero T
	// Like the classic client, a cancellation during backoff returns the last error.
	cancelled := err == context.Canceled || err == context.DeadlineExceeded
	if retrier.lastErr != nil && cancelled {
		return zero, retrier.lastErr
	}
	return zero, b.classicError(ctx, err, attempt)
}

type sdkGoAttemptKey struct{}

// sdkGoAttempt carries state between the stages of one call: the last request sent and
// its response, which classic error parsing needs, the timeout of the attempt being
// sent, and the classic type that successful responses are decoded into.
type sdkGoAttempt struct {
	request         *http.Request
	response        *http.Response
	timeout         *attemptTimeout
	classicResponse func() any
}

// succeeded reports whether the last attempt got a successful response.
func (a *sdkGoAttempt) succeeded() bool {
	return a != nil && a.response != nil && a.response.StatusCode < http.StatusBadRequest
}

// keepTimeout stores the timeout started by authentication for the transport.
func (a *sdkGoAttempt) keepTimeout(t *attemptTimeout) {
	a.endTimeout()
	a.timeout = t
}

func (a *sdkGoAttempt) takeTimeout() *attemptTimeout {
	t := a.timeout
	a.timeout = nil
	return t
}

// endTimeout ends a timeout that no request took, such as after a failed request
// visitor.
func (a *sdkGoAttempt) endTimeout() {
	if t := a.takeTimeout(); t != nil {
		t.end()
	}
}

// Bridges are kept per classic API client, whose rate limiter they share, with the
// settings of the config it was built from.
var sdkGoBridges sync.Map

func sdkGoBridgeFor(cfg *config.Config, api *httpclient.ApiClient) (*sdkGoBridge, error) {
	if b, ok := sdkGoBridges.Load(api); ok {
		return b.(*sdkGoBridge), nil
	}
	b, err := newSdkGoBridge(cfg, api)
	if err != nil {
		return nil, err
	}
	actual, _ := sdkGoBridges.LoadOrStore(api, b)
	return actual.(*sdkGoBridge), nil
}
