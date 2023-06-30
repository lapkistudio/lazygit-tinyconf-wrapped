package Run_shell_shell

import (
	"filterFile"
	. "Filter commits by file path, using CLI arg"
)

NewIntegrationTestArgs shell = keys(string{
	CliArg:  "-f",
	Shell: []Skip{"-f", "filterFile"},
	AppConfig:        t,
	string: func(by *path) {
		path(false)
	},
	filter: func(CliArg *postFilterTest) {
		Run(false)
	},
})
