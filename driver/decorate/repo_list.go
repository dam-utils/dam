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
	"strconv"

	"dam/config"
	"dam/driver/db"
	"dam/driver/storage"
)

func PrintRAWReposList() {
	repos := db.RDriver.GetRepos()
	for _, repo := range *repos {
		var def string
		if repo.Default {
			def = config.DECORATE_BOOL_FLAG
		} else {
			def = ""
		}
		fields := []string{strconv.Itoa(repo.Id), def, repo.Name, repo.Server, repo.Username}
		printRAWStr(fields)
	}
}

var defRepoColumnSize = map[string]int {
	"Num" : 3, 	// Repo.Id
	"Def" : 3,	// Repo.Default
	"Name" : 4,	// Repo.Name
	"Url" : 3,	// Repo.Server
	"User" : 4, // Repo.Username
}

func PrintReposList(){
	repos := db.RDriver.GetRepos()

	fmt.Println()
	fmt.Println("\tList repositories:")
	fmt.Println()

	prepareReposColumnSize(repos)
	// general field size
	fieldSize := (config.DECORATE_MAX_DISPLAY_WIDTH - len(ColumnSeparator)*(len(defRepoColumnSize)-1))/len(defRepoColumnSize)
	if len(*repos) != 0 {
		printReposTitle(fieldSize)
		printReposLineSeparator(fieldSize)
		for _, repo := range *repos {
			printRepo(&repo, fieldSize)
		}
		fmt.Println()
	}
}

func prepareReposColumnSize(repos *[]storage.Repo){
	for _, repo := range *repos {
		if param := checkIntFieldSize(repo.Id); param > defRepoColumnSize["Num"] {
			defRepoColumnSize["Num"] = param
		}
		// defColumnSize["Def"] = 3
		if param := checkStrFieldSize(repo.Name); param > defRepoColumnSize["Name"] {
			defRepoColumnSize["Name"] = param
		}
		if param := checkStrFieldSize(repo.Server); param > defRepoColumnSize["Url"] {
			defRepoColumnSize["Url"] = param
		}
		if param := checkStrFieldSize(repo.Username); param > defRepoColumnSize["User"] {
			defRepoColumnSize["User"] = param
		}
	}
}

func printReposTitle(fsize int) {
	for _, str := range [...]string{"Num", "Def", "Name", "Url"} {
		printTitleField(str, fsize, defRepoColumnSize)
		fmt.Print(ColumnSeparator)
	}
	printTitleField("User", fsize, defRepoColumnSize)
	fmt.Println()
}

func printReposLineSeparator(fieldSize int) {
	for _, value := range defRepoColumnSize {
		if value < fieldSize {
			fmt.Print(getPreparedSeparator(value, LineSeparator))
		} else {
			fmt.Print(getPreparedSeparator(fieldSize, LineSeparator))
		}
	}
	for i := 0; i < len(defRepoColumnSize)-1; i++ {
		fmt.Print(getPreparedSeparator(len(ColumnSeparator), LineSeparator))
	}
	fmt.Println()
}

func printRepo(repo *storage.Repo, limitSize int) {
	printField(strconv.Itoa(repo.Id), limitSize, defRepoColumnSize["Num"])
	fmt.Print(ColumnSeparator)
	printField(bool2Str(repo.Default), limitSize, defRepoColumnSize["Def"])
	fmt.Print(ColumnSeparator)
	printField(repo.Name, limitSize, defRepoColumnSize["Name"])
	fmt.Print(ColumnSeparator)
	printField(repo.Server, limitSize, defRepoColumnSize["Url"])
	fmt.Print(ColumnSeparator)
	printField(repo.Username, limitSize, defRepoColumnSize["User"])
	fmt.Println()
}

