package terraformtest

import (
	"testing"
)

func TestAuthReplay(t *testing.T) {
	Run(t, "testdata/auth")
}
