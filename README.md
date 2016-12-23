# Golog
--
### Minimal overhead leveled logging routines for golang.

Introduction
--
Wraps log into a higher level abstraction that handles the following:

* Log levels
* File output
* Syslog output
* Performance timings

The goals of Golog are to get as close as possible to zero overhead, while allowing fine grained control of logging.


Dependencies
--

This package has no external dependencies.


Installation
--
```bash
	go get github/mlavergn/golog
```

Examples
--
```go
	// import ."golog"

	SetLogLevel(LOG_ALL)
	LogDebug("debug")
	LogInfo("info")
	LogWarn("warning")
	LogError("error")

	LogConfigure(LOG_WARN, LOG_STDOUT)
	LogDebug("debug")
	LogInfo("info")
	LogWarn("warning")
	LogError("error")
```
