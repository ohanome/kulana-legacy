package _app

import (
	"fmt"
	"kulana/_email"
	"kulana/_fetcher"
	"kulana/_misc"
	"kulana/_template"
	"time"
)

func runStatus(application Application) {
	for {
		response, filteredResponse := _fetcher.FetchAndFilter(application.Url, application.Filter, application.ForeignID)
		fmt.Print(_template.Render(application.OutputFormat, filteredResponse) + _misc.GetNLChar())
		if application.NotifyViaMail {
			_email.SendNotificationMail(application.NotifyMailTo, application.Url, response.Status)
		}
		if application.FollowRedirect && !application.RunInLoop && response.Status < 400 && response.Status >= 300 {
			application.Url = response.Destination
			response, filteredResponse = _fetcher.FetchAndFilter(application.Url, application.Filter, application.ForeignID)
			fmt.Print(_template.Render(application.OutputFormat, filteredResponse) + _misc.GetNLChar())
			if application.NotifyViaMail {
				_email.SendNotificationMail(application.NotifyMailTo, application.Url, response.Status)
			}
		}

		if !application.RunInLoop {
			return
		}
		time.Sleep(time.Duration(application.Delay))
	}
}
