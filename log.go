// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

//
// Package golog provides performance sensitive logging routines.
//
package golog

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strconv"
	"time"
)

//
// Log levels
//
const (
	LOG_ALL   = iota
	LOG_DEBUG = iota
	LOG_INFO  = iota
	LOG_WARN  = iota
	LOG_ERROR = iota
	LOG_FATAL = iota
	LOG_OFF   = iota
)

//
// Log destinations
//
const (
	LOG_STDOUT = iota
	LOG_STDERR = iota
	LOG_FILE   = iota
)

func _DevNull(v ...interface{}) {
}

func _DevNullf(f string, v ...interface{}) {
}

type _output func(v ...interface{})
type _outputf func(f string, v ...interface{})

var _Output _output = log.Println
var _Outputf _outputf = log.Printf

var _OutputExit _output = log.Fatalln
var _OutputExitf _outputf = log.Fatalf

// LogDebug accepts a value, outputting when LOG_LEVEL <= LOG_DEBUG
var LogDebug _output = _DevNull

// LogDebugf accepts a format mask and a value, outputting when LOG_LEVEL <= LOG_DEBUG
var LogDebugf _outputf = _DevNullf

// LogInfo accepts a value, outputting when LOG_LEVEL <= LOG_INFO
var LogInfo _output = _Output

// LogInfof accepts a format mask and a value, outputting when LOG_LEVEL <= LOG_INFO
var LogInfof _outputf = _Outputf

// LogWarn accepts a value, outputting when LOG_LEVEL <= LOG_WARN
var LogWarn _output = _Output

// LogWarnf accepts a format mask and a value, outputting when LOG_LEVEL <= LOG_WARN
var LogWarnf _outputf = _Outputf

// LogError accepts a value, outputting when LOG_LEVEL <= LOG_ERROR
var LogError _output = _Output

// LogErrorf accepts a format mask and a value, outputting when LOG_LEVEL <= LOG_ERROR
var LogErrorf _outputf = _Outputf

// LogFatal accepts a value, outputting when LOG_LEVEL <= LOG_FATAL, then exitting
var LogFatal _output = _OutputExit

// LogFatalf accepts a format mask and a value, outputting when LOG_LEVEL <= LOG_FATAL, then exitting
var LogFatalf _outputf = _OutputExitf

//
// Modify the log level.
// The default log level is LOG_WARN.
//
func SetLogLevel(level int) {
	if level <= LOG_DEBUG {
		LogDebug = _Output
		LogDebugf = _Outputf
	} else {
		LogDebug = _DevNull
		LogDebugf = _DevNullf
	}

	if level <= LOG_INFO {
		LogInfo = _Output
		LogInfof = _Outputf
	} else {
		LogInfo = _DevNull
		LogInfof = _DevNullf
	}

	if level <= LOG_WARN {
		LogWarn = _Output
		LogWarnf = _Outputf
	} else {
		LogWarn = _DevNull
		LogWarnf = _DevNullf
	}

	if level <= LOG_ERROR {
		LogError = _Output
		LogErrorf = _Outputf
	} else {
		LogError = _DevNull
		LogErrorf = _DevNullf
	}

	if level <= LOG_FATAL {
		LogFatal = log.Fatalln
		LogFatalf = log.Fatalf
	} else {
		LogFatal = _DevNull
		LogFatalf = _DevNullf
	}
}

//
// Dump to file
//
func LogDumpFile(modulename string, output string) {
	usr, err := user.Current()
	if err == nil {
		path := usr.HomeDir + "/log/" + modulename + "/"

		epoch := strconv.Itoa(int(time.Now().Unix()))
		logpath := path + epoch + ".log"

		_, err = os.Stat(path)
		if os.IsNotExist(err) {
			os.MkdirAll(path, os.ModePerm)
		}
		err = ioutil.WriteFile(logpath, []byte(output), 0644)
	}

	if err != nil {
		log.Println(err)
	}
}

//
// Performance timing routines
//

var _timerShared time.Time

//
// TimerMark marks the beginning of a timing period
//
func TimerMark() {
	_timerShared = time.Now()
}

//
// TimerMeasure outputs at level LOG_DEBUG the elapsed time since the last call to TimerMark
//
func TimerMeasure() {
	LogDebugf("ELAPSED [%s]", time.Since(_timerShared))
}
