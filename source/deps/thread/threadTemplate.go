package thread

const (
	fileName = "thread.go"

	template = `package thread

import (
	"fmt"
	"runtime/debug"
	"strconv"
	"strings"
)

var mainThreadID int64

func SetMainThreadID() (err error) {
	mainThreadID = threadID()
	if mainThreadID == -1 {
		err = fmt.Errorf("main thread ID is unknown")
	}
	return
}

func IsMainThread() (is bool) {
	is = threadID() == mainThreadID
	return
}

func threadID() (id int64) {
	// Set id to invalid number.
	id = -1
	var stack []byte
	if stack = debug.Stack(); len(stack) == 0 {
		return
	}
	var parts []string
	// Format: "goroutine ${N} [running]: ..."
	if parts = strings.Split(string(stack), " "); len(parts) < 2 {
		return
	}
	trimmed := strings.TrimSpace(parts[1])
	var err error
	if id, err = strconv.ParseInt(trimmed, 10, 64); err != nil {
		id = -1
	}
	return
}
`
)
