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
package internal

import (
	"dam/config"
	fs "dam/driver/filesystem"
)

func PrepareTmpMetaPath(meta string) string {
	path := fs.GetAbsolutePath(meta)
	fs.Remove(path)
	return path
}

func BoolToString(b bool) string {
	if b {
		return config.MULTIVERSION_TRUE_FLAG
	}
	return config.MULTIVERSION_FALSE_FLAG
}
