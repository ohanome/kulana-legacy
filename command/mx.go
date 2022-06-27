package command

import (
	"fmt"
	"kulana/filter"
	"kulana/hostinfo"
	"kulana/output"
	"kulana/template"
)

type MXCommand struct {
	Hostname string `long:"hostname" alias:"host" description:"The hostname to ping" value-name:"HOSTNAME" required:"true"`
}

var mxCommand MXCommand

var mxCommandDescription = CommandDescription{
	CommandName:      "mx",
	ShortDescription: "Fetches MX records for a hostname",
	LongDescription:  "Fetches MX records for a hostname",
	Command:          &mxCommand,
}

func (c *MXCommand) Execute(args []string) error {
	SetFormat()
	f := filter.Filter{
		Url:                   false,
		Status:                false,
		Time:                  false,
		Destination:           false,
		ContentLength:         false,
		IpAddress:             false,
		MXRecords:             true,
		ICMPCode:              false,
		PingSuccessful:        false,
		Hostname:              true,
		Port:                  false,
		Content:               false,
		ForeignID:             false,
		CertificateValid:      false,
		CertificateValidUntil: false,
		CertificateIssuer:     false,
	}

	of := output.Output{}
	mx := hostinfo.LookupMX(c.Hostname)
	o := output.Output{MXRecords: mx, Hostname: c.Hostname}
	of = o.Filter(f)
	formatted := template.Render(defaultOptions.Format, of, defaultOptions.NoColor)
	fmt.Println(formatted)
	return nil
}
