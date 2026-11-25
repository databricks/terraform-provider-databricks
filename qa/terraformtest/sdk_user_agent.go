package terraformtest

import (
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strings"
	"sync"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

// sdkGoRequest is a request that a resource sends through sdk-go when a step selects it.
type sdkGoRequest struct {
	method, pathPrefix string
}

// sdkGoRequests lists, by resource type, the requests sent through sdk-go. The other
// requests of a selected resource, such as its writes, stay on the existing clients.
var sdkGoRequests = map[string][]sdkGoRequest{
	"databricks_volume": {{http.MethodGet, "/api/2.1/unity-catalog/volumes/"}},
}

// sdkGoResource returns the resource type that req's User-Agent names with a
// resource/<name> segment, if s selects it as sdk_go. Data sources, whose User-Agent also
// has data/yes, stay on the existing clients.
func (s sdkImplementations) sdkGoResource(req *http.Request) (string, bool) {
	segments := strings.Fields(req.UserAgent())
	if slices.Contains(segments, "data/yes") {
		return "", false
	}
	for name, implementation := range s {
		resource := "resource/" + strings.TrimPrefix(name, "databricks_")
		if implementation == sdkGoImplementation && slices.Contains(segments, resource) {
			return name, true
		}
	}
	return "", false
}

// checkUserAgent checks the sdk-go module segment, such as sdk-go-uc-volumes/0.8.0, in a
// User-Agent that names a resource with a resource/<name> segment. A read that s sends
// through sdk-go must carry one, and only a resource that s selects as sdk_go may, for
// example on the requests that follow a redirect. sdk-go requests are otherwise the
// existing clients' requests, so only this shows which one ran. Requests that name no
// resource, such as the importer's read, aren't checked.
func (s sdkImplementations) checkUserAgent(req *http.Request) error {
	userAgent := req.UserAgent()
	segments := strings.Fields(userAgent)
	if !slices.ContainsFunc(segments, func(segment string) bool {
		return strings.HasPrefix(segment, "resource/")
	}) {
		return nil
	}
	hasModule := slices.ContainsFunc(segments, func(segment string) bool {
		return strings.HasPrefix(segment, "sdk-go-")
	})
	name, sdkGo := s.sdkGoResource(req)
	if hasModule && !sdkGo {
		return fmt.Errorf("%s %s doesn't select sdk_go, but its User-Agent has an "+
			"sdk-go module: %q", req.Method, req.URL.EscapedPath(), userAgent)
	}
	if !hasModule && sdkGo && sentThroughSdkGo(name, req) {
		return fmt.Errorf("%s %s selects sdk_go, but its User-Agent has no sdk-go "+
			"module: %q", req.Method, req.URL.EscapedPath(), userAgent)
	}
	return nil
}

// sentThroughSdkGo reports whether resource sends req through sdk-go.
func sentThroughSdkGo(resource string, req *http.Request) bool {
	return slices.ContainsFunc(sdkGoRequests[resource], func(r sdkGoRequest) bool {
		return req.Method == r.method &&
			strings.HasPrefix(req.URL.EscapedPath(), r.pathPrefix)
	})
}

// userAgentCheck fails a step whose requests break checkUserAgent for the step's
// sdk_implementations. The fixtures still answer those requests.
type userAgentCheck struct {
	*httpfixture.Handler
	// selections holds the sdk_implementations of each step label, and destroy's.
	selections map[string]sdkImplementations
	mu         sync.Mutex
	label      string
	failures   []error
}

func newUserAgentCheck(handler *httpfixture.Handler, d *testDefinition) *userAgentCheck {
	selections := map[string]sdkImplementations{}
	for i, s := range d.Steps {
		selections[stepLabel(i, s.Name)] = s.SdkImplementations
	}
	// Destroy uses the selection retained with its cleanup configuration.
	retained := d.cleanupConfigSteps()
	if len(retained) > 0 && retained[len(retained)-1] >= 0 {
		cleanup := d.Steps[retained[len(retained)-1]]
		selections[stepLabel(len(d.Steps), "destroy")] = cleanup.SdkImplementations
	}
	return &userAgentCheck{Handler: handler, selections: selections}
}

// StartStep reports the previous step's failures, like httpfixture.Handler.StartStep.
func (c *userAgentCheck) StartStep(label string, calls []httpfixture.Call) error {
	c.mu.Lock()
	failures := c.failures
	c.label, c.failures = label, nil
	c.mu.Unlock()
	return errors.Join(append(failures, c.Handler.StartStep(label, calls))...)
}

// Verify reports the current step's failures without clearing them, like
// httpfixture.Handler.Verify.
func (c *userAgentCheck) Verify() error {
	c.mu.Lock()
	failures := slices.Clone(c.failures)
	c.mu.Unlock()
	return errors.Join(append(failures, c.Handler.Verify())...)
}

func (c *userAgentCheck) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c.mu.Lock()
	if err := c.selections[c.label].checkUserAgent(req); err != nil {
		c.failures = append(c.failures, fmt.Errorf("%s: %w", c.label, err))
	}
	c.mu.Unlock()
	c.Handler.ServeHTTP(w, req)
}
