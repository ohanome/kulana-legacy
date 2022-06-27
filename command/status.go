package command

import (
	"fmt"
	"kulana/fetcher"
	"kulana/filter"
	"kulana/template"
	"time"
)

type StatusCommand struct {
	Url      string `short:"u" long:"url" description:"The URL to get te status from" value-name:"URL" required:"true"`
	CheckSSL bool   `short:"s" long:"check-ssl" description:"Check the SSL certificate"`
}

var statusCommand StatusCommand

var statusFilter = filter.Filter{
	Url:                   true,
	Status:                true,
	Time:                  true,
	Destination:           true,
	ContentLength:         true,
	IpAddress:             true,
	MXRecords:             false,
	ICMPCode:              false,
	PingSuccessful:        false,
	Hostname:              false,
	Port:                  false,
	Content:               false,
	ForeignID:             true,
	CertificateValid:      true,
	CertificateValidUntil: true,
	CertificateIssuer:     true,
}

func (c *StatusCommand) Execute(args []string) error {
	SetFormat()
	for {
		out := fetcher.FetchHTTPHost(c.Url, defaultOptions.ForeignId, c.CheckSSL)
		out = out.Filter(statusFilter)
		formatted := template.Render(defaultOptions.Format, out, defaultOptions.NoColor)
		fmt.Println(formatted)

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
