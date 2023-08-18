package main

import (
	. "golog"
)

func main() {
	SetLogLevel(LOG_ALL)
	LogDebug("debug")
}
