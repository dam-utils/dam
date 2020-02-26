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
	"dam/config"
	d_log "dam/decorate/log"
)

func SortAppNames(names *[]string) *[]string {
	switch config.SORT_APP_TYPE {
	case "alphabetic":
		return AlphabeticSort(names)
	default:
		d_log.Fatal("Config option SORT_APP_TYPE='"+config.SORT_APP_TYPE+"' not valid. Sorting type is bad." )
	}
	return nil
}

func SortVersions(vers *[]string) *[]string {
	switch config.SORT_VERSION_TYPE {
	case "semantic_version":
		return SemanticVersionSort(vers)
	default:
		d_log.Fatal("Config option SORT_VERSION_TYPE='"+config.SORT_VERSION_TYPE+"' not valid. Sorting type is bad." )
	}
	return nil
}