package databricks

import (
	"testing"
)

//go:generate easytags $GOFILE

func TestCheckers(t *testing.T) {
	t.Log(fetchStringFromCheckers([]string{"test", "abc"}))
}
