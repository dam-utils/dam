package decorate

import (
	"dam/driver/logger/color"
	"dam/driver/structures"
	"fmt"
	"strconv"

	"dam/config"
	"dam/driver/db"
)

var defAppColumnSize = map[string]int {
	"Name" : 4, 	// App.ImageName
	"Version" : 7,	// App.ImageVersion
	"Repository" : 10,	// App.RepoID
}

func PrintRAWAppsList() {
	apps := db.ADriver.GetApps()
	for _, app := range apps {
		var multiVers string
		if app.MultiVersion {
			multiVers = config.DECORATE_BOOL_FLAG
		} else {
			multiVers = ""
		}
		fields := []string{strconv.Itoa(app.Id), app.DockerID, app.ImageName, app.ImageVersion, strconv.Itoa(app.RepoID), multiVers, app.Family}
		printRAWStr(fields)
	}
}

func PrintAppsList(){
	apps := db.ADriver.GetApps()

	fmt.Println()
	fmt.Println("\tList apps:")
	fmt.Println()

	prepareAppsColumnSize(apps)
	// general field size
	fieldSize := (config.DECORATE_MAX_DISPLAY_WIDTH - len(ColumnSeparator)*(len(defAppColumnSize)-1))/len(defAppColumnSize)
	if len(apps) != 0 {
		printAppsTitle(fieldSize)
		printAppsLineSeparator(fieldSize)
		for _, app := range apps {
			printApp(app, fieldSize)
		}
		fmt.Println()
	}
}

func prepareAppsColumnSize(apps []*structures.App){
	for _, app := range apps {
		if param := checkStrFieldSize(app.ImageName); param > defAppColumnSize["Name"] {
			defAppColumnSize["Name"] = param
		}
		if param := checkStrFieldSize(app.ImageVersion); param > defAppColumnSize["Version"] {
			defAppColumnSize["Version"] = param
		}
		if param := checkStrFieldSize(getRepoNameByApp(app)); param > defAppColumnSize["Repository"] {
			defAppColumnSize["Repository"] = param
		}
	}
}

func printAppsTitle(fsize int) {
	printTitleField("Name", fsize, defAppColumnSize)
	fmt.Print(ColumnSeparator)
	printTitleField("Version", fsize, defAppColumnSize)
	fmt.Print(ColumnSeparator)
	printTitleField("Repository", fsize, defAppColumnSize)
	fmt.Println()
}

func getRepoNameByApp(app *structures.App) string {
	name := config.UNKNOWN_REPO_NAME
	repo := db.RDriver.GetRepoById(app.RepoID)
	if repo != nil {
		name = repo.Name
	}
	return name
}

func printAppsLineSeparator(fieldSize int) {
	for _, value := range defAppColumnSize {
		if value < fieldSize {
			fmt.Print(getPreparedSeparator(value, LineSeparator))
		} else {
			fmt.Print(getPreparedSeparator(fieldSize, LineSeparator))
		}
	}
	for i := 0; i < len(defAppColumnSize)-1; i++ {
		fmt.Print(getPreparedSeparator(len(ColumnSeparator), LineSeparator))
	}
	fmt.Println()
}

func printApp(app *structures.App, limitSize int) {
	printField(app.ImageName, limitSize, defAppColumnSize["Name"])
	fmt.Print(ColumnSeparator)
	printField(app.ImageVersion, limitSize, defAppColumnSize["Version"])
	fmt.Print(ColumnSeparator)
	printField(getRepoNameByApp(app), limitSize, defAppColumnSize["Repository"])
	fmt.Println()
}

func PrintAppList(title string, appSkipList []*structures.ImportApp, c string) {
	if len(appSkipList) == 0 {
		return
	}

	fmt.Println(c + title)

	for _, app := range appSkipList {
		fmt.Println("\t"+ app.CurrentName())
	}

	fmt.Println(color.Reset)
}