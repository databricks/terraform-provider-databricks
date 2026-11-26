package useragent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserAgentRuntime(t *testing.T) {
	for _, tc := range []string{
		"1.2.3",
		"0.0.0-dev+2e014739024a",
		"client.0",
		"foo",
		"15.0",
		"13.3",
	} {
		t.Run(tc, func(t *testing.T) {
			t.Setenv("DATABRICKS_RUNTIME_VERSION", tc)

			v := getRuntimeVersion()
			assert.Equal(t, tc, v)
			assert.NoError(t, matchAlphanumOrSemVer(v))
		})
	}
}
