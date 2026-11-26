package useragent

import (
	"os"
	"sync"
)

// Clear cached user agent values like the DBR version or the CI/CD provider
// being used. This is useful for testing.
func ClearCache() {
	// Reset the sync.Once to their default values. This will recompute the
	// values on the next call to Runtime(), CiCdProvider(), or AgentProvider().
	runtimeOnce = sync.Once{}
	providerOnce = sync.Once{}
	agentOnce = sync.Once{}
	metaHarnessOnce = sync.Once{}
}

var runtimeOnce sync.Once
var runtimeVersion string

func getRuntimeVersion() string {
	v := os.Getenv("DATABRICKS_RUNTIME_VERSION")
	return Sanitize(v)
}

func Runtime() string {
	runtimeOnce.Do(func() {
		runtimeVersion = getRuntimeVersion()
	})
	return runtimeVersion
}
