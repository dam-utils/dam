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
	"sort"
	"strings"
)

type Alphabetic []string
func (list Alphabetic) Len() int      { return len(list) }
func (list Alphabetic) Swap(i, j int) { list[i], list[j] = list[j], list[i] }
func (list Alphabetic) Less(i, j int) bool {
	var si string = list[i]
	var sj string = list[j]
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