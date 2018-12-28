package config

var (
	dateFormat = "%Y%m%d%H.%i"
	kindList   map[string]string
)

func init() {
	kindList = map[string]string{
		"info":     "access.log",
		"debug":    "debug.log",
		"trace":    "trace.log",
		"warn":     "warn.log",
		"error":    "error.log",
		"critical": "critical.log",
	}
}

func InitConfigString(data map[string]string) string {

	str := `<?xml version="1.0" encoding="utf-8" ?>
			<logs>`
	for k, v := range kindList {
		tmp := buildXmlConfig(k, v, data["path"], data["size"])
		str = str + `
			  ` + tmp
	}

	str = str + `
		  </logs>`

	return str
}

func buildXmlConfig(key, value, path, size string) string {
	str := "<" + key + `>
				<rotate filename="` + dateFormat + `.` + value + `" dir="` + path + `" size="` + size + `" />
		    </` + key + ">"

	return str
}

func xmlbak(data map[string]string) string {
	xml := `
		<?xml version="1.0" encoding="utf-8" ?>
		<logs>
			<info >
				<rotate filename="` + dateFormat + `.access.log" dir="` + data["path"] + `" size="` + data["size"] + `" />
			</info>

			<debug>
				<rotate  filename="%Y%m%d.%H.%i.debug.log" dir="` + data["path"] + `" size="` + data["size"] + `" />
			</debug>

			<trace>
				<rotate  filename="%Y%m%d.%i.trace.log" dir="` + data["path"] + `" size="` + data["size"] + `" />
			</trace>

			<warn>
				<buffer size="5">
					<rotate  filename="%Y%m%d.%i.warn.log" dir="` + data["path"] + `" size="` + data["size"] + `" />
				</buffer>
			</warn>

			<error>
				<rotate filename="%Y%m%d.%H.%i.error.log" dir="` + data["path"] + `" size="` + data["size"] + `" />
				<console output="stderr" foreground="red" background="blue" />
			</error>

			<critical>
				<rotate filename="%Y%m%d.%i.critical.log" dir="` + data["path"] + `" size="` + data["size"] + `" />
				<console output="stderr" foreground="red" background="blue" />
			</critical>
		</logs>
		`
	return xml
}
