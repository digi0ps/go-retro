package logger

import (
	"fmt"
	"time"
)

func getTime() string {
	return time.Now().Format(time.RFC3339)
}

// Log : Prints a log taking type and content
func Log(logType, content interface{}) {
	_time := getTime()
	fmt.Printf("\n%v [%v] %v", _time, logType, content)
}

// Info : Prints [INFO] log
func Info(content interface{}) {
	Log("INFO", content)
}

// Debug : Print [DEBUG] log
func Debug(content interface{}) {
	Log("DEBUG", content)
}

// Error : Prints [ERROR] log
func Error(err error) {
	Log("ERROR", err.Error())
}
