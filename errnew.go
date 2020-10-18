package errnew

import (
	"errors"
	"fmt"
	"math"
	"os"
	"runtime"
	"strings"
	"time"
)

var __traceContinue_ = []string{"/usr/local/", "/github.com/"}

// Join error description
func Join(e error, errMsg string) error {
	return errors.New(f() + ": " + errMsg + "\n" + e.Error())
}

// Trace and archive to log
func Trace(e interface{}) {
	var now = time.Now().Local().Format(time.RFC3339)
	var res = panicTrace(e, now)
	println(res)
	var f *os.File
	f, e = os.OpenFile("panic-"+now[:10]+".log", 0x2|0x400|0x40, 0600)
	if e == nil {
		_, _ = f.WriteString(res)
		_ = f.Close()
	}
}

func panicTrace(e interface{}, now string) (res string) {
	var msg string
	var trace = make([]byte, 1<<16)
	var stack = strings.Split(string(trace[:int(math.Min(float64(runtime.Stack(trace, true)), 5000))]), "\n")
	for i := range stack {
		if !strings.HasPrefix(stack[i], "	") {
			continue
		}
		has := true
		for j := range __traceContinue_ {
			if strings.Contains(stack[i], __traceContinue_[j]) {
				has = false
				break
			}
		}
		if has {
			msg += stack[i] + "\n"
		}
	}
	res = fmt.Sprintf("[  ERROR  ] %s\npanic: %v\n%s\n", now, e, msg)
	return
}

func f() string {
	var pc = make([]uintptr, 1)
	runtime.Callers(3, pc)
	return runtime.FuncForPC(pc[0]).Name()
}
