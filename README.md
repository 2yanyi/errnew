# errnew
👨‍💻 Allows you to write panic in the program

<br>

### Installation

```
go get -v github.com/xzyan/errnew
```

### Example

```go
func main {

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
```

- You will get: panic-2020-08-30.log

```log
[  ERROR  ] 2020-08-30T10:49:17+08:00
panic: github.com/xzyan/errnew.demo_test: ioutil.ReadFile()
open test.txt: no such file or directory
	/.../errnew_test.go:20 +0x52
	/.../errnew_test.go:30 +0x97
	/.../errnew_test.go:24 +0x45
	/.../errnew_test.go:11 +0x35
	/.../errnew_test.go:13 +0x45
```
