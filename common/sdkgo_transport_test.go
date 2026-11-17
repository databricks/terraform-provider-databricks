package common

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"sync"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// redirectTo answers with a 307 to the same path on target.
func redirectTo(target *httptest.Server) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, target.URL+r.URL.RequestURI(), http.StatusTemporaryRedirect)
	}
}

func TestSdkGoScimClientFollowsRedirectsLikeClassicClient(t *testing.T) {
	read := func(sdkGo bool) (originCalls, destinationCalls int32, err error) {
		destination, destinationRequests := serveSequence(t,
			func(w http.ResponseWriter, r *http.Request) {
				path := "/api/2.0/preview/scim/v2/ServicePrincipals/201"
				assert.Equal(t, path, r.URL.Path)
				respond(http.StatusOK, sdkGoTestServicePrincipal)(w, r)
			})
		origin, originRequests := serveSequence(t, redirectTo(destination))
		c := newSdkGoTestClient(t, origin, nil)
		if sdkGo {
			_, err = getServicePrincipal(context.Background(), t, c, "")
		} else {
			err = classicGetServicePrincipal(context.Background(), c)
		}
		return originRequests.Load(), destinationRequests.Load(), err
	}
	originCalls, destinationCalls, err := read(true)
	require.NoError(t, err)
	assert.EqualValues(t, 1, originCalls)
	assert.EqualValues(t, 1, destinationCalls)
	classicOriginCalls, classicDestinationCalls, classicErr := read(false)
	require.NoError(t, classicErr)
	assert.Equal(t, classicOriginCalls, originCalls)
	assert.Equal(t, classicDestinationCalls, destinationCalls)
}

func TestSdkGoScimClientReportsPrivateLinkRedirectsLikeClassicClient(t *testing.T) {
	toLogin := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login.html?error=private-link-validation-error",
			http.StatusFound)
	}
	loginPage := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/login.html", r.URL.Path)
		w.Header().Set("Content-Type", "text/html")
		respond(http.StatusOK, "<html>login</html>")(w, r)
	}
	r := readBothWays(t, context.Background, nil, toLogin, loginPage)
	require.Error(t, r.err)
	assert.EqualError(t, r.classicErr, r.err.Error())
	assert.ErrorIs(t, r.err, apierr.ErrPermissionDenied)
	var apiErr *apierr.APIError
	require.ErrorAs(t, r.err, &apiErr)
	assert.Equal(t, "PRIVATE_LINK_VALIDATION_ERROR", apiErr.ErrorCode)
	assert.False(t, apierr.IsMissing(r.err))
	assert.EqualValues(t, 2, r.calls)
	assert.Equal(t, r.classicCalls, r.calls)
}

// capturedLogger records the classic SDK's log lines at every level.
type capturedLogger struct {
	mu    sync.Mutex
	lines []string
}

func (l *capturedLogger) Enabled(context.Context, logger.Level) bool { return true }

func (l *capturedLogger) record(level, format string, v ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lines = append(l.lines, level+" "+fmt.Sprintf(format, v...))
}

func (l *capturedLogger) Tracef(_ context.Context, format string, v ...any) {
	l.record("TRACE", format, v...)
}

func (l *capturedLogger) Debugf(_ context.Context, format string, v ...any) {
	l.record("DEBUG", format, v...)
}

func (l *capturedLogger) Infof(_ context.Context, format string, v ...any) {
	l.record("INFO", format, v...)
}

func (l *capturedLogger) Warnf(_ context.Context, format string, v ...any) {
	l.record("WARN", format, v...)
}

func (l *capturedLogger) Errorf(_ context.Context, format string, v ...any) {
	l.record("ERROR", format, v...)
}

// matching returns the captured lines that contain every one of parts.
func (l *capturedLogger) matching(parts ...string) []string {
	var matched []string
	for _, line := range l.lines {
		missing := func(part string) bool { return !strings.Contains(line, part) }
		if !slices.ContainsFunc(parts, missing) {
			matched = append(matched, line)
		}
	}
	return matched
}

func TestSdkGoScimClientLogsLikeClassicClient(t *testing.T) {
	sdkGoLogger, classicLogger := &capturedLogger{}, &capturedLogger{}
	// readBothWays reads through sdk-go first.
	next := []*capturedLogger{sdkGoLogger, classicLogger}
	newContext := func() context.Context {
		l := next[0]
		next = next[1:]
		return logger.NewContext(context.Background(), l)
	}
	r := readBothWays(t, newContext, nil,
		respond(429, `{"detail": "Too many requests.", "status": "429"}`),
		respond(http.StatusOK, sdkGoTestServicePrincipal))
	require.NoError(t, r.err)
	require.NoError(t, r.classicErr)
	path := "/api/2.0/preview/scim/v2/ServicePrincipals/201"
	for name, l := range map[string]*capturedLogger{"sdk-go": sdkGoLogger,
		"classic": classicLogger} {
		assert.Len(t, l.matching("DEBUG GET "+path, "< HTTP/1.1 429 Too Many Requests",
			"Too many requests."), 1, name)
		assert.Len(t, l.matching("DEBUG GET "+path, "< HTTP/1.1 200 OK", "app-201"), 1,
			name)
		assert.Len(t, l.matching("TRACE Too many requests. Sleeping"), 1, name)
	}
}
