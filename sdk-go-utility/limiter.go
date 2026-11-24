package sdkgoutility

import (
	"context"
	"fmt"
	"sync"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/sdk-go/core/ops"
	"golang.org/x/time/rate"
)

// defaultRateLimitPerSecond is the rate of a config that sets none. It is the
// value of defaultRateLimitPerSeconds in package httpclient of the classic
// SDK, which is unexported.
const defaultRateLimitPerSecond = 15

// budgets holds the limiter of every config that Limiter has seen, as an
// ops.Limiter by *config.Config.
var budgets sync.Map

// Limiter returns the rate limiter for the modular SDK calls made with cfg.
// The classic client waits on a rate limiter before every request, so that a
// large apply doesn't flood the workspace API. A modular client has none, and
// limits only the calls that are given one: call.WithLimiter(Limiter(cfg)).
//
// The limiter is built as httpclient.NewApiClient of the classic SDK builds
// the limiter of a classic client for cfg. It admits cfg.RateLimitPerSecond
// calls per second, or 15 when that is unset, and at most that many at once.
// It doesn't hold back the calls of a config whose HTTPTransport reports
// SkipRetryOnIO, as the HTTP fixtures of the classic SDK do. A nil cfg is
// limited as a config that sets no rate.
//
// Wait blocks until the limiter admits the call. It returns an error at once
// when ctx has already ended, when the deadline of ctx comes before the turn
// of the call or when the rate is negative, and as soon as ctx is canceled
// during the wait. The error starts with "failed in rate limiter: ", the
// message of the classic client, and wraps the error of the limiter. That is
// ctx.Err() when ctx has ended or is canceled, and otherwise an error of
// package rate, which errors.Is doesn't match with context.DeadlineExceeded.
//
// All limiters of one *config.Config share one budget, so the modular calls
// of one provider configuration are limited together, whichever client makes
// them. Two configs have separate budgets, even when their fields are equal.
// The budget is built on the first call for cfg, from the rate and the
// transport that cfg has then. It is kept, and keeps cfg, for as long as the
// process runs.
//
// Limiter leaves to the caller:
//
//   - resolving cfg. Limiter reads the rate that cfg holds and doesn't load
//     it from the provider block or the environment.
//   - passing the limiter to every call. A call without it isn't limited and
//     doesn't draw from the budget.
//   - how long a call may wait. The modular client hands Wait a ctx that
//     ends with the timeout of the call, which call.WithTimeout and
//     client.WithTimeout set.
//   - the budget of the classic client. Its limiter is unexported, so the
//     classic and the modular calls of one config are limited separately,
//     each at the rate of cfg.
func Limiter(cfg *config.Config) ops.Limiter {
	limiter, ok := budgets.Load(cfg)
	if !ok {
		limiter, _ = budgets.LoadOrStore(cfg, configLimiter{newBudget(cfg)})
	}
	return limiter.(ops.Limiter)
}

// newBudget returns the limiter that httpclient.NewApiClient of the classic
// SDK builds for cfg: the burst is the rate per second, also when the
// transport lifts the limit.
func newBudget(cfg *config.Config) *rate.Limiter {
	if cfg == nil {
		cfg = &config.Config{}
	}
	perSecond := cfg.RateLimitPerSecond
	if perSecond == 0 {
		perSecond = defaultRateLimitPerSecond
	}
	limit := rate.Limit(perSecond)
	skippable, ok := cfg.HTTPTransport.(interface{ SkipRetryOnIO() bool })
	if ok && skippable.SkipRetryOnIO() {
		limit = rate.Inf
	}
	return rate.NewLimiter(limit, perSecond)
}

// configLimiter waits on the budget of one config and reports a failed wait
// as the classic client does.
type configLimiter struct {
	budget *rate.Limiter
}

func (l configLimiter) Wait(ctx context.Context) error {
	if err := l.budget.Wait(ctx); err != nil {
		return fmt.Errorf("failed in rate limiter: %w", err)
	}
	return nil
}
