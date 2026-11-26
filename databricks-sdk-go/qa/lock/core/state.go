package core

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/qa"
)

// LockState represents the state of a lock.
type LockState struct {
	// Lockable is the lock id of the lockable resource.
	Lockable string

	// LeaseId is the lease id of the lock.
	LeaseId string

	// Start is the time the lock was acquired.
	Start time.Time

	// Expiry is the time the lock expires.
	Expiry time.Time

	// IsInDebug is true if the test is running in debug mode.
	IsInDebug bool

	// Test is the name of the test, if applicable.
	Test string

	// User is the user running the test.
	User string `env:"USER"`

	// GithubActions is true if the test is running in Github Actions.
	GithubActions string `env:"GITHUB_ACTIONS"`

	// GithubRef is the ref of the Github Actions workflow.
	GithubRef string `env:"GITHUB_REF"`

	// GithubRepository is the repository of the Github Actions workflow, e.g. `databricks/databricks-sdk-go`.
	GithubRepository string `env:"GITHUB_REPOSITORY"`

	// GithubRunId is the ID of the Github Actions run.
	GithubRunId string `env:"GITHUB_RUN_ID"`

	// GithubRunNumber is the number of the Github Actions run.
	GithubRunNumber string `env:"GITHUB_RUN_NUMBER"`

	// GithubSha is the SHA of the Github Actions commit.
	GithubSha string `env:"GITHUB_SHA"`

	// GithubWorkflow is the name of the Github Actions workflow.
	GithubWorkflow string `env:"GITHUB_WORKFLOW"`
}

func (l *LockState) loadFromEnv() {
	rv := reflect.ValueOf(l).Elem()
	t := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		tag := t.Field(i).Tag.Get("env")
		if tag == "" {
			continue
		}
		val := os.Getenv(tag)
		if val != "" {
			field.SetString(val)
		}
	}
}

// IsValid() returns true if the lock has not yet expired.
func (l *LockState) IsValid() bool {
	return time.Now().Before(l.Expiry)
}

// Extend() resets the lock's expiry to the current time plus the given duration.
// This method only updates the state object; it does not actually persist the
// change to the lock backend.
func (l *LockState) Extend(duration time.Duration) {
	l.Expiry = time.Now().Add(duration)
}

// NewLockState creates a new LockState object with the given lockable, lease ID,
// and lease duration. If t is not nil, the test name is also set.
func NewLockState(lockable Lockable, leaseId string, leaseDuration time.Duration, t *testing.T) *LockState {
	now := time.Now()
	state := &LockState{
		Start:     now,
		Expiry:    now.Add(leaseDuration),
		Lockable:  lockable.GetLockId(),
		LeaseId:   leaseId,
		IsInDebug: qa.IsInDebug(),
	}
	state.loadFromEnv()
	if t != nil {
		state.Test = t.Name()
	}
	return state
}
