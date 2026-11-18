package terraformtest

import (
	"testing"
)

func TestHostsReplay(t *testing.T) {
	Run(t, "testdata/hosts")
}
