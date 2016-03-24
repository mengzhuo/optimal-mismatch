package om

import (
	"sort"
)

type Pat struct {
	loc  int
	char byte
}

type StringFinder struct {
	pattern       string
	patternLength int
	pats          []*Pat
	qsBc, freq    [256]int
	adaptedGs     []int
}

func MakeStringFinder(pattern string) *StringFinder {

	l := len(pattern)

	f := &StringFinder{
		pattern:       pattern,
		patternLength: l,
	}

	return f
}

func (f *StringFinder) orderPattern() {

	f.pats = make([]*Pat, f.patternLength)
	for i := 0; i < f.patternLength; i++ {
		f.pats[i] = &Pat{i, f.pattern[i]}
	}
}
