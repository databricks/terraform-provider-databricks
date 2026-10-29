package terraformtest

import "testing"

func TestRecordedReplay(t *testing.T) {
	Run(t, "testdata/recording")
}
