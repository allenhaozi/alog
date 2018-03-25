package config

func InitConfigString() string {
	xml := `
		<?xml version="1.0" encoding="utf-8" ?>
		<logs>
			<info >
				<!-- info内容，先缓存到一定10条，再一次性输出 -->
				<buffer size="1">
					<rotate filename="%Y%m%d.%i.access.log" dir="-" size="20M" />
				</buffer>
			</info>

			<debug>
				<rotate  filename="%Y%m%d.%i.debug.log" dir="-" size="20M" />
			</debug>

			<trace>
				<rotate  filename="%Y%m%d.%i.trace.log" dir="-" size="20M" />
			</trace>

			<warn>
				<buffer size="5">
					<rotate  filename="%Y%m%d.%i.warn.log" dir="-" size="20M" />
				</buffer>
			</warn>

			<error>
				<console output="stderr" foreground="red" background="blue" />
				<rotate filename="%Y%m%d.%i.error.log" dir="-" size="20M" />
			</error>

			<critical>
				<console output="stderr" foreground="red" background="blue" />
				<rotate filename="%Y%m%d.%i.critical.log" dir="-" size="20M" />
			</critical>
		</logs>
`
	return xml
}
