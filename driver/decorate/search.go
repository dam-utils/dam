package decorate

import (
	"dam/driver/conf/option"
	"dam/driver/logger/color"
	"fmt"
)

func PrintSearchedApp(app string) {
	if option.Config.Decoration.GetColorOn() {
		fmt.Print(color.Yellow + app + color.Reset + ":")
	} else {
		fmt.Print(app + ":")
	}
}

func PrintSearchedVersions(list []string){
	lenList := len(list)
	for i, str := range list {
		if i == lenList -1 {
			fmt.Println(str)
		} else {
			fmt.Print(str+", ")
		}
	}
}