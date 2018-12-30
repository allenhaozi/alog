package config

const (
	info = "info"
)

var (
	dateFormat = "%Y%m%d%H.%i"
	kindList   = map[string]string{
		info:       "access.log",
		"debug":    "debug.log",
		"trace":    "trace.log",
		"warn":     "warn.log",
		"error":    "error.log",
		"critical": "critical.log",
	}
	defBufCnt = "10"
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
		if len(cnt) <= 0 {
			cnt = defBufCnt
		}
		str = "<" + key + `>
				<buffer size="cnt">
				<rotate filename="` + dateFormat + `.` + value + `" dir="` + path + `" size="` + size + `" />
				</buffer>
		    </` + key + ">"
	} else {
		str = "<" + key + `>
				<rotate filename="` + dateFormat + `.` + value + `" dir="` + path + `" size="` + size + `" />
		    </` + key + ">"
	}
	return str
}
