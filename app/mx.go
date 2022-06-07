package app

import (
	"fmt"
	"kulana/hostinfo"
	"kulana/misc"
	"kulana/template"
)

func runMX(application Application) {
	_, f := hostinfo.FetchAsOutput(application.Host, application.Filter)
	fmt.Print(template.Render(application.OutputFormat, f) + misc.GetNLChar())
}
