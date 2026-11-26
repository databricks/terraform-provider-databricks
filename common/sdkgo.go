package common

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

// SdkGoVolumes calls github.com/databricks/sdk-go/uc/volumes/v1. Each call sends the
// request that the classic workspace client it was built from sends.
type SdkGoVolumes struct {
	client *sdkvolumes.Client
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
	return &SdkGoVolumes{client: vc, bridge: b}, nil
}

// ReadByName reads the volume fullName with the request that the classic
// Volumes.ReadByName sends for it, so "#", "?" and "%" in the name act as they do there.
func (vc *SdkGoVolumes) ReadByName(ctx context.Context,
	fullName string) (*sdkvolumes.VolumeInfo, error) {
	path := classicVolumePath(fullName)
	// The classic client adds an empty query to GETs.
	classic := classicVolumeRequest(http.MethodGet, path, path+"?")
	classic.decode = decodeClassicVolumeInfo
	get := func(ctx context.Context,
		opts ...call.Option) (*sdkvolumes.VolumeInfo, error) {
		req := sdkvolumes.GetVolumeRequest{FullNameArg: &fullName}
		return vc.client.GetVolume(ctx, req, opts...)
	}
	return callSdkGo(ctx, vc.bridge, classic, get)
}

// Create creates a volume with the request that the classic Volumes.Create sends.
func (vc *SdkGoVolumes) Create(ctx context.Context,
	req sdkvolumes.CreateVolumeRequest) (*sdkvolumes.VolumeInfo, error) {
	path := "/api/2.1/unity-catalog/volumes"
	classic := classicVolumeRequest(http.MethodPost, path, path)
	classic.body = func(body []byte) ([]byte, error) {
		return classicVolumeBody(body)
	}
	classic.decode = decodeClassicVolumeInfo
	create := func(ctx context.Context,
		opts ...call.Option) (*sdkvolumes.VolumeInfo, error) {
		return vc.client.CreateVolume(ctx, req, opts...)
	}
	return callSdkGo(ctx, vc.bridge, classic, create)
}

// Update updates the volume req.FullNameArg with the request that the classic
// Volumes.Update sends for it. The classic body has no full_name_arg, which sdk-go
// serializes from the path.
func (vc *SdkGoVolumes) Update(ctx context.Context,
	req sdkvolumes.UpdateVolumeRequest) (*sdkvolumes.VolumeInfo, error) {
	var fullName string
	if req.FullNameArg != nil {
		fullName = *req.FullNameArg
	}
	path := classicVolumePath(fullName)
	classic := classicVolumeRequest(http.MethodPatch, path, path)
	classic.body = func(body []byte) ([]byte, error) {
		return classicVolumeBody(body, "full_name_arg")
	}
	classic.decode = decodeClassicVolumeInfo
	update := func(ctx context.Context,
		opts ...call.Option) (*sdkvolumes.VolumeInfo, error) {
		return vc.client.UpdateVolume(ctx, req, opts...)
	}
	return callSdkGo(ctx, vc.bridge, classic, update)
}

// DeleteByName deletes the volume fullName with the request that the classic
// Volumes.DeleteByName sends for it.
func (vc *SdkGoVolumes) DeleteByName(ctx context.Context, fullName string) error {
	path := classicVolumePath(fullName)
	// The classic client adds an empty query to DELETEs.
	classic := classicVolumeRequest(http.MethodDelete, path, path+"?")
	del := func(ctx context.Context,
		opts ...call.Option) (*sdkvolumes.DeleteVolumeResponse, error) {
		req := sdkvolumes.DeleteVolumeRequest{FullNameArg: &fullName}
		return vc.client.DeleteVolume(ctx, req, opts...)
	}
	_, err := callSdkGo(ctx, vc.bridge, classic, del)
	return err
}

// decodeClassicVolumeInfo decodes a success body as the classic Volumes calls that return
// a volume do.
func decodeClassicVolumeInfo(body []byte) error {
	var v catalog.VolumeInfo
	return json.Unmarshal(body, &v)
}

// classicVolumePath is the path of a classic Volumes call for fullName, which the classic
// client formats unescaped.
func classicVolumePath(fullName string) string {
	return "/api/2.1/unity-catalog/volumes/" + fullName
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

// callSdkGo runs one sdk-go method with the classic client's retries and rate limit.
// Every attempt is sent as classic, and a failure returns the classic client's error.
func callSdkGo[T any](ctx context.Context, b *sdkGoBridge, classic classicRequest,
	method func(context.Context, ...call.Option) (T, error)) (T, error) {
	var zero T
	classicURL, err := url.Parse(classic.url)
	if err != nil {
		// Like the classic client, the call fails after the rate limiter, without
		// authenticating, sending or retrying.
		if err := (classicLimiter{b.limiter}).Wait(ctx); err != nil {
			return zero, err
		}
		return zero, fmt.Errorf("failed creating new request: %w", err)
	}
	attempt := &sdkGoAttempt{classic: classic, classicURL: classicURL}
	ctx = context.WithValue(ctx, sdkGoAttemptKey{}, attempt)
	retrier := b.newRetrier(ctx, attempt)
	result, err := method(ctx,
		call.WithRetrier(func() ops.Retrier { return retrier }),
		call.WithLimiter(classicLimiter{b.limiter}))
	attempt.endTimeout()
	if err == nil {
		return result, nil
	}
	// Like the classic client, a cancellation during backoff returns the last error.
	cancelled := err == context.Canceled || err == context.DeadlineExceeded
	if retrier.lastErr != nil && cancelled {
		return zero, retrier.lastErr
	}
	return zero, b.classicError(ctx, err, attempt)
}

type sdkGoAttemptKey struct{}

// sdkGoAttempt carries state between the stages of one call: the classic request to
// send, the last exchange, which classic error parsing needs, and the timeout of the
// attempt being sent.
type sdkGoAttempt struct {
	classic    classicRequest
	classicURL *url.URL
	request    *http.Request
	// requestBody, response and responseBody are those of the last response, if any.
	requestBody  []byte
	response     *http.Response
	responseBody []byte
	timeout      *attemptTimeout
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
