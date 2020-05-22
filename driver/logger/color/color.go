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
package color

//from https://golangbyexample.com/print-output-text-color-console/
const (
	Reset = string("\033[0m")

	Red    = string("\033[31m")
	Green  = string("\033[32m")
	Yellow = string("\033[33m")
	Blue   = string("\033[34m")
	Purple = string("\033[35m")
	Cyan   = string("\033[36m")
	White  = string("\033[37m")
)