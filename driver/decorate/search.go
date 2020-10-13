package decorate

import "fmt"

func PrintSearchedApp(app string) {
	fmt.Print(app + ":")
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