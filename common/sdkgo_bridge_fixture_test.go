package common

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// For a test fixture's transport, httpclient.NewApiClient drops the configured retry
// timeout as well as the rate limit, and the bridge drops it too. Both then retry for
// the retries package's default of 20 minutes, which retries.Poll gives the classic
// client for a zero timeout.
func TestSdkGoBridgeSkipsRetryTimeoutForTestFixtures(t *testing.T) {
	cfg := newCredentialsTestConfig(t, &countingCredentials{})
	cfg.HTTPTransport = skipRetryTransport{}
	cfg.RetryTimeoutSeconds = 30
	classic := newBridgeTestClient(t, cfg)
	b, err := newSdkGoBridge(classic.Config, classic.ApiClient())
	require.NoError(t, err)

	// httpclient.ApiClient keeps its ClientConfig unexported.
	classicTimeout := reflect.ValueOf(classic.ApiClient()).Elem().
		FieldByName("config").FieldByName("RetryTimeout").Int()
	assert.Zero(t, classicTimeout)
	assert.Equal(t, 20*time.Minute, b.retry.Timeout())
	before := time.Now()
	r := b.newRetrier(context.Background(), &sdkGoAttempt{})
	assert.WithinDuration(t, before.Add(20*time.Minute), r.deadline, time.Minute)
}
