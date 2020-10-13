package sort

import (
	"sort"
	"strings"
)

type Alphabetic []string
func (list Alphabetic) Len() int      { return len(list) }
func (list Alphabetic) Swap(i, j int) { list[i], list[j] = list[j], list[i] }
func (list Alphabetic) Less(i, j int) bool {
	var si = list[i]
	var sj = list[j]
	var si_lower = strings.ToLower(si)
	var sj_lower = strings.ToLower(sj)
	if si_lower == sj_lower {
		return si < sj
	}
	return si_lower < sj_lower
}

func AlphabeticSort(names *[]string) *[]string {
	var sortNames = *names
	sort.Sort(Alphabetic(sortNames))
	return &sortNames
}