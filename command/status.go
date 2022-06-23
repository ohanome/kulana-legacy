package command

import (
	"fmt"
	"kulana/fetcher"
	"kulana/filter"
	"kulana/template"
	"time"
)

type StatusCommand struct {
	Url     string   `short:"u" long:"url" description:"The URL to get te status from" value-name:"URL" required:"true"`
	Include []string `long:"include" description:"Includes specific fields" choice:"url" choice:"status" choice:"time" choice:"destination" choice:"content_length" choice:"ip_address" choice:"mx_records" choice:"icmp_code" choice:"hostname" choice:"port" choice:"content" choice:"foreign_id"`
}

var statusCommand StatusCommand

func (c *StatusCommand) Execute(args []string) error {
	SetFormat()
	for {
		out := fetcher.FetchHTTPHost(c.Url, defaultOptions.ForeignId)
		f := filter.FromOptions(c.Include)
		// Apply defaults
		f.Url = true
		f.Status = true
		f.Time = true
		out = out.Filter(f)
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
