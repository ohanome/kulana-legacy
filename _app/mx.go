package _app

import (
	"fmt"
	"kulana/_hostinfo"
	"kulana/_misc"
	"kulana/_template"
)

func runMX(application Application) {
	_, f := _hostinfo.FetchAsOutput(application.Host, application.Filter)
	fmt.Print(_template.Render(application.OutputFormat, f) + _misc.GetNLChar())
}
