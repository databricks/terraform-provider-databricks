package db

import "testing"

func TestDiff(t *testing.T) {
	tf_main := []string{"a", "b", "c"}
	remote := []string{"b", "d"}

	t.Log(diff(tf_main, remote))
	t.Log(diff(remote, tf_main))

}
