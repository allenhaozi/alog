package config

func InitConfigString(data map[string]string) string {
	xml := `
		<?xml version="1.0" encoding="utf-8" ?>
		<logs>
			<info >
				<!-- info内容，先缓存到一定10条，再一次性输出 -->
				<buffer size="1">
					<rotate filename="%Y%m%d.%i.access.log" dir="` + data["path"] + `" size="20M" />
				</buffer>
			</info>

			<debug>
				<rotate  filename="%Y%m%d.%i.debug.log" dir="` + data["path"] + `" size="20M" />
			</debug>

			<trace>
				<rotate  filename="%Y%m%d.%i.trace.log" dir="` + data["path"] + `" size="20M" />
			</trace>

			<warn>
				<buffer size="5">
					<rotate  filename="%Y%m%d.%i.warn.log" dir="` + data["path"] + `" size="20M" />
				</buffer>
			</warn>

			<error>
				<rotate filename="%Y%m%d.%i.error.log" dir="` + data["path"] + `" size="20M" />
				<console output="stderr" foreground="red" background="blue" />
			</error>

			<critical>
				<rotate filename="%Y%m%d.%i.critical.log" dir="` + data["path"] + `" size="20M" />
				<console output="stderr" foreground="red" background="blue" />
			</critical>
		</logs>
`
	return xml
}
