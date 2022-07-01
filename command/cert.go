package command

import (
	"fmt"
	"kulana/filter"
	"kulana/hostinfo"
	"kulana/output"
	"kulana/template"
)

type CertCommand struct {
	Hostname string `long:"hostname" alias:"host" description:"The hostname to ping" value-name:"HOSTNAME" required:"true"`
}

var certCommand CertCommand

var certCommandDescription = Description{
	CommandName:      "cert",
	ShortDescription: "Resolves SSL certificate information for a hostname",
	LongDescription:  "Resolves SSL certificate information for a hostname",
	Command:          &certCommand,
}

func (c *CertCommand) Execute(args []string) error {
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
		Port:                  false,
		Content:               false,
		ForeignID:             false,
		CertificateValid:      true,
		CertificateValidUntil: true,
		CertificateIssuer:     true,
	}

	of := output.Output{}
	isValid, validUntil, issuer := hostinfo.CheckCertificate(c.Hostname)
	valid := 0
	if isValid {
		valid = 1
	}
	o := output.Output{Hostname: c.Hostname, CertificateValid: valid, CertificateValidUntil: validUntil.Format("2006-01-02 15:04:05"), CertificateIssuer: issuer}
	of = o.Filter(f)
	formatted := template.Render(defaultOptions.Format, of, defaultOptions.NoColor)
	fmt.Println(formatted)
	return nil
}
