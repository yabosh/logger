package logger

//
// Standard console logger with support for levels.
//
// This is a wrapper around the standard "log" library
// that adds some basic support for levels.  The current log output
// level can be set by calling SetLevel()
//

import (
	"fmt"
	"log"
	"strings"
)

const (
	ERROR = 1
	WARN  = 2
	INFO  = 3
	DEBUG = 4
	TRACE = 5
)

var levelPrefix = map[int]string{
	ERROR: "ERROR - ",
	WARN:  "WARN - ",
	INFO:  "INFO - ",
	DEBUG: "DEBUG - ",
	TRACE: "TRACE - ",
}

var levels = map[string]int{
	"error": ERROR,
	"warn":  WARN,
	"info":  INFO,
	"debug": DEBUG,
	"trace": TRACE,
}

// logLevel contains the current log level.  The default is INFO.
var logLevel = INFO

// SetLevel sets the current log level.  levelName should be one of the following strings:
//
//	"ERROR"
//	"WARN"
//	"INFO"
//	"DEBUG"
//	"TRACE"
func SetLevel(levelName string) {
	logLevel = levels[strings.ToLower(levelName)]
	if logLevel == 0 {
		logLevel = INFO
	}
}

// GetLevel returns the current log level
func GetLevel() int {
	return logLevel
}

// Log sends output to the standard logger.  Arguments are handled in the manner of fmt.Println
func Log(level int, format string, v ...interface{}) {
	var entry string

	if logLevel < level {
		return
	}

	prefix := levelPrefix[level]
	if v == nil {
		entry = prefix + format
	} else {
		entry = prefix + fmt.Sprintf(format, v...)
	}
	log.Println(entry)
}

// Trace sends trace-level output to the standard logger
func Trace(format string, v ...interface{}) {
	Log(TRACE, format, v...)
}

// Debug sends debug-level output to the standard logger
func Debug(format string, v ...interface{}) {
	Log(DEBUG, format, v...)
}

// Info sends info-level output to the standard logger
func Info(format string, v ...interface{}) {
	Log(INFO, format, v...)
}

// Warn sends warning-level output to the standard logger
func Warn(format string, v ...interface{}) {
	Log(WARN, format, v...)
}

// Error sends error-level output to the standard logger
func Error(format string, v ...interface{}) {
	Log(ERROR, format, v...)
}
