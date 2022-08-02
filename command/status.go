package command

import (
	"fmt"
	"kulana/fetcher"
	"kulana/filter"
	"kulana/template"
	"time"
)

type StatusCommand struct {
	Url      []string `short:"u" long:"url" description:"The URL to get te status from" value-name:"URL" required:"true"`
	CheckSSL bool     `short:"s" long:"check-ssl" description:"Check the SSL certificate"`
}

var statusCommand StatusCommand

var statusCommandDescription = Description{
	CommandName:      "status",
	ShortDescription: "Fetches the status of the given URL",
	LongDescription:  "Fetches the status of the given URL",
	Command:          &statusCommand,
}

var statusFilter = filter.Filter{
	Url:            true,
	Status:         true,
	Time:           true,
	Destination:    true,
	ContentLength:  true,
	IpAddress:      true,
	MXRecords:      false,
	ICMPCode:       false,
	PingSuccessful: false,
	Hostname:       false,
	Port:           false,
	Content:        false,
	ForeignID:      true,
	Certificate: filter.CertificateFilter{
		Valid:      true,
		ValidUntil: true,
		Issuer:     true,
	},
}

func (c *StatusCommand) Execute(args []string) error {
	SetFormat()
	for {
		for _, url := range c.Url {
			out := fetcher.FetchHTTPHost(url, defaultOptions.ForeignId, c.CheckSSL)
			out = out.Filter(statusFilter)
			formatted := template.Render(defaultOptions.Format, out, defaultOptions.NoColor)
			fmt.Println(formatted)
		}

		if !defaultOptions.Loop {
			break
		}

		if defaultOptions.Delay < 1000 {
			defaultOptions.Delay = 1000
		}

		time.Sleep(time.Duration(defaultOptions.Delay) * time.Millisecond)
	}
	return nil
}
