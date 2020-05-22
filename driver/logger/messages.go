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
package logger

import (
	"log"
	"os"

	"dam/config"
	"dam/driver/logger/color"
)

var DebugMode bool

func Fatal(message string, args ...interface{}) {
	message = "ERROR: " + message
	if config.COLLOR_ON == true {
		message = color.Red + message + color.Reset
	}

	if len(args) == 0 {
		log.Println(message)
	} else {
		log.Printf(message, args...)
	}
	os.Exit(1)
}

func Debug(message string, args ...interface{}) {
	message = "DEBUG: "+message
	if config.COLLOR_ON == true {
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
	if config.COLLOR_ON == true {
		message = color.Yellow + message + color.Reset
	}

	if len(args) == 0 {
		log.Println(message)
	} else {
		log.Printf(message, args...)
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
	if config.COLLOR_ON == true {
		message = color.Green + message + color.Reset
	}

	if len(args) == 0 {
		log.Println(message)
	} else {
		log.Printf(message, args...)
	}
}