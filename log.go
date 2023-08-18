package golog

import (
	"io"
	"io/ioutil"
	"log"
	"log/syslog"
	"os"
	"os/user"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Log levels
const (
	LOG_ALL = iota
	LOG_DEBUG
	LOG_INFO
	LOG_WARN
	LOG_ERROR
	LOG_FATAL
	LOG_OFF
)

// Log destinations
const (
	LOG_STDOUT = iota
	LOG_STDERR
	LOG_FILE
	LOG_SYSTEM
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

// Modify the log level.
// The default log level is LOG_WARN.
func SetLogLevel(level int) {
	LogConfigure(level, LOG_STDOUT)
}

//
// singleton
//

var (
	once sync.Once
)

// logConfigOutput configures the logger.
func LogConfigure(level int, dest int) {
	once.Do(func() {
		log.SetFlags(log.Lshortfile | log.LstdFlags | log.LUTC | log.Lmicroseconds)
	})

	// if level == LOG_OFF, shutdown all logging
	if level == LOG_OFF {
		log.SetOutput(ioutil.Discard)
		return
	}

	switch dest {
	case LOG_STDOUT:
		log.SetOutput(os.Stdout)
	case LOG_STDERR:
		log.SetOutput(os.Stderr)
	case LOG_FILE:
		logPath := obtainLogDirectory(obtainProcessName())
		logfile, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE, 0640)
		if err == nil {
			gologger := log.New(io.MultiWriter(logfile, os.Stdout), log.Prefix(), log.Flags())
			_Output = gologger.Println
			_Outputf = gologger.Printf
			_OutputExit = gologger.Fatalln
			_OutputExitf = gologger.Fatalf
		}
	case LOG_SYSTEM:
		gologger, err := syslog.NewLogger(syslog.LOG_NOTICE, log.Flags())
		gologgerErr, errErr := syslog.NewLogger(syslog.LOG_ERR, log.Flags())
		if err == nil && errErr == nil {
			_Output = gologger.Println
			_Outputf = gologger.Printf
			_OutputExit = gologgerErr.Fatalln
			_OutputExitf = gologgerErr.Fatalf
		}
	}

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

// obtainProcessName derives the process name from argv
func obtainProcessName() (result string) {
	result = os.Args[0]
	last := strings.LastIndex(result, "/")
	if last > -1 {
		result = result[last+1:]
		result = strings.Replace(result, " ", "_", -1)
	}

	return
}

// obtainLogDirectory derives and creates the log director path for a given log name
func obtainLogDirectory(logName string) (result string) {
	usr, err := user.Current()
	if err == nil {
		logPath := path.Join(usr.HomeDir, "log", logName)

		_, err = os.Stat(logPath)
		if os.IsNotExist(err) {
			os.MkdirAll(logPath, os.ModePerm)
		}

		epoch := strconv.Itoa(int(time.Now().Unix()))
		logFile := epoch + ".log"

		result = path.Join(logPath, logFile)
	}

	return
}

// Dump to file
func LogDumpFile(modulename string, output string) {
	logPath := obtainLogDirectory(modulename)

	if len(logPath) > 0 {
		err := ioutil.WriteFile(logPath, []byte(output), 0644)

		if err != nil {
			LogError(err)
		}
	}
}

//
// Performance timing routines
//

var timerShared time.Time

// TimerMark marks the beginning of a timing period
func TimerMark() {
	timerShared = time.Now()
}

// TimerMeasure outputs at level LOG_DEBUG the elapsed time since the last call to TimerMark
func TimerMeasure() {
	LogDebugf("ELAPSED [%s]ns", time.Since(timerShared))
}
