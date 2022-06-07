package app

import (
	"fmt"
	"kulana/filter"
	"kulana/misc"
	"kulana/ping"
	"kulana/template"
	"time"
)

func processPingArgs(application Application) Application {
	return application
}

func runPing(application Application) {
	for {
		var f filter.Output
		if application.Port < 1 {
			_, f = ping.ICMPAsOutput(application.Host, application.Timeout, application.Filter)
		} else {
			_, f = ping.PortAsOutput(application.Host, application.Port, application.Protocol, application.Timeout, application.Filter)
		}
		fmt.Print(template.Render(application.OutputFormat, f) + misc.GetNLChar())

		if !application.RunInLoop {
			return
		}
		time.Sleep(time.Duration(application.Delay))
	}
}
