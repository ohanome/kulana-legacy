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

var mxCommandDescription = Description{
	CommandName:      "mx",
	ShortDescription: "Fetches MX records for a hostname",
	LongDescription:  "Fetches MX records for a hostname",
	Command:          &mxCommand,
}

var mxFilter = filter.Filter{
	Url:            false,
	Status:         false,
	Time:           false,
	Destination:    false,
	ContentLength:  false,
	IpAddress:      false,
	MXRecords:      true,
	ICMPCode:       false,
	PingSuccessful: false,
	Hostname:       true,
	Port:           false,
	Content:        false,
	ForeignID:      false,
	Certificate: filter.CertificateFilter{
		Valid:      false,
		ValidUntil: false,
		Issuer:     false,
	},
}

func (c *MXCommand) Execute(args []string) error {
	SetFormat()

	of := output.Output{}
	mx := hostinfo.LookupMX(c.Hostname)
	o := output.Output{MXRecords: mx, Hostname: c.Hostname}
	of = o.Filter(mxFilter)
	formatted := template.Render(defaultOptions.Format, of, defaultOptions.NoColor)
	fmt.Println(formatted)
	return nil
}
