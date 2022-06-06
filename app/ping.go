package app

import (
	"fmt"
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
		_, f := ping.PortAsOutput(application.Host, application.Port, application.Protocol, application.Timeout, application.Filter)
		fmt.Print(template.Render(application.OutputFormat, f) + misc.GetNLChar())

		if !application.RunInLoop {
			return
		}
		time.Sleep(time.Duration(application.Delay))
	}
}
