package decorate

import (
	"fmt"
	"strconv"

	"dam/driver/conf/option"
	"dam/driver/db"
	"dam/driver/structures"
)

func PrintRAWReposList() {
	repos := db.RDriver.GetRepos()
	for _, repo := range repos {
		var def string
		if repo.Default {
			def = option.Config.Decoration.GetBoolFlagSymbol()
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
	fieldSize := (option.Config.Decoration.GetMaxDisplayWidth() - len(ColumnSeparator)*(len(defRepoColumnSize)-1))/len(defRepoColumnSize)
	if len(repos) != 0 {
		printReposTitle(fieldSize)
		printReposLineSeparator(fieldSize)
		for _, repo := range repos {
			printRepo(repo, fieldSize)
		}
		fmt.Println()
	}
}

func prepareReposColumnSize(repos []*structures.Repo){
	for _, repo := range repos {
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

func printRepo(repo *structures.Repo, limitSize int) {
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

