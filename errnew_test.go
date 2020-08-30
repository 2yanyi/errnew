package errnew

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func Test(T *testing.T) {

	go demo()

	_ = http.ListenAndServe(":2000", nil)
}

func demo() {
	defer func() {
		if e := recover(); e != nil {
			PanicTrace(e)
		}
	}()

	demo_test()
}

func demo_test() {
	if _, e := ioutil.ReadFile("test.txt"); e != nil {
		e = Join(e, "ioutil.ReadFile()")
		panic(e)
	}
}
