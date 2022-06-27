package command

import (
	"fmt"
	"kulana/filter"
	"kulana/hostinfo"
	"kulana/output"
	"kulana/template"
)

type CNAMECommand struct {
	Hostname string `long:"hostname" alias:"host" description:"The hostname to ping" value-name:"HOSTNAME" required:"true"`
}

var cnameCommand CNAMECommand

var cnameCommandDescription = CommandDescription{
	CommandName:      "cname",
	ShortDescription: "Fetches CNAME records for a hostname",
	LongDescription:  "Fetches CNAME records for a hostname",
	Command:          &cnameCommand,
}

func (c *CNAMECommand) Execute(args []string) error {
	SetFormat()
	f := filter.Filter{
		Url:                   false,
		Status:                false,
		Time:                  false,
		Destination:           false,
		ContentLength:         false,
		IpAddress:             false,
		MXRecords:             false,
		ICMPCode:              false,
		PingSuccessful:        false,
		Hostname:              true,
		CNAME:                 true,
		Port:                  false,
		Content:               false,
		ForeignID:             false,
		CertificateValid:      false,
		CertificateValidUntil: false,
		CertificateIssuer:     false,
	}

	of := output.Output{}
	cname := hostinfo.LookupCNAME(c.Hostname)
	o := output.Output{CNAME: cname, Hostname: c.Hostname}
	of = o.Filter(f)
	formatted := template.Render(defaultOptions.Format, of, defaultOptions.NoColor)
	fmt.Println(formatted)
	return nil
}
