# errnew
👨‍💻 Allows you to write panic in the program

<br>

### Installation

```
go get -v github.com/xzyan/errnew
```

### Example

```go
defer func() {
    if e := recover(); e != nil {
        errnew.PanicTrace(e)
    }
}()

if _, e := ioutil.ReadFile("test.txt"); e != nil {
    e = errnew.Join(e, "ioutil.ReadFile()")
    panic(e)
}
```

- You will get: panic-2020-08-30.log

```log
[  ERROR  ] 2020-08-30T10:49:17+08:00
panic: github.com/xzyan/errnew.demo_test: ioutil.ReadFile()
open test.txt: no such file or directory
```
