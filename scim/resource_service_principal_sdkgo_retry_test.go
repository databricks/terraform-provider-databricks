package scim

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"sync/atomic"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/common"
)

// These tests repeat the retry, timeout and cancellation tests of
// resource_service_principal_retry_test.go with reads through sdk-go scim/v1.

// readServicePrincipalWithSdkGo refreshes service principal abc through the resource
// with sdk-go selected, as Terraform does.
func readServicePrincipalWithSdkGo(ctx context.Context,
	c *common.DatabricksClient) (*schema.ResourceData, diag.Diagnostics) {
	r := ResourceServicePrincipal().ToResource()
	common.EnableSdkGo(r)
	d := r.TestResourceData()
	d.SetId("abc")
	ctx = context.WithValue(ctx, common.ResourceName, "service_principal")
	return d, r.ReadContext(ctx, d, c)
}

// newSdkGoPersistent429Server answers every service principal read with 429. sdk-go
// sends no attributes query, so the server matches the path only.
func newSdkGoPersistent429Server(t *testing.T) string {
	server := newServicePrincipalServer(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Empty(t, r.URL.RawQuery)
		w.WriteHeader(http.StatusTooManyRequests)
		_, err := w.Write([]byte(retryTestTooManyRequestsJSON))
		assert.NoError(t, err)
	})
	return server.URL
}

func TestServicePrincipalSdkGoReadFailsWithLast429AtRetryTimeout(t *testing.T) {
	transport := newAttemptTransport()
	// As on the existing path, a 3s retry timeout allows exactly two attempts.
	c := newRetryTestClient(t, &config.Config{
		Host:                newSdkGoPersistent429Server(t),
		RetryTimeoutSeconds: 3,
		HTTPTransport:       transport,
	})

	d, diags := readServicePrincipalWithSdkGo(context.Background(), c)

	assert.Equal(t, "cannot read service principal: Too many requests.",
		requireSingleDiagnostic(t, diags))
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, int32(2), transport.attempts.Load())
}

func TestServicePrincipalSdkGoReadRetriesConnectionReset(t *testing.T) {
	var served atomic.Int32
	server := newServicePrincipalServer(t, func(w http.ResponseWriter, r *http.Request) {
		if served.Add(1) == 1 {
			resetConnection(t, w)
			return
		}
		_, err := w.Write([]byte(retryTestServicePrincipalJSON))
		assert.NoError(t, err)
	})
	transport := newAttemptTransport()
	c := newRetryTestClient(t, &config.Config{Host: server.URL, HTTPTransport: transport})

	d, diags := readServicePrincipalWithSdkGo(context.Background(), c)

	require.Empty(t, diags)
	assert.Equal(t, "bcd", d.Get("application_id"))
	assert.Equal(t, int32(2), transport.attempts.Load())
	assert.Equal(t, int32(2), served.Load())
}

func TestServicePrincipalSdkGoReadRetriesRefusedConnection(t *testing.T) {
	var served atomic.Int32
	server := newServicePrincipalServer(t, func(w http.ResponseWriter, r *http.Request) {
		served.Add(1)
		_, err := w.Write([]byte(retryTestServicePrincipalJSON))
		assert.NoError(t, err)
	})
	refusedAddress := closedListenerAddress(t)
	transport := newAttemptTransport()
	transport.send = func(attempt int32, r *http.Request) (*http.Response, error) {
		if attempt == 1 {
			r = r.Clone(r.Context())
			r.URL.Host = refusedAddress
		}
		return transport.base.RoundTrip(r)
	}
	c := newRetryTestClient(t, &config.Config{Host: server.URL, HTTPTransport: transport})

	d, diags := readServicePrincipalWithSdkGo(context.Background(), c)

	require.Empty(t, diags)
	assert.Equal(t, "bcd", d.Get("application_id"))
	assert.Equal(t, int32(2), transport.attempts.Load())
	assert.Equal(t, int32(1), served.Load())
}

func TestServicePrincipalSdkGoReadDoesNotRetryHTTPTimeout(t *testing.T) {
	server := newServicePrincipalServer(t, func(w http.ResponseWriter, r *http.Request) {
		<-r.Context().Done()
	})
	transport := newAttemptTransport()
	c := newRetryTestClient(t, &config.Config{
		Host:                server.URL,
		HTTPTimeoutSeconds:  1,
		RetryTimeoutSeconds: 3,
		HTTPTransport:       transport,
	})

	d, diags := readServicePrincipalWithSdkGo(context.Background(), c)

	// The URL has no attributes query, unlike on the existing path.
	timeout := fmt.Sprintf(`Get "%s%s": request timed out after 1s of inactivity`,
		server.URL, retryTestServicePrincipalPath)
	assert.Equal(t, "cannot read service principal: "+timeout,
		requireSingleDiagnostic(t, diags))
	assert.Equal(t, "abc", d.Id())
	assert.Equal(t, int32(1), transport.attempts.Load())
}

func TestServicePrincipalSdkGoReadStopsRetryingWhenCancelledDuringBackoff(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	transport := newAttemptTransport()
	transport.send = func(_ int32, r *http.Request) (*http.Response, error) {
		resp, err := transport.base.RoundTrip(r)
		if err != nil {
			return nil, err
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		resp.Body = io.NopCloser(bytes.NewReader(body))
		cancel()
		return resp, err
	}
	c := newRetryTestClient(t, &config.Config{
		Host:          newSdkGoPersistent429Server(t),
		HTTPTransport: transport,
	})

	_, diags := readServicePrincipalWithSdkGo(ctx, c)

	assert.Equal(t, "cannot read service principal: Too many requests.",
		requireSingleDiagnostic(t, diags))
	assert.Equal(t, int32(1), transport.attempts.Load())
}
