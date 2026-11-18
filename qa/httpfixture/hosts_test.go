package httpfixture

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var testHosts = map[string]string{"account": "accounts.example.com", "workspace": "workspace.example.com:8443"}

func sendTo(h *Handler, host, method, path string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(""))
	req.Host = host
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	return w
}

func hostCall(host, method, path, result string) Call {
	c := call(method, path, nil, result)
	c.Request.Host = host
	return c
}

func TestCallsMatchTheirNamedHost(t *testing.T) {
	h := NewHandler()
	require.NoError(t, h.SetHosts(testHosts, "account"))
	require.NoError(t, h.SetReusableFixtures([]Call{
		hostCall("", "GET", "/metadata", "account metadata"),
		hostCall("workspace", "GET", "/metadata", "workspace metadata"),
	}))
	require.NoError(t, h.StartStep("step 1 (create)", []Call{
		hostCall("", "GET", "/lookup", "lookup"),
		hostCall("workspace", "POST", "/items", "created"),
		hostCall("account", "GET", "/items", "account read"),
	}))
	require.Equal(t, `"workspace metadata"`, sendTo(h, "workspace.example.com:8443", "GET", "/metadata").Body.String())
	require.Equal(t, `"account metadata"`, sendTo(h, "accounts.example.com", "GET", "/metadata").Body.String())
	// Calls without a host expect the default host, not every host.
	require.Equal(t, 400, sendTo(h, "workspace.example.com:8443", "GET", "/lookup").Code)
	require.Equal(t, `"lookup"`, sendTo(h, "accounts.example.com", "GET", "/lookup").Body.String())
	require.Equal(t, 400, sendTo(h, "accounts.example.com", "POST", "/items").Code)
	require.Equal(t, `"created"`, sendTo(h, "workspace.example.com:8443", "POST", "/items").Body.String())
	require.Equal(t, 400, sendTo(h, "127.0.0.1:1234", "GET", "/items").Code)
	require.Equal(t, `"account read"`, sendTo(h, "accounts.example.com", "GET", "/items").Body.String())
	err := h.Verify()
	require.ErrorContains(t, err, `step 1 (create): unexpected HTTP call: GET /lookup on host "workspace" body=`)
	require.ErrorContains(t, err, `step 1 (create): unexpected HTTP call: POST /items on host "account" body=`)
	// A request to an unnamed host names no host.
	require.ErrorContains(t, err, "step 1 (create): unexpected HTTP call: GET /items body=")
	require.NotContains(t, err.Error(), "unconsumed")
}

func TestCallsWithoutHostMatchUnnamedHostsWithoutDefault(t *testing.T) {
	h := NewHandler()
	require.NoError(t, h.SetHosts(testHosts, ""))
	require.NoError(t, h.StartStep("read", []Call{call("GET", "/items", nil, "local"), hostCall("workspace", "GET", "/items", "workspace")}))
	require.Equal(t, `"workspace"`, sendTo(h, "workspace.example.com:8443", "GET", "/items").Body.String())
	require.Equal(t, 400, sendTo(h, "accounts.example.com", "GET", "/items").Code)
	require.Equal(t, `"local"`, sendTo(h, "127.0.0.1:1234", "GET", "/items").Body.String())
	require.ErrorContains(t, h.Verify(), `read: unexpected HTTP call: GET /items on host "account"`)
}

func TestUnconsumedAndOutOfOrderCallsNameTheirHost(t *testing.T) {
	h := NewHandler()
	h.SetEnforceCallOrder(true)
	require.NoError(t, h.SetHosts(testHosts, "account"))
	require.NoError(t, h.StartStep("create", []Call{hostCall("", "GET", "/lookup", "lookup"), hostCall("workspace", "POST", "/items", "created")}))
	response := sendTo(h, "workspace.example.com:8443", "POST", "/items")
	require.Equal(t, 400, response.Code)
	require.Contains(t, response.Body.String(), `out-of-order HTTP call: POST /items on host "workspace" body= matched call 2; next expected call 1: GET /lookup on host "account"`)
	err := h.Verify()
	require.ErrorContains(t, err, `create: unconsumed call 1: GET /lookup on host "account"`)
	require.ErrorContains(t, err, `create: unconsumed call 2: POST /items on host "workspace"`)
}

func TestRejectInvalidHostsWithoutReplacingThem(t *testing.T) {
	for _, tc := range []struct {
		name        string
		hosts       map[string]string
		defaultHost string
		message     string
	}{
		{"empty name", map[string]string{"": "accounts.example.com"}, "", `host "": expected a nonempty name and authority`},
		{"empty authority", map[string]string{"account": ""}, "", `host "account": expected a nonempty name and authority`},
		{"shared authority", map[string]string{"account": "example.com", "workspace": "example.com"}, "", `hosts "account" and "workspace" have the same authority example.com`},
		{"unknown default", testHosts, "workspac", `default host "workspac" is not a named host`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			h := NewHandler()
			require.NoError(t, h.SetHosts(testHosts, "account"))
			require.ErrorContains(t, h.SetHosts(tc.hosts, tc.defaultHost), tc.message)
			require.NoError(t, h.StartStep("read", []Call{hostCall("workspace", "GET", "/items", "workspace")}))
			require.Equal(t, 200, sendTo(h, "workspace.example.com:8443", "GET", "/items").Code)
		})
	}
}

func TestRejectCallsNamingUnknownHosts(t *testing.T) {
	h := NewHandler()
	require.NoError(t, h.SetHosts(testHosts, "account"))
	require.NoError(t, h.StartStep("original", []Call{call("GET", "/original", nil, "original")}))
	invalid := hostCall("workspac", "GET", "/items", "invalid")
	require.EqualError(t, h.StartStep("step 2 (update)", []Call{call("GET", "/other", nil, "other"), invalid}), `step 2 (update): call 2: host "workspac" is not a named host`)
	require.ErrorContains(t, h.SetReusableFixtures([]Call{invalid}), `call 1: host "workspac" is not a named host`)
	require.ErrorContains(t, NewHandler().StartStep("no hosts", []Call{hostCall("account", "GET", "/items", "")}), `host "account" is not a named host`)
	require.Equal(t, `"original"`, sendTo(h, "accounts.example.com", "GET", "/original").Body.String())
	require.NoError(t, h.Verify())
}
