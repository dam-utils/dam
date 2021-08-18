package logger

import (
	"fmt"
	"log"
	"runtime/debug"

	"dam/config"
	"dam/driver/logger/color"
)

var DebugMode bool

func init(){
	log.SetFlags(0)
}


func Fatal(message string, args ...interface{}) {
	debug.SetTraceback("")

	message = "ERROR: " + message
	if config.COLOR_ON == true {
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
	if config.COLOR_ON == true {
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
	if config.COLOR_ON == true {
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
	if config.COLOR_ON == true {
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
	if config.COLOR_ON == true {
		message = color.Green + message + color.Reset
	}

	if len(args) == 0 {
		log.Println(message)
	} else {
		log.Printf(message, args...)
	}
}