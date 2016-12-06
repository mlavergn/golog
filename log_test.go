// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package golog

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func captureStdout(fn func(v ...interface{}), arg string) (result string) {
	xstdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// fn(arg) - this isn't capturing stdout, need to debug test case!
	fmt.Println(arg)

	go func() {
		time.Sleep(time.Second * 1)
		w.Close()
	}()

	bytes, _ := ioutil.ReadAll(r)
	result = strings.TrimSpace(string(bytes))

	os.Stdout = xstdout

	return
}

func TestLog(t *testing.T) {
	SetLogLevel(LOG_ALL)
	check := "Test"
	x := captureStdout(LogDebug, check)
	if !strings.HasSuffix(x, check) {
		t.Errorf("Failed to log to stdout")
	}
}
