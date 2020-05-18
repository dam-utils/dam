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
package decorate

import (
	"fmt"
	"log"
	"strconv"

	"dam/config"
)

const (
	SpaceSeparator = " "
	ColumnSeparator = "  |  "
	LineSeparator = "-")

func printField(str string, limitFieldSize int, defaultFieldSize int) {
	strLen := len(str)

	if limitFieldSize < defaultFieldSize {
		defaultFieldSize = limitFieldSize
	}

	if strLen < limitFieldSize {
		fmt.Print(str + getPreparedSeparator(defaultFieldSize-strLen, SpaceSeparator))
	} else {
		fmt.Print(str[:defaultFieldSize] + "~")
	}
}

func bool2Str(b bool) string {
	if !b {
		return SpaceSeparator
	} else {
		return config.DECORATE_BOOL_FLAG
	}
}

func getPreparedSeparator(i int, sep string) string {
	acc := ""
	for s:=0; s<=i;s++ {
		acc = acc + sep
	}
	return acc
}

func printTitleField(fName string, fSize int, columnMap map[string]int) {
	if columnMap[fName] < fSize {
		fmt.Print(fName + getPreparedSeparator(columnMap[fName]-len(fName), SpaceSeparator))
	} else {
		fmt.Print(fName + getPreparedSeparator(fSize-len(fName), SpaceSeparator))
	}
}

func checkStrFieldSize(s string) int {
	return len(s)
}

func checkIntFieldSize(i int) int {
	return len(strconv.Itoa(i))
}

func printRAWStr(fields []string){
	sep := config.DECORATE_RAW_SEPARATOR
	lenF := len(fields)
	str := ""
	for j, field := range fields {
		if j == lenF -1 {
			str = str + field
		} else {
			str = str + field + sep
		}
	}
	log.Println(str)
}