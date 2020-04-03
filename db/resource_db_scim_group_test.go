package db

import "testing"

func diff(sliceA []string, sliceB []string) []string {
	var output []string
	m := make(map[string]int)
	for _, y := range sliceB {
		m[y]++
	}
	for _, x := range sliceA {
		if m[x] > 0 {
			m[x]--
			continue
		}
		output = append(output, x)
	}
	return output
}

func TestDiff(t *testing.T) {
	tf_main := []string{"a", "b", "c"}
	remote := []string{"b", "d"}

	t.Log(diff(tf_main, remote))
	t.Log(diff(remote, tf_main))

}
