package errnew

import (
	"io/ioutil"
	"testing"
)

func Test(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			// errnew.Trace
			Trace(e)
		}
	}()

	if _, e := ioutil.ReadFile("test.txt"); e != nil {
		// errnew.Join
		e = Join(e, "ioutil.ReadFile()")
		panic(e)
	}

}
