package acceptance

import (
	"testing"
)

type ResourceFixturePluginFramework struct {
	Steps []Step
}

func (f ResourceFixturePluginFramework) RunUnitTest(t *testing.T) error {
	setDebugLogger()
	return runWithFixtureServer(t, f)
}
