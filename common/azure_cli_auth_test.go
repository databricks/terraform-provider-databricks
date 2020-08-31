package common

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAzureCliAuth(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")
	// fake expiration date for az mock cli
	os.Setenv("EXPIRE", "15M")

	cnt := []int{0}
	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			cnt[0]++
			if req.RequestURI == "/api/2.0/clusters/list-zones" {
				assert.Equal(t, "Bearer ...", req.Header.Get("Authorization"))
				_, err := rw.Write([]byte(`{"zones": ["a", "b", "c"]}`))
				assert.NoError(t, err)
				return
			}
			assert.Fail(t, fmt.Sprintf("Received unexpected call: %s %s",
				req.Method, req.RequestURI))
		}))
	defer server.Close()

	client := DatabricksClient{
		Host: server.URL,
		AzureAuth: AzureAuth{
			ResourceID: "/a/b/c",
		},
		InsecureSkipVerify: true,
	}
	err := client.Configure()
	assert.NoError(t, err)

	type ZonesInfo struct {
		Zones       []string `json:"zones,omitempty"`
		DefaultZone string   `json:"default_zone,omitempty"`
	}
	var zi ZonesInfo
	err = client.Get("/clusters/list-zones", nil, &zi)
	assert.NoError(t, err)
	assert.NotNil(t, zi)
	assert.Len(t, zi.Zones, 3)

	err = client.Get("/clusters/list-zones", nil, &zi)
	assert.NoError(t, err)

	assert.Equal(t, 2, cnt[0], "There should be only one HTTP call")
}
