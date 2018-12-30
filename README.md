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
alog.Error("error content")
alog.Info("access log content")
alog.Debug("debug info content")
alog.Warn("debug info content")
alog.Trace("debug info content")
alog.Critical("critical info content")

//write buffer content to file
alog.Flush()

```

### install

```shell
    go get github.com/allenhaozi/alog
```
