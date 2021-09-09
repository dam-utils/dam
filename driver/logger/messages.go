package logger

import (
	"fmt"
	"log"
	"runtime/debug"

	"dam/driver/conf/option"
	"dam/driver/logger/color"
)

var DebugMode bool

func init(){
	log.SetFlags(0)
}


func Fatal(message string, args ...interface{}) {
	debug.SetTraceback("")

	message = "ERROR: " + message
	if option.Config.Decoration.GetColorOn() {
		message = color.Red + message + color.Reset
	}

	if len(args) != 0 {
		message = fmt.Sprintf(message, args...)
	}

	log.Println(message)
	panic(nil)
}

func Error(message string, args ...interface{}) {
	message = "ERROR: " + message
	if option.Config.Decoration.GetColorOn() {
		message = color.Red + message + color.Reset
	}

	if len(args) == 0 {
		log.Println(message)
	} else {
		log.Printf(message, args...)
	}
}

func Debug(message string, args ...interface{}) {
	message = "DEBUG: "+message
	if option.Config.Decoration.GetColorOn() {
		message = color.White + message + color.Reset
	}

	if DebugMode {
		if len(args) == 0 {
			log.Println(message)
		} else {
			log.Printf(message, args...)
		}
	}
}

func Warn(message string, args ...interface{}) {
	message = "WARN: "+message
	if option.Config.Decoration.GetColorOn() {
		message = color.Yellow + message + color.Reset
	}

	if DebugMode {
		if len(args) == 0 {
			log.Println(message)
		} else {
			log.Printf(message, args...)
		}
	}
}

func Info(message string, args ...interface{}) {
	if len(args) == 0 {
		log.Println(message)
	} else {
		log.Printf(message, args...)
	}
}

func Success(message string, args ...interface{}) {
	if option.Config.Decoration.GetColorOn() {
		message = color.Green + message + color.Reset
	}

	if len(args) == 0 {
		log.Println(message)
	} else {
		log.Printf(message, args...)
	}
}