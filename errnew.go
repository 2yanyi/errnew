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

// 对error追加描述
func Join(err error, errMsg string) error {
	return errors.New(f() + ": " + errMsg + "\n" + err.Error())
}

// 对panic信息截取并记录
func PanicTrace(err interface{}) {
	var now = time.Now().Local().Format(time.RFC3339)
	var res = panicTrace(err, now)
	println(res)
	var f *os.File
	f, err = os.OpenFile("panic-"+now[:10]+".log", 0x2|0x400|0x40, 0600)
	if err == nil {
		_, _ = f.WriteString(res)
		_ = f.Close()
	}
}

func panicTrace(err interface{}, now string) (res string) {
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
	res = fmt.Sprintf("[  ERROR  ] %s\npanic: %v\n%s\n", now, err, msg)
	return
}

func f() string {
	var pc = make([]uintptr, 1)
	runtime.Callers(3, pc)
	return runtime.FuncForPC(pc[0]).Name()
}
