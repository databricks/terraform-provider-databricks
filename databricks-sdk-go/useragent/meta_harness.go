package useragent

import (
	"os"
	"sync"
)

// knownMetaHarness describes a meta-harness that orchestrates AI agents (not
// itself an agent). Detected if envVar is set (any value, including empty, counts).
type knownMetaHarness struct {
	envVar  string
	product string
}

// listKnownMetaHarnesses returns the canonical list of agent meta-harnesses.
// Keep in sync with databricks-sdk-py and databricks-sdk-java.
func listKnownMetaHarnesses() []knownMetaHarness {
	return []knownMetaHarness{
		{envVar: "OMNIGENT", product: "omnigent"}, // https://github.com/omnigent-ai/omnigent
	}
}

// lookupMetaHarnessProvider returns the matched meta-harness product, "multiple"
// if more than one matched, or "" if none matched.
func lookupMetaHarnessProvider() string {
	var matches []string
	for _, h := range listKnownMetaHarnesses() {
		if _, ok := os.LookupEnv(h.envVar); ok {
			matches = append(matches, h.product)
		}
	}
	switch len(matches) {
	case 1:
		return matches[0]
	case 0:
		return ""
	default:
		return "multiple"
	}
}

var (
	metaHarnessOnce sync.Once
	metaHarnessName string
)

// MetaHarnessProvider returns the detected meta-harness name, cached for the process lifetime.
func MetaHarnessProvider() string {
	metaHarnessOnce.Do(func() {
		metaHarnessName = lookupMetaHarnessProvider()
	})
	return metaHarnessName
}
