package config

import (
	"strconv"
)

const (
	info = "info"
)

var (
	//log filename format
	dateFormatYmdHi  = "%Y%m%d%H.%i"
	dateFormatYmdHmi = "%Y%m%d%H%n.%i"
	dateFormatYmdi   = "%Y%m%d.%i"

	kindList = map[string]string{
		info:       "access.log",
		"debug":    "debug.log",
		"trace":    "trace.log",
		"warn":     "warn.log",
		"error":    "error.log",
		"critical": "critical.log",
	}

	//default buffer size
	defBufCnt = "10"
	//custom log flag
	flagCustom1 = "log.ldate|log.lmicroseconds"
	flagCustom2 = "log.ldate|log.lmicroseconds|log.lshortfile"
	flagCustom3 = "log.ldate|log.lmicroseconds|log.llongfile"
)

func InitConfigString(data map[string]string) string {

	str := `<?xml version="1.0" encoding="utf-8" ?>
			<logs>`
	for k, v := range kindList {
		tmp := buildXmlConfig(k, v, data["path"], data["size"], data["buf_cnt"])
		str = str + `
			  ` + tmp
	}

	str = str + `
		  </logs>`

	return str
}

func buildXmlConfig(key, value, path, size, cnt string) string {
	var str string
	if key == info {
		i, err := strconv.Atoi(cnt)
		if i < 0 || err != nil {
			cnt = defBufCnt
		}
		str = "<" + key + ` prefix="` + key + ` " flag="` + flagCustom1 + `">
				<buffer size="` + cnt + `">
				<rotate filename="` + dateFormatYmdHi + `.` + value + `" dir="` + path + `" size="` + size + `" />
				</buffer>
		    </` + key + ">"
	} else {
		str = "<" + key + ` prefix="` + key + ` " flag="` + flagCustom2 + `">
				<rotate filename="` + dateFormatYmdi + `.` + value + `" dir="` + path + `" size="` + size + `" />
		    </` + key + ">"
	}
	return str
}
