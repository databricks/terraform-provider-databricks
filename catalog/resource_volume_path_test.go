package catalog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// The expected values are what sdk-go uc/volumes/v1 puts in the request path.
func TestVolumeFullNameInPath(t *testing.T) {
	for name, want := range map[string]string{
		"main.replay.raw":      "main.replay.raw",
		"main.replay.raw#2":    "main.replay.raw%232",
		"main.replay.raw?2":    "main.replay.raw%3F2",
		"main.replay.raw%41":   "main.replay.raw%2541",
		"main.replay.raw%zz":   "main.replay.raw%25zz",
		"main.replay.a(b);c,d": "main.replay.a(b);c,d",
		"main.replay.a(b)#c":   "main.replay.a%28b%29%23c",
		"main.replay.résum":    "main.replay.r%C3%A9sum",
	} {
		assert.Equal(t, want, volumeFullNameInPath(name), name)
	}
}
