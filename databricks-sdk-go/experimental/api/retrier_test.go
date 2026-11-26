package api

import (
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func deterministicRand(n int64) int64 {
	return n - 1 // always return max value
}

func TestRetrier_IsRetriable_success(t *testing.T) {
	testCases := []struct {
		fn        func(error) bool
		wantDelay time.Duration
		wantOk    bool
	}{
		{
			wantDelay: 100 * time.Millisecond,
			wantOk:    true,
			fn: func(err error) bool {
				return true
			},
		},
		{
			wantDelay: 0,
			wantOk:    false,
			fn: func(err error) bool {
				return false
			},
		},
	}

	for _, tc := range testCases {
		bp := BackoffPolicy{
			Initial: tc.wantDelay,
			int63n:  deterministicRand,
		}

		r := RetryOn(bp, tc.fn)
		gotDelay, gotOk := r.IsRetriable(errors.New("an error"))

		if gotOk != tc.wantOk {
			t.Errorf("IsRetriable() = %v, want %v", gotOk, tc.wantOk)
		}
		if gotDelay != tc.wantDelay {
			t.Errorf("Delay() = %v, want %v", gotDelay, tc.wantDelay)
		}
	}
}

func TestBackoffPolicy_Delay_initializeation(t *testing.T) {
	testCases := []struct {
		name   string
		bp     BackoffPolicy
		wantBP BackoffPolicy
	}{
		{
			name: "default",
			bp:   BackoffPolicy{},
			wantBP: BackoffPolicy{
				Initial: 1 * time.Second,
				Maximum: 60 * time.Second,
				Factor:  2,
			},
		},
		{
			name: "custom initial smaller than maximum",
			bp: BackoffPolicy{
				Initial: 100 * time.Millisecond,
			},
			wantBP: BackoffPolicy{
				Initial: 100 * time.Millisecond,
				Maximum: 60 * time.Second,
				Factor:  2,
			},
		},
		{
			name: "custom initial greater than maximum",
			bp: BackoffPolicy{
				Initial: 10 * time.Second,
				Maximum: 1 * time.Second,
			},
			wantBP: BackoffPolicy{
				Initial: 1 * time.Second,
				Maximum: 1 * time.Second,
				Factor:  2,
			},
		},
		{
			name: "custom factor less than 1",
			bp: BackoffPolicy{
				Factor: 0.5,
			},
			wantBP: BackoffPolicy{
				Initial: 1 * time.Second,
				Maximum: 60 * time.Second,
				Factor:  2,
			},
		},
		{
			name: "custom factor greater than 1",
			bp: BackoffPolicy{
				Factor: 1.5,
			},
			wantBP: BackoffPolicy{
				Initial: 1 * time.Second,
				Maximum: 60 * time.Second,
				Factor:  1.5,
			},
		},
	}
	for _, tc := range testCases {
		tc.bp.Delay() // initializes defaults

		if diff := cmp.Diff(tc.bp, tc.wantBP, cmpopts.IgnoreUnexported(BackoffPolicy{})); diff != "" {
			t.Errorf("unexpected error (-want +got):\n%s", diff)
		}
	}
}

func TestBackoffPolicy_Delay_exponential(t *testing.T) {
	bp := BackoffPolicy{
		Initial: 100 * time.Millisecond,
		Maximum: 10 * time.Second,
		Factor:  2.0,
		int63n:  deterministicRand,
	}

	wantDelays := []time.Duration{
		100 * time.Millisecond,
		200 * time.Millisecond,
		400 * time.Millisecond,
		800 * time.Millisecond,
		1600 * time.Millisecond,
		3200 * time.Millisecond,
		6400 * time.Millisecond,
		10000 * time.Millisecond, // capped by Maximum
	}

	for i, want := range wantDelays {
		got := bp.Delay()

		if got != want {
			t.Errorf("Delay() %d = %v, want %v", i+1, got, want)
		}
	}
}
