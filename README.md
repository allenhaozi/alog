alog
======
base on: https://github.com/issue9/logs

```xml
<?xml version="1.0" encoding="utf-8" ?>
    <logs>
        <debug>
            <buffer size="10">
                <rotate dir="/var/log/" size="5M" />
                <stmp username=".." password=".." />
            </buffer>
        </debug>
        <info>
            ....
        </info>
    </logs>
```

```go
logs.InitFromXMLFile("./config.xml")
logs.Debug("debug start...")
logs.Debugf("%v start...", "debug")
logs.DEBUG().Println("debug start...")
```

add new Funciton nitALog, you can config it through map
```go
config := make(map[string]string)
config["path"] = "your/log/path/"
//default size is 20M
config["size"] = "5M"
alog.InitALog(config)
alog.Error("occur error")
alog.Info("access log")
alog.Info(config)
```

### install

```shell
    go get github.com/Allenhaozi/alog
```
