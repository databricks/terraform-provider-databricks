package sdkgoutility

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"slices"
	"sync"
	"testing"
	"testing/synctest"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/sdk-go/auth/credentials"
	"github.com/databricks/sdk-go/core/ops"
	modelregistry "github.com/databricks/sdk-go/modelregistry/v1"
	"github.com/databricks/sdk-go/options/call"
	"github.com/databricks/sdk-go/options/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The tests that measure a rate run in a synctest bubble, whose fake clock
// makes the time of every admission exact and the tests instant. The tests
// that compare the limiter with the classic client run a real classic client
// and a real modular client: in a bubble against an in-memory registry for the
// rate, and against an httptest.Server for the errors.

const (
	// getModelPath is the path of the call that both clients make in these
	// tests: GetModel of the classic client and GetRegisteredModelDatabricks
	// of the modular one.
	getModelPath = "/api/2.0/mlflow/databricks/registered-models/get"

	// unlimited is the rate of a limiter that admits every call at once.
	unlimited = 0
)

// wantAdmissions returns when a limiter that the classic client builds for
// perSecond admits n calls made one after the other, counted from the first:
// perSecond calls at once and then one call every 1/perSecond of a second.
func wantAdmissions(perSecond, n int) []time.Duration {
	times := make([]time.Duration, n)
	for i := perSecond; perSecond != unlimited && i < n; i++ {
		times[i] = time.Duration(i-perSecond+1) * time.Second / time.Duration(perSecond)
	}
	return times
}

// assertAdmissions checks that calls were admitted at the times want. It
// allows a microsecond, because the limiter computes its waits in floating
// point.
func assertAdmissions(t *testing.T, want, got []time.Duration) {
	t.Helper()
	require.Len(t, got, len(want))
	for i := range want {
		assert.InDelta(t, want[i], got[i], float64(time.Microsecond), "call %d", i+1)
	}
}

// admissions makes n calls one after the other and returns when each of them
// returned, counted from the first.
func admissions(t *testing.T, n int, makeCall func() error) []time.Duration {
	t.Helper()
	start := time.Now()
	times := make([]time.Duration, n)
	for i := range times {
		require.NoError(t, makeCall())
		times[i] = time.Since(start)
	}
	return times
}

// forgetNilConfig gives the nil config a new budget for the test. All tests
// of the nil config would otherwise draw from one.
func forgetNilConfig(t *testing.T) {
	t.Helper()
	var cfg *config.Config
	budgets.Delete(cfg)
	t.Cleanup(func() { budgets.Delete(cfg) })
}

// noopLoader keeps the DATABRICKS_ variables and the config file of the
// machine out of a classic config: the SDK loads them only when the config
// names no loader.
type noopLoader struct{}

func (noopLoader) Name() string                   { return "noop" }
func (noopLoader) Configure(*config.Config) error { return nil }

// registry answers the call that both clients make and notes when each call
// arrived. It is the transport of the clients in a synctest bubble, which
// can't wait on a real connection, and the handler of an httptest.Server.
type registry struct {
	mu       sync.Mutex
	arrivals []time.Time
}

func (r *registry) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet || req.URL.Path != getModelPath {
		http.NotFound(w, req)
		return
	}
	r.mu.Lock()
	r.arrivals = append(r.arrivals, time.Now())
	r.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"registered_model_databricks": {"name": "model"}}`))
}

func (r *registry) RoundTrip(req *http.Request) (*http.Response, error) {
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)
	resp := recorder.Result()
	resp.Request = req
	return resp, nil
}

// take returns when the calls since the last take arrived, counted from
// start. The clients take turns, so these are the calls of one client.
func (r *registry) take(start time.Time) []time.Duration {
	r.mu.Lock()
	defer r.mu.Unlock()
	times := make([]time.Duration, len(r.arrivals))
	for i, arrival := range r.arrivals {
		times[i] = arrival.Sub(start)
	}
	r.arrivals = nil
	return times
}

// fixtureRegistry is a registry that says whether the classic client retries
// on IO errors, as the HTTP fixtures of the classic SDK do. The classic client
// doesn't limit the rate of a transport that says it doesn't.
type fixtureRegistry struct {
	*registry
	skipRetryOnIO bool
}

func (f fixtureRegistry) SkipRetryOnIO() bool { return f.skipRetryOnIO }

// clients are a real classic client and a real modular client of one config.
// The modular client limits its calls with Limiter of that config.
type clients struct {
	classic *databricks.WorkspaceClient
	modular *modelregistry.Client
	limiter ops.Limiter
}

// newClients returns the clients of a config for host with the rate limit.
// A nil transport is the default transport of each client.
func newClients(t *testing.T, host string, perSecond int, transport http.RoundTripper) clients {
	t.Helper()
	cfg := &config.Config{
		Host:               host,
		Token:              "token",
		RateLimitPerSecond: perSecond,
		HTTPTransport:      transport,
		Loaders:            []config.Loader{noopLoader{}},
		// The registry has no metadata of the host for the classic client.
		HostMetadataResolver: func(context.Context, string) (*config.HostMetadata, error) {
			return nil, nil
		},
	}
	classic, err := databricks.NewWorkspaceClient((*databricks.Config)(cfg))
	require.NoError(t, err)
	creds, err := credentials.NewPATCredentials("token")
	require.NoError(t, err)
	modular, err := modelregistry.NewClient(t.Context(),
		client.WithHost(host),
		client.WithCredentials(creds),
		client.WithHTTPClient(&http.Client{Transport: transport}),
		client.WithoutEnv(),
		client.WithoutConfigFile(),
	)
	require.NoError(t, err)
	return clients{classic: classic, modular: modular, limiter: Limiter(cfg)}
}

func (c clients) classicCall(ctx context.Context) error {
	_, err := c.classic.ModelRegistry.GetModel(ctx, ml.GetModelRequest{Name: "model"})
	return err
}

func (c clients) modularCall(ctx context.Context) error {
	name := "model"
	_, err := c.modular.GetRegisteredModelDatabricks(ctx,
		modelregistry.GetRegisteredModelDatabricksRequest{Name: &name},
		call.WithLimiter(c.limiter))
	return err
}

func TestLimiterRate(t *testing.T) {
	tests := []struct {
		name string
		cfg  *config.Config
		// want is the rate that the classic client applies to the config.
		want int
	}{
		{"unset rate is the classic default", &config.Config{}, 15},
		{"nil config is the classic default", nil, 15},
		{"one call per second", &config.Config{RateLimitPerSecond: 1}, 1},
		{"rate below the default", &config.Config{RateLimitPerSecond: 4}, 4},
		{"rate of the default", &config.Config{RateLimitPerSecond: 15}, 15},
		{"rate above the default", &config.Config{RateLimitPerSecond: 100}, 100},
		{"transport without an opinion on retries",
			&config.Config{RateLimitPerSecond: 4, HTTPTransport: http.DefaultTransport}, 4},
		{"transport that retries on IO errors",
			&config.Config{RateLimitPerSecond: 4, HTTPTransport: fixtureRegistry{}}, 4},
		{"transport that doesn't retry on IO errors",
			&config.Config{RateLimitPerSecond: 4, HTTPTransport: fixtureRegistry{skipRetryOnIO: true}}, unlimited},
		{"mapping fixtures of the classic SDK",
			&config.Config{HTTPTransport: fixtures.MappingTransport{}}, unlimited},
		{"slice fixtures of the classic SDK",
			&config.Config{HTTPTransport: fixtures.SliceTransport{}}, unlimited},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			forgetNilConfig(t)
			synctest.Test(t, func(t *testing.T) {
				limiter := Limiter(tc.cfg)
				wait := func() error { return limiter.Wait(t.Context()) }
				n := 2*tc.want + 40
				assertAdmissions(t, wantAdmissions(tc.want, n), admissions(t, n, wait))
				// An idle budget fills up to the burst and no further.
				time.Sleep(time.Hour)
				assertAdmissions(t, wantAdmissions(tc.want, n), admissions(t, n, wait))
			})
		})
	}
}

// TestLimiterHoldsModularCallsToClassicRate compares what a registry receives
// from a classic client with what it receives from a modular client that uses
// the limiter, for one config.
func TestLimiterHoldsModularCallsToClassicRate(t *testing.T) {
	plain := func(r *registry) http.RoundTripper { return r }
	fixture := func(skipRetryOnIO bool) func(*registry) http.RoundTripper {
		return func(r *registry) http.RoundTripper { return fixtureRegistry{r, skipRetryOnIO} }
	}
	tests := []struct {
		name      string
		perSecond int
		transport func(*registry) http.RoundTripper
		want      int
	}{
		{"unset rate", 0, plain, 15},
		{"one call per second", 1, plain, 1},
		{"rate below the default", 4, plain, 4},
		{"rate above the default", 100, plain, 100},
		{"transport that retries on IO errors", 4, fixture(false), 4},
		{"transport that doesn't retry on IO errors", 4, fixture(true), unlimited},
		{"transport that doesn't retry on IO errors, unset rate", 0, fixture(true), unlimited},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				api := &registry{}
				c := newClients(t, "https://registry.test", tc.perSecond, tc.transport(api))
				n := 2*tc.want + 40
				want := wantAdmissions(tc.want, n)

				start := time.Now()
				admissions(t, n, func() error { return c.classicCall(t.Context()) })
				classic := api.take(start)

				start = time.Now()
				admissions(t, n, func() error { return c.modularCall(t.Context()) })
				modular := api.take(start)

				assertAdmissions(t, want, classic)
				assertAdmissions(t, want, modular)
				assertAdmissions(t, classic, modular)
			})
		})
	}
}

func TestLimiterBudgets(t *testing.T) {
	one, changed := &config.Config{RateLimitPerSecond: 3}, &config.Config{RateLimitPerSecond: 3}
	tests := []struct {
		name          string
		first, second *config.Config
		// change runs between the calls of Limiter for first and for second.
		change func()
		// spend is the number of calls that take the whole budget of first.
		spend int
		// want is when the limiter of second admits calls after that.
		want []time.Duration
	}{
		{name: "one config has one budget",
			first: one, second: one, spend: 3, want: wantAdmissions(3, 9)[3:]},
		{name: "the nil config has one budget",
			first: nil, second: nil, spend: 15, want: wantAdmissions(15, 45)[15:]},
		{name: "a config keeps its budget when its rate and its transport change",
			first: changed, second: changed, spend: 3, want: wantAdmissions(3, 9)[3:],
			change: func() {
				changed.RateLimitPerSecond = 5
				changed.HTTPTransport = fixtures.MappingTransport{}
			}},
		{name: "configs with equal fields have a budget each",
			first: &config.Config{RateLimitPerSecond: 3}, second: &config.Config{RateLimitPerSecond: 3},
			spend: 3, want: wantAdmissions(3, 9)},
		{name: "configs with different rates have a budget each",
			first: &config.Config{RateLimitPerSecond: 3}, second: &config.Config{RateLimitPerSecond: 5},
			spend: 3, want: wantAdmissions(5, 15)},
		{name: "a config and the nil config have a budget each",
			first: &config.Config{}, second: nil, spend: 15, want: wantAdmissions(15, 45)},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			forgetNilConfig(t)
			synctest.Test(t, func(t *testing.T) {
				first := Limiter(tc.first)
				if tc.change != nil {
					tc.change()
				}
				second := Limiter(tc.second)
				spent := admissions(t, tc.spend, func() error { return first.Wait(t.Context()) })
				assertAdmissions(t, make([]time.Duration, tc.spend), spent)

				got := admissions(t, len(tc.want), func() error { return second.Wait(t.Context()) })
				assertAdmissions(t, tc.want, got)
			})
		})
	}
}

func TestLimiterConcurrentWaits(t *testing.T) {
	tests := []struct {
		name string
		// rates are the rate limits of the configs that wait at the same time,
		// and want the rates that their waits are held to.
		rates, want []int
		// waits is the number of concurrent waits on each config.
		waits int
	}{
		{"one config at the default rate", []int{0}, []int{15}, 50},
		{"one config at a set rate", []int{4}, []int{4}, 21},
		{"two configs at one rate", []int{4, 4}, []int{4, 4}, 21},
		{"three configs at different rates", []int{0, 2, 7}, []int{15, 2, 7}, 33},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				var mu sync.Mutex
				var wg sync.WaitGroup
				admitted := make([][]time.Duration, len(tc.rates))
				start := time.Now()
				for i, perSecond := range tc.rates {
					cfg := &config.Config{RateLimitPerSecond: perSecond}
					for range tc.waits {
						wg.Go(func() {
							// Every wait asks for the limiter itself, so that the
							// first ones race to create the budget of the config.
							err := Limiter(cfg).Wait(t.Context())
							mu.Lock()
							defer mu.Unlock()
							assert.NoError(t, err)
							admitted[i] = append(admitted[i], time.Since(start))
						})
					}
				}
				wg.Wait()
				for i, got := range admitted {
					slices.Sort(got)
					assertAdmissions(t, wantAdmissions(tc.want[i], tc.waits), got)
				}
			})
		})
	}
}

// limiterFailures are the ways in which a wait fails. The budget of every
// case refills in a second, and the contexts end well before that.
var limiterFailures = []struct {
	name      string
	perSecond int
	// spend is the number of calls that pass before the call that fails.
	spend int
	// context returns the context of the call that fails.
	context func(t *testing.T) context.Context
	// want is the error of the call, and cause the context error that it
	// wraps, if any.
	want  string
	cause error
	// after is how long the call waits before it fails.
	after time.Duration
}{
	{
		name:      "context canceled before the wait",
		perSecond: 1,
		context: func(t *testing.T) context.Context {
			ctx, cancel := context.WithCancel(t.Context())
			cancel()
			return ctx
		},
		want:  "failed in rate limiter: context canceled",
		cause: context.Canceled,
	},
	{
		name:      "context past its deadline before the wait",
		perSecond: 1,
		context: func(t *testing.T) context.Context {
			ctx, cancel := context.WithDeadline(t.Context(), time.Now().Add(-time.Minute))
			t.Cleanup(cancel)
			return ctx
		},
		want:  "failed in rate limiter: context deadline exceeded",
		cause: context.DeadlineExceeded,
	},
	{
		name:      "context canceled during the wait",
		perSecond: 1,
		spend:     1,
		context: func(t *testing.T) context.Context {
			ctx, cancel := context.WithCancel(t.Context())
			timer := time.AfterFunc(50*time.Millisecond, cancel)
			t.Cleanup(func() { timer.Stop() })
			return ctx
		},
		want:  "failed in rate limiter: context canceled",
		cause: context.Canceled,
		after: 50 * time.Millisecond,
	},
	{
		name:      "deadline before the turn of the call",
		perSecond: 1,
		spend:     1,
		context: func(t *testing.T) context.Context {
			ctx, cancel := context.WithTimeout(t.Context(), 250*time.Millisecond)
			t.Cleanup(cancel)
			return ctx
		},
		want: "failed in rate limiter: rate: Wait(n=1) would exceed context deadline",
	},
	{
		name:      "negative rate",
		perSecond: -1,
		context:   func(t *testing.T) context.Context { return t.Context() },
		want:      "failed in rate limiter: rate: Wait(n=1) exceeds limiter's burst -1",
	},
}

func TestLimiterWaitFails(t *testing.T) {
	for _, tc := range limiterFailures {
		t.Run(tc.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				limiter := Limiter(&config.Config{RateLimitPerSecond: tc.perSecond})
				for range tc.spend {
					require.NoError(t, limiter.Wait(t.Context()))
				}

				start := time.Now()
				err := limiter.Wait(tc.context(t))

				require.EqualError(t, err, tc.want)
				assert.Equal(t, tc.after, time.Since(start))
				assertContextCause(t, tc.cause, err)
			})
		})
	}
}

// assertContextCause checks that errors.Is finds the context error cause in
// err and no other one.
func assertContextCause(t *testing.T, cause, err error) {
	t.Helper()
	for _, target := range []error{context.Canceled, context.DeadlineExceeded} {
		assert.Equal(t, target == cause, errors.Is(err, target), "errors.Is(%q, %q)", err, target)
	}
}

// TestLimiterFailsAsClassicClient compares what the caller of a classic client
// gets back with what the caller of a modular client that uses the limiter
// gets back, when the limiter fails the call, and what the server receives
// from each of them.
func TestLimiterFailsAsClassicClient(t *testing.T) {
	for _, tc := range limiterFailures {
		t.Run(tc.name, func(t *testing.T) {
			api := &registry{}
			server := httptest.NewServer(api)
			defer server.Close()
			c := newClients(t, server.URL, tc.perSecond, nil)
			start := time.Now()
			// fail makes the calls that pass and the call that fails. It returns
			// the number of calls that the server received and the error.
			fail := func(makeCall func(context.Context) error) (int, error) {
				for range tc.spend {
					require.NoError(t, makeCall(t.Context()))
				}
				err := makeCall(tc.context(t))
				return len(api.take(start)), err
			}

			classicReceived, classicErr := fail(c.classicCall)
			modularReceived, modularErr := fail(c.modularCall)

			require.EqualError(t, classicErr, tc.want)
			require.EqualError(t, modularErr, tc.want)
			assertContextCause(t, tc.cause, classicErr)
			assertContextCause(t, tc.cause, modularErr)
			assert.Equal(t, tc.spend, classicReceived)
			assert.Equal(t, tc.spend, modularReceived)
		})
	}
}
