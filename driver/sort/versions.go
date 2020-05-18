// Copyright 2020 The Docker Applications Manager Authors
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
package sort

import (
	"github.com/Masterminds/semver"
	"sort"
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
