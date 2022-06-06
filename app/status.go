package app

import (
	"fmt"
	"kulana/email"
	"kulana/fetcher"
	"kulana/misc"
	"kulana/template"
	"time"
)

func runStatus(application Application) {
	for {
		response, filteredResponse := fetcher.FetchAndFilter(application.Url, application.Filter)
		fmt.Print(template.Render(application.OutputFormat, filteredResponse) + misc.GetNLChar())
		if application.NotifyViaMail {
			email.SendNotificationMail(application.NotifyMailTo, application.Url, response.Status)
		}
		if application.FollowRedirect && !application.RunInLoop && response.Status < 400 && response.Status >= 300 {
			application.Url = response.Destination
			response, filteredResponse = fetcher.FetchAndFilter(application.Url, application.Filter)
			fmt.Print(template.Render(application.OutputFormat, filteredResponse) + misc.GetNLChar())
			if application.NotifyViaMail {
				email.SendNotificationMail(application.NotifyMailTo, application.Url, response.Status)
			}
		}

		if !application.RunInLoop {
			return
		}
		time.Sleep(time.Duration(application.Delay))
	}
}
