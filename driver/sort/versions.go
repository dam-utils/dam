package sort

import (
	"sort"

	"github.com/Masterminds/semver/v3"
)

func SemanticVersionSort(vers *[]string) *[]string {
	var vs []*semver.Version
	var notSemanticVersions []string
	for _, r := range *vers  {
		v, err := semver.NewVersion(r)
		if err != nil {
			notSemanticVersions = append(notSemanticVersions, r)
			continue
		}
		vs = append(vs, v)
	}
	var results []string
	if len(vs) != 0 {
		sort.Sort(semver.Collection(vs))
		for _, v := range vs {
			if v != nil {
				results = append(results, v.String())
			}
		}
	}
	if len(notSemanticVersions) != 0 {
		sortedNoSemVer := AlphabeticSort(&notSemanticVersions)
		results = append(results, *sortedNoSemVer...)
	}
	return &results
}
