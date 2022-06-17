package _app

import (
	"fmt"
	"kulana/_filter"
	"kulana/_misc"
	"kulana/_ping"
	"time"
)

func processPingArgs(application Application) Application {
	return application
}

func runPing(application Application) {
	for {
		var f _filter.Output
		if application.Port < 1 {
			_, f = _ping.ICMPAsOutput(application.Host, application.Timeout, application.Filter)
		} else {
			_, f = _ping.PortAsOutput(application.Host, application.Port, application.Protocol, application.Timeout, application.Filter)
		}
		fmt.Print(template.Render(application.OutputFormat, f) + _misc.GetNLChar())

		if !application.RunInLoop {
			return
		}
		time.Sleep(time.Duration(application.Delay))
	}
}
