package option

import (
	"fmt"
	"log"
	"runtime/debug"

	"dam/driver/logger/color"
)

var Config Conf

type Conf struct {
	DB             DB
	FilesDB        FilesDB
	Decoration     Decoration
	DefaultRepo    DefaultRepo
	Docker         Docker
	Export         Export
	FileSystem     FileSystem
	Global         Global
	Multiversion   Multiversion
	OfficialRepo   OfficialRepo
	ReservedEnvs   ReservedEnvs
	Save           Save
	Search         Search
	Sort           Sort
	Virtualization Virtualization
}

func printFatal(message string, args ...interface{}) {
	debug.SetTraceback("")

	message = "CONFIG ERROR: " + message
	if Config.Decoration.GetColorOn() {
		message = color.Red + message + color.Reset
	}

	if len(args) != 0 {
		message = fmt.Sprintf(message, args...)
	}

	log.Println(message)
	panic(nil)
}
