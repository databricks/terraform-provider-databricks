package acceptance

import (
	"net/http/httptest"
	"os"
	"sync"

	"github.com/databricks/cli/libs/fakebricks"
)

// fakeServerEnvVar, when set to a non-empty value, routes the acceptance tests at
// an in-process fakebricks backend instead of a real workspace. This lets the
// existing integration tests run with no cloud credentials.
const fakeServerEnvVar = "DATABRICKS_FAKE"

// fakeToken is the bearer token every fake-mode client authenticates with. A
// single token binds every client to one shared workspace.
const fakeToken = "fake-token"

var fakeOnce sync.Once

// FakeServerEnabled reports whether tests should run against the in-process fake.
func FakeServerEnabled() bool {
	return os.Getenv(fakeServerEnvVar) != ""
}

// startFakeServerIfEnabled boots a single process-wide fakebricks backend behind
// a real httptest socket and points the Databricks SDK at it via the standard
// DATABRICKS_HOST / DATABRICKS_TOKEN environment variables.
//
// Driving the integration through env vars (rather than an injected HTTP
// transport) is what makes it generic: every Databricks client built during a
// test — the provider's own client, and the auxiliary clients that Check
// callbacks construct via client.New — resolves host and token from the
// environment, so they all reach the same backend and share its state.
//
// A single backend is shared across the whole test process so that state created
// in one step (e.g. a resource create) remains visible to later steps and to
// Check callbacks. Because all clients authenticate with the same token they bind
// to the same fake workspace; tests within a package therefore share state and
// should be run individually or in curated groups.
func startFakeServerIfEnabled() {
	if !FakeServerEnabled() {
		return
	}
	fakeOnce.Do(func() {
		fake := fakebricks.New()
		server := httptest.NewServer(fake.Handler())
		// Responses that echo the workspace host (e.g. a job run's run_page_url)
		// must match the URL clients actually reached the fake at. SetHost must be
		// called before any token binds a workspace.
		fake.SetHost(server.URL)
		// Bind the shared token to a workspace with the default user's home
		// directory seeded, as a real workspace always has it. This also pins all
		// clients (provider and Check callbacks) that authenticate with this token
		// to one shared workspace, instead of lazily minting a fresh one per token.
		fakebricks.BindWorkspace(fake, fakeToken, fakebricks.DefaultUser)
		os.Setenv("DATABRICKS_HOST", server.URL)
		os.Setenv("DATABRICKS_TOKEN", fakeToken)
		os.Setenv("DATABRICKS_AUTH_TYPE", "pat")
	})
}
