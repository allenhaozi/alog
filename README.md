alog
======
base on: https://github.com/issue9/logs
```go

config := map[string]string{
	"path"    : "point/to/your/log/path",
	"size"    : "200M",
	"buf_cnt" : "100",
}
//description
//buf_cnt : set type of info log buffer size, default is 10
//the log won't write to file until reach the buffer size or call alog.Flush

alog.InitALog(config)
alog.Error("occur error")
alog.Info("access log")
alog.Info(config)

```

### install

```shell
    go get github.com/allenhaozi/alog
```
