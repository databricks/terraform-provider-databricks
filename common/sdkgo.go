package common

import (
	"context"
	"net/http"
	"sync"

	classicclient "github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/sdk-go/core/ops"
	"github.com/databricks/sdk-go/options/call"
	"github.com/databricks/sdk-go/options/client"
	sdkscim "github.com/databricks/sdk-go/scim/v1"
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

// SdkGoScim is a github.com/databricks/sdk-go/scim/v1 client that sends requests the
// way the classic client it was built from does.
type SdkGoScim struct {
	*sdkscim.Client
	// IsAccount is true when the account-level methods must be used.
	IsAccount bool
	bridge    *sdkGoBridge
}

// SdkGoScimClient builds a scim/v1 client for apiLevel. Call it on the client already
// resolved for the resource, for example by DatabricksClientForDualResource.
func (c *DatabricksClient) SdkGoScimClient(ctx context.Context,
	apiLevel string) (*SdkGoScim, error) {
	b, err := sdkGoBridgeFor(c.DatabricksClient)
	if err != nil {
		return nil, err
	}
	isAccount := isAccountLevelFromApiLevel(apiLevel, c)
	opts := []client.Option{
		client.WithHost(c.Config.Host),
		client.WithCredentials(classicCredentials{c.Config, b.httpTimeout}),
		client.WithHTTPClient(b.httpClient),
		// retry_timeout_seconds and http_timeout_seconds bound calls instead.
		client.WithTimeout(0),
		client.WithoutConfigFile(),
		client.WithoutEnv(),
	}
	if isAccount {
		opts = append(opts, client.WithAccountID(c.Config.AccountID))
	} else {
		opts = append(opts, client.WithWorkspaceID(c.Config.WorkspaceID))
	}
	sc, err := sdkscim.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &SdkGoScim{Client: sc, IsAccount: isAccount, bridge: b}, nil
}

// CallSdkGoScim runs one method of sc with the classic client's retries and rate limit,
// and returns the error the classic client would return for the same response.
func CallSdkGoScim[T any](ctx context.Context, sc *SdkGoScim,
	method func(context.Context, ...call.Option) (T, error)) (T, error) {
	return callSdkGo(ctx, sc.bridge, method)
}

func callSdkGo[T any](ctx context.Context, b *sdkGoBridge,
	method func(context.Context, ...call.Option) (T, error)) (T, error) {
	attempt := &sdkGoAttempt{}
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
	return zero, b.classicError(ctx, err, attempt.request)
}

type sdkGoAttemptKey struct{}

// sdkGoAttempt carries state between the stages of one call: the last request sent,
// which classic error parsing needs, and the timeout of the attempt being sent.
type sdkGoAttempt struct {
	request *http.Request
	timeout *attemptTimeout
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

// Bridges are kept per classic client, whose settings and rate limiter they use.
var sdkGoBridges sync.Map

func sdkGoBridgeFor(classic *classicclient.DatabricksClient) (*sdkGoBridge, error) {
	if b, ok := sdkGoBridges.Load(classic); ok {
		return b.(*sdkGoBridge), nil
	}
	b, err := newSdkGoBridge(classic)
	if err != nil {
		return nil, err
	}
	actual, _ := sdkGoBridges.LoadOrStore(classic, b)
	return actual.(*sdkGoBridge), nil
}
