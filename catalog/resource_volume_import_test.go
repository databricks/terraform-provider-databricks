package catalog_test

import (
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/terraform-provider-databricks/qa/terraformtest"
)

// TestVolumeImportUserAgent pins the User-Agent of the two GETs that an import makes,
// on both clients. The default importer reads outside the CRUD wrappers that add
// sdk/sdkv2 and resource/volume (common/context.go), so only the refresh GET after
// import carries them. YAML header matchers compare exact values, and the User-Agent
// also holds version and environment tokens, so the test compares the two values in Go.
func TestVolumeImportUserAgent(t *testing.T) {
	type factories = map[string]func() (tfprotov6.ProviderServer, error)
	for name, providers := range map[string]func(func(*config.Config) error) factories{
		"existing_path": func(customize func(*config.Config) error) factories {
			return terraformtest.ProviderFactories(customize)
		},
		"sdk_go": func(customize func(*config.Config) error) factories {
			return terraformtest.ProviderFactoriesWithSdkGo(customize,
				[]string{"databricks_volume"})
		},
	} {
		t.Run(name, func(t *testing.T) {
			recorder := &importUserAgentRecorder{}
			terraformtest.RunNative(t, "testdata/volume_import_user_agent.yaml",
				func(tc *resource.TestCase, customize func(*config.Config) error) {
					tc.ProtoV6ProviderFactories = providers(func(c *config.Config) error {
						if err := customize(c); err != nil {
							return err
						}
						c.HTTPTransport = importUserAgentTransport{recorder, c.HTTPTransport}
						return nil
					})
				})
			agents := recorder.snapshot()
			require.Len(t, agents, 2)
			importer, refresh := agents[0], agents[1]
			const wrapperTokens = " sdk/sdkv2 resource/volume"
			assert.Contains(t, refresh, wrapperTokens)
			assert.Equal(t, strings.Replace(refresh, wrapperTokens, "", 1), importer)
			assert.Equal(t, name == "sdk_go", strings.Contains(importer, " sdk-go-uc-volumes/"),
				importer)
		})
	}
}

// importUserAgentRecorder keeps the User-Agent of each volume call, in order.
type importUserAgentRecorder struct {
	mu     sync.Mutex
	agents []string
}

func (r *importUserAgentRecorder) record(req *http.Request) {
	if !strings.HasPrefix(req.URL.Path, "/api/2.1/unity-catalog/volumes") {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.agents = append(r.agents, req.Header.Get("User-Agent"))
}

func (r *importUserAgentRecorder) snapshot() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]string(nil), r.agents...)
}

type importUserAgentTransport struct {
	recorder  *importUserAgentRecorder
	transport http.RoundTripper
}

func (u importUserAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	u.recorder.record(req)
	return u.transport.RoundTrip(req)
}
