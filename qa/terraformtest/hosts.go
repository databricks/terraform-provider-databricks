package terraformtest

import (
	"fmt"
	"maps"
	"net/url"
	"slices"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

// namedHost is a host the provider may reach besides the local fixture server,
// such as an account host or a workspace's own host. The local server answers it
// too; calls select it by name and its metadata answers host discovery there.
type namedHost struct {
	URL      string         `yaml:"url"`
	Metadata map[string]any `yaml:"metadata,omitempty"`
}

func (d *testDefinition) validateHosts() error {
	if d.ProviderHost != "" {
		host, ok := d.Hosts[d.ProviderHost]
		if !ok {
			return fmt.Errorf("provider_host %q is not one of hosts", d.ProviderHost)
		}
		if d.HostMetadata != nil && host.Metadata != nil {
			return fmt.Errorf("host_metadata and hosts.%s.metadata both describe the "+
				"provider host; keep one", d.ProviderHost)
		}
	}
	authorities := map[string]string{}
	for _, name := range slices.Sorted(maps.Keys(d.Hosts)) {
		if name == "" {
			return fmt.Errorf("hosts: a host name must not be empty")
		}
		u, err := parseHostURL(d.Hosts[name].URL)
		if err != nil {
			return fmt.Errorf("hosts.%s.url: %w", name, err)
		}
		// The local server tells hosts apart by the Host header, which has no scheme.
		if other, duplicate := authorities[u.Host]; duplicate {
			return fmt.Errorf("hosts.%s.url: %s is also the host of hosts.%s",
				name, u.Host, other)
		}
		authorities[u.Host] = name
	}
	return nil
}

// parseHostURL accepts only an origin, so a fixture path cannot hide in the host.
func parseHostURL(raw string) (*url.URL, error) {
	u, err := url.Parse(raw)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" ||
		u.User != nil || (u.Path != "" && u.Path != "/") || u.RawQuery != "" ||
		u.Fragment != "" {
		return nil, fmt.Errorf("expected an http or https URL without a path, query "+
			"or fragment, got %q", raw)
	}
	return u, nil
}

// setHosts gives the handler each named host's authority. The definition's
// validation has already checked every URL.
func (d *testDefinition) setHosts(handler *httpfixture.Handler) error {
	authorities := map[string]string{}
	for name, host := range d.Hosts {
		u, err := parseHostURL(host.URL)
		if err != nil {
			return err
		}
		authorities[name] = u.Host
	}
	return handler.SetHosts(authorities, d.ProviderHost)
}

// replayHosts are the origins the provider may reach during replay. The local
// fixture server serves all of them.
type replayHosts struct {
	server   string
	named    []*url.URL
	provider string
}

func (d *testDefinition) replayHosts(server string) replayHosts {
	hosts := replayHosts{server: server}
	for _, name := range slices.Sorted(maps.Keys(d.Hosts)) {
		u, _ := parseHostURL(d.Hosts[name].URL)
		hosts.named = append(hosts.named, u)
		if name == d.ProviderHost {
			hosts.provider = d.Hosts[name].URL
		}
	}
	return hosts
}

// providerHost keeps an HCL host that names a journey host, so a step can move
// the provider between hosts. Without named hosts the HCL host is ignored, as it
// always was, so journeys cannot reach anything but the local server.
func (h replayHosts) providerHost(configured string) (string, error) {
	switch {
	case len(h.named) == 0:
		return h.server, nil
	case configured == "" && h.provider != "":
		return h.provider, nil
	case configured == "":
		return h.server, nil
	}
	u, err := url.Parse(configured)
	if err == nil && u.Host == "" {
		// The SDK assumes https for a host without a scheme.
		u, err = url.Parse("https://" + configured)
	}
	if err == nil && isNamedHost(h.named, u) {
		return configured, nil
	}
	return "", fmt.Errorf("hermetic replay: provider host %s is not one of the "+
		"journey's hosts", configured)
}

func isNamedHost(named []*url.URL, u *url.URL) bool {
	return slices.ContainsFunc(named, func(host *url.URL) bool {
		return host.Scheme == u.Scheme && host.Host == u.Host
	})
}
