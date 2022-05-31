package app

import (
	"fmt"
	"kulana/email"
	"kulana/fetcher"
	"kulana/misc"
	"kulana/template"
	"net/mail"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func processStatusArgs(application Application) Application {
	if len(os.Args) == 2 {
		misc.Usage(CommandStatus)
		os.Exit(1)
	}

	if len(os.Args) >= 2 {
		for _, arg := range os.Args[2:] {
			switch arg {
			case "--help":
			case "-h":
				misc.Usage(CommandStatus)
				os.Exit(0)

			case "--json":
				application.OutputFormat = template.FormatJSON
				break

			case "--csv":
				application.OutputFormat = template.FormatCSV
				break

			case "--loop":
				application.RunInLoop = true
				if application.FollowRedirect {
					misc.Die("Cannot follow redirects in a loop")
				}
				break

			case "--include-length":
			case "-l":
				application.Filter.ContentLength = true
				break

			case "--follow-redirect":
			case "-f":
				application.FollowRedirect = true
				if application.RunInLoop {
					misc.Die("Cannot follow redirects in a loop")
				}
				break

			case "--status-only":
				application.Filter = fetcher.BuildFilterFromNumeric(fetcher.FilterStatus)
				break

			case "--url-only":
				application.Filter = fetcher.BuildFilterFromNumeric(fetcher.FilterUrl)
				break

			case "--time-only":
				application.Filter = fetcher.BuildFilterFromNumeric(fetcher.FilterTime)
				break

			case "--check-env":
				emailSetup := email.CheckMailEnvironment(false)

				if emailSetup {
					fmt.Println("[OK] Email setup")
				}

				os.Exit(0)

			case "--notify":
			case "-n":
				application.NotifyViaMail = true
				break

			default:
				delayMatch, _ := regexp.Match(`^--delay=\d.+$`, []byte(arg))
				if delayMatch {
					d := strings.ReplaceAll(arg, "--delay=", "")
					dInt, _ := strconv.ParseInt(d, 10, 64)
					application.Delay = dInt * 1000000
				}

				urlMatch, _ := regexp.Match(`^http(s|)://\w.+\.\w{2,3}$`, []byte(arg))
				if urlMatch {
					application.Url = arg
				}

				mailMatch, _ := regexp.Match(`^--notify-mail=.+$`, []byte(arg))
				if mailMatch {
					m := strings.ReplaceAll(arg, "--notify-mail=", "")
					_, mErr := mail.ParseAddress(m)
					if mErr == nil {
						application.NotifyMailTo = m
					}
				}
				break
			}
		}

		if application.Url == "" {
			misc.Die("No URL given.")
		}

		if application.NotifyViaMail && application.NotifyMailTo == "" {
			fmt.Println("Email address is missing, no email will be sent.")
			application.NotifyViaMail = false
		}
	}

	// Correct some values if they're wrong.
	if application.Delay < 100 {
		application.Delay = 100
	}

	if application.NotifyViaMail && application.Delay < 30000 {
		application.Delay = 30000
		fmt.Println("When sending notification emails, the minimum delay is 30000 (30 seconds).")
	}

	return application
}

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
			os.Exit(0)
		}
		time.Sleep(time.Duration(application.Delay))
	}
}
