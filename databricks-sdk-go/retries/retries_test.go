package retries

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/stretchr/testify/assert"
)

func TestRetriesDefaults(t *testing.T) {
	r := Retrier[struct{}]{}
	assert.Equal(t, r.config.Timeout(), 20*time.Minute)
	assert.Equal(t, r.config.ShouldRetry(nil), false)
	assert.Equal(t, r.config.ShouldRetry(errors.New("an error")), true)
}

func forTesting[T any](opts ...RetryOption) Retrier[T] {
	r := New[T](opts...)
	r.config.backoff = func(_ int) time.Duration {
		return 100 * time.Millisecond
	}
	return r
}

func TestRetriesRunSuccess(t *testing.T) {
	r := forTesting[int]()
	res, err := r.Run(context.Background(), func(_ context.Context) (*int, error) {
		res := 3
		return &res, nil
	})
	assert.NoError(t, err)
	assert.Equal(t, *res, 3)
}

func TestRetriesRunSuccessAfterError(t *testing.T) {
	r := forTesting[int]()
	hasRun := false
	res, err := r.Run(context.Background(), func(_ context.Context) (*int, error) {
		if !hasRun {
			hasRun = true
			return nil, errors.New("an error")
		}
		res := 3
		return &res, nil
	})
	assert.NoError(t, err)
	assert.Equal(t, *res, 3)
}

func TestRetriesRunError(t *testing.T) {
	r := forTesting[int](WithTimeout(200 * time.Millisecond))
	_, err := r.Run(context.Background(), func(_ context.Context) (*int, error) {
		return nil, errors.New("an error")
	})
	assert.Equal(t, err.Error(), "timed out: an error")
}

func TestRetriesRunOnErrorRetry(t *testing.T) {
	e := errors.New("an error")
	hasRun := false
	r := forTesting[int](OnErrors(e))
	_, err := r.Run(context.Background(), func(_ context.Context) (*int, error) {
		if !hasRun {
			hasRun = true
			return nil, e
		}
		res := 3
		return &res, nil
	})
	assert.NoError(t, err)
}

func TestRetriesRunOnErrorNoRetry(t *testing.T) {
	e := errors.New("an error")
	hasRun := false
	r := forTesting[int](OnErrors(e))
	_, err := r.Run(context.Background(), func(_ context.Context) (*int, error) {
		if !hasRun {
			hasRun = true
			return nil, errors.New("a different error")
		}
		res := 3
		return &res, nil
	})
	assert.Equal(t, err.Error(), "a different error")
}

func TestRetriesRunWithRetryFunc(t *testing.T) {
	hasRun := false
	r := forTesting[int](WithRetryFunc(func(err error) bool {
		return strings.Contains(err.Error(), "hi")
	}))
	_, err := r.Run(context.Background(), func(_ context.Context) (*int, error) {
		if !hasRun {
			hasRun = true
			return nil, errors.New("hi hello")
		}
		res := 3
		return &res, nil
	})
	assert.NoError(t, err)
}

func TestRetriesRunWithRetryFuncNoRetry(t *testing.T) {
	hasRun := false
	r := forTesting[int](WithRetryFunc(func(err error) bool {
		return strings.Contains(err.Error(), "hi")
	}))
	_, err := r.Run(context.Background(), func(_ context.Context) (*int, error) {
		if !hasRun {
			hasRun = true
			return nil, errors.New("oh no")
		}
		res := 3
		return &res, nil
	})
	assert.Equal(t, err.Error(), "oh no")
}

func TestRetriesNegativeTimeout(t *testing.T) {
	r := forTesting[int](WithTimeout(time.Duration(-1)))
	assert.Equal(t, r.config.Timeout(), time.Duration(-1))
}

func ExampleNew() {
	// Default retries for 20 minutes on any error
	New[int]()

	// Override the timeout with NewTimeout
	New[int](WithTimeout(5 * time.Minute))

	// Retry on specific errors only with OnErrors
	New[int](OnErrors(apierr.ErrResourceConflict))

	// Customize retry logic with WithRetryFunc
	retryCount := 3
	New[int](WithRetryFunc(func(err error) bool {
		if retryCount > 0 {
			retryCount--
			return true
		}
		return false
	}))
}

func ExampleRetrier_Run() {
	hasRun := false
	e := errors.New("an error")
	res, _ := New[string](WithTimeout(5*time.Minute), OnErrors(e)).Run(
		context.Background(),
		func(ctx context.Context) (*string, error) {
			if !hasRun {
				hasRun = true
				fmt.Println("failed, retrying")
				return nil, e
			}
			fmt.Println("succeeded")
			res := "hello!"
			return &res, nil
		},
	)
	fmt.Println(*res)
	// Output:
	// failed, retrying
	// succeeded
	// hello!
}

func ExampleRetrier_Wait() {
	hasRun := false
	e := errors.New("an error")
	err := New[string](WithTimeout(5*time.Minute), OnErrors(e)).Wait(
		context.Background(),
		func(ctx context.Context) error {
			if !hasRun {
				hasRun = true
				fmt.Println("failed, retrying")
				return e
			}
			fmt.Println("succeeded")
			return nil
		},
	)
	fmt.Println(err)
	// Output:
	// failed, retrying
	// succeeded
	// <nil>
}
