package decorate

import (
	"bufio"
	"dam/driver/conf/option"
	fs "dam/driver/filesystem"
	"dam/driver/logger"
	"dam/driver/logger/color"
	"fmt"
	"os"
)

func PrintDescription(desc string) {
	if fs.IsExistFile(desc) {
		f, err := os.Open(desc)
		defer func() {
			if f != nil {
				f.Close()
			}
		}()
		if err != nil {
			logger.Fatal("Cannot open file '%s' with error: %s", desc, err)
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			fmt.Println(color.Yellow + scanner.Text() + color.Reset)
		}
		if err := scanner.Err(); err != nil {
			logger.Warn("Cannot read full description with error: %s", err)
		}
	} else {
		logger.Warn("Not found %s file with the app description", option.Config.FileSystem.GetDescriptionFileName())
	}
}

func Println() {
	fmt.Println()
}

func PrintLabel(name, s string) {
	fmt.Printf("Label %s: '%s'\n", name, s)
}