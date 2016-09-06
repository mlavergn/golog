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

// Debug accepts a value, outputting when LOG_LEVEL <= LOG_DEBUG
var Debug _output = _DevNull

// Debugf accepts a format mask and a value, outputting when LOG_LEVEL <= LOG_DEBUG
var Debugf _outputf = _DevNullf

// Info accepts a value, outputting when LOG_LEVEL <= LOG_INFO
var Info _output = _Output

// Infof accepts a format mask and a value, outputting when LOG_LEVEL <= LOG_INFO
var Infof _outputf = _Outputf

// Warn accepts a value, outputting when LOG_LEVEL <= LOG_WARN
var Warn _output = _Output

// Warnf accepts a format mask and a value, outputting when LOG_LEVEL <= LOG_WARN
var Warnf _outputf = _Outputf

// Error accepts a value, outputting when LOG_LEVEL <= LOG_ERROR
var Error _output = _Output

// Errorf accepts a format mask and a value, outputting when LOG_LEVEL <= LOG_ERROR
var Errorf _outputf = _Outputf

// Fatal accepts a value, outputting when LOG_LEVEL <= LOG_FATAL, then exitting
var Fatal _output = _OutputExit

// Fatalf accepts a format mask and a value, outputting when LOG_LEVEL <= LOG_FATAL, then exitting
var Fatalf _outputf = _OutputExitf

//
// Modify the log level.
// The default log level is LOG_WARN.
//
func SetLogLevel(level int) {
	if level <= LOG_DEBUG {
		Debug = _Output
		Debugf = _Outputf
	} else {
		Debug = _DevNull
		Debugf = _DevNullf
	}

	if level <= LOG_INFO {
		Info = _Output
		Infof = _Outputf
	} else {
		Info = _DevNull
		Infof = _DevNullf
	}

	if level <= LOG_WARN {
		Warn = _Output
		Warnf = _Outputf
	} else {
		Warn = _DevNull
		Warnf = _DevNullf
	}

	if level <= LOG_ERROR {
		Error = _Output
		Errorf = _Outputf
	} else {
		Error = _DevNull
		Errorf = _DevNullf
	}

	if level <= LOG_FATAL {
		Fatal = log.Fatalln
		Fatalf = log.Fatalf
	} else {
		Fatal = _DevNull
		Fatalf = _DevNullf
	}
}

//
// Dump to file
//
func DumpFile(modulename string, output string) {
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
	Debugf("ELAPSED [%s]", time.Since(_timerShared))
}
