package _app

import (
	"fmt"
	"kulana/_hostinfo"
	"kulana/_misc"
)

func runMX(application Application) {
	_, f := _hostinfo.FetchAsOutput(application.Host, application.Filter)
	fmt.Print(template.Render(application.OutputFormat, f) + _misc.GetNLChar())
}
