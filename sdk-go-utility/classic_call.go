package sdkgoutility

import (
	"context"
	"errors"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/sdk-go/options/call"
)

// ClassicCall runs method, which wraps one method of a modular SDK client, as
// the classic client of cfg runs a request. It puts the other helpers of this
// package together, so that a resource needs one line for a call:
//
//   - it retries by a new Retrier of cfg, by the rules and until the retry
//     timeout of the classic client;
//   - it waits on Limiter(cfg) before every attempt, so all calls made with
//     one cfg share the rate limit of that cfg;
//   - it returns the error that ClassicError makes of a failure, which is also
//     the error that the retrier decides on.
//
// On failure ClassicCall returns the zero value of T. It returns an error
// without running method when cfg or method is nil, and the error of Retrier
// when cfg doesn't resolve.
//
// ClassicCall leaves to the caller:
//
//   - the modular client. ClassicCall takes from cfg only the retry timeout,
//     the rules of what to retry and the rate. The host, the credentials, the
//     HTTP client and the User-Agent are those of the client that method
//     calls.
//   - the timeout of that client. A modular client ends a call after 5 minutes
//     unless client.WithTimeout or call.WithTimeout says otherwise, and the
//     classic client has no such limit, so a client that is to retry as long
//     as the classic one is built with client.WithTimeout(0).
//   - cfg itself: the resolved config of the provider, and the same pointer
//     for every call, because the rate limit is kept by pointer. A config
//     that isn't resolved is resolved by its first call, which calls that run
//     at once must not do: a config.Config can't be resolved by two of them.
//   - a method that makes exactly one call of the client and returns its
//     result. A method that returns an iterator makes its requests after it
//     has returned, outside ClassicCall.
func ClassicCall[T any](
	ctx context.Context,
	cfg *config.Config,
	method func(context.Context, ...call.Option) (T, error),
) (T, error) {
	var zero T
	if cfg == nil {
		return zero, errors.New("sdkgoutility.ClassicCall: cfg is nil")
	}
	if method == nil {
		return zero, errors.New("sdkgoutility.ClassicCall: method is nil")
	}
	classic := func(err error) error { return ClassicError(ctx, err) }
	retrier, err := Retrier(ctx, cfg, classic)
	if err != nil {
		return zero, err
	}
	return Call(ctx, method, retrier, Limiter(cfg), classic)
}
