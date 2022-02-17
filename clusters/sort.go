package clusters

import (
	"sort"
)

// readable chained sorting helper
func sortByChain(s interface{}, fn func(int) sortCmp) {
	sort.Slice(s, func(i, j int) bool {
		return fn(i).Less(fn(j))
	})
}

type sortCmp interface {
	Less(o sortCmp) bool
}

type boolAsc bool

func (b boolAsc) Less(o sortCmp) bool {
	return bool(b) != bool(o.(boolAsc)) && !bool(b)
}

type intAsc int

func (ia intAsc) Less(o sortCmp) bool {
	return int(ia) < int(o.(intAsc))
}

type strAsc string

func (s strAsc) Less(o sortCmp) bool {
	return string(s) < string(o.(strAsc))
}

type sortChain []sortCmp

func (c sortChain) Less(other sortCmp) bool {
	o := other.(sortChain)
	for i := range c {
		if c[i].Less(o[i]) {
			return true
		}
		if o[i].Less(c[i]) {
			break
		}
	}
	return false
}
