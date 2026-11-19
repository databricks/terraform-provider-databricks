package terraformtest

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/databricks/databricks-sdk-go/config"
)

type recordingProfile struct{ config *config.Config }

// Load only the named file profile, never ambient host/token environment values.
// ConfigFile.Configure does not resolve metadata or authenticate, so validation
// and unsupported-feature checks complete before any network request.
func loadRecordingProfile(name string) (*recordingProfile, error) {
	cfg := &config.Config{Profile: name, ConfigFile: os.Getenv("DATABRICKS_CONFIG_FILE")}
	if err := config.ConfigFile.Configure(cfg); err != nil {
		return nil, fmt.Errorf("cannot load recorder profile %q; check .databrickscfg", name)
	}
	host, err := url.Parse(cfg.Host)
	if err != nil || host.Host == "" || host.User != nil || host.RawQuery != "" || host.Fragment != "" ||
		(host.Scheme != "https" && host.Scheme != "http") || strings.Trim(host.Path, "/") != "" {
		return nil, fmt.Errorf("recorder profile must contain an HTTP(S) host without credentials, path or query")
	}
	return &recordingProfile{config: cfg}, nil
}

// Only connectivity and authentication come from the profile. Other provider
// settings in HCL, including deliberately invalid configurations, remain under test.
func (p *recordingProfile) customize(transport http.RoundTripper) func(*config.Config) error {
	return func(cfg *config.Config) error {
		for _, attr := range config.ConfigAttributes {
			if attr.HasAuthAttribute() || attr.Sensitive || attr.Name == "host" || attr.Name == "auth_type" ||
				attr.Name == "account_id" || attr.Name == "workspace_id" || attr.Name == "profile" || attr.Name == "config_file" {
				if err := attr.SetS(cfg, attr.GetString(p.config)); err != nil {
					return fmt.Errorf("cannot configure recorder profile field %s", attr.Name)
				}
			}
		}
		cfg.Loaders = []config.Loader{fixtureLoader{}}
		cfg.HTTPTransport = recordingTransport{host: strings.TrimRight(p.config.Host, "/"), capture: transport, base: p.transport()}
		return nil
	}
}

func (p *recordingProfile) transport() http.RoundTripper {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	return transport
}

type recordingTransport struct {
	host    string
	capture http.RoundTripper
	base    http.RoundTripper
}

// OAuth traffic must reach the selected credential strategy but must never be
// recorded as a test fixture. Calls to another API host cannot be represented by
// the single-host replay format and fail instead of producing a partial transcript.
func (r recordingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	path := req.URL.Path
	if strings.HasPrefix(path, "/oidc/") || strings.HasPrefix(path, "/oauth") ||
		strings.HasPrefix(path, "/.well-known/oauth-") || strings.HasPrefix(path, "/.well-known/openid-") {
		return r.base.RoundTrip(req)
	}
	if req.URL.Scheme+"://"+req.URL.Host != r.host {
		return nil, fmt.Errorf("recording requires one API host; use a profile for this test's API")
	}
	return r.capture.RoundTrip(req)
}
