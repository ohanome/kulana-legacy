package command

import (
	"fmt"
	"kulana/fetcher"
	"kulana/hostinfo"
	"kulana/output"
	"kulana/ping"
	"kulana/template"
	"strings"
)

type InfoCommand struct {
	Url string `short:"u" long:"url" description:"The URL to get te status from" value-name:"URL" required:"true"`
}

var infoCommand InfoCommand

var infoCommandDescription = Description{
	CommandName:      "info",
	ShortDescription: "Fetches all available information about the given URL",
	LongDescription:  "Fetches all available information about the given URL, including port pings, MX records, IP address, CNAME, certificate information, etc.",
	Command:          &infoCommand,
}

func (c *InfoCommand) Execute(args []string) error {
	SetFormat()
	url := c.Url
	hostname := strings.Split(url, "//")[1]

	cname := hostinfo.LookupCNAME(hostname)
	mxRecords := hostinfo.LookupMX(hostname)
	ipAddress := hostinfo.ResolveIPAddress(hostname)

	sOut := fetcher.FetchHTTPHost(c.Url, defaultOptions.ForeignId, true)
	status := sOut.Status
	time := sOut.Time
	destination := sOut.Destination
	contentLength := sOut.ContentLength
	isValid := sOut.Certificate.Valid
	validUntil := sOut.Certificate.ValidUntil
	issuer := sOut.Certificate.Issuer

	var pings []output.PingOutput
	for _, port := range manyPorts {
		elapsed, _, err := ping.Port(hostname, port, ping.ProtocolTCP, 1)
		pingSuccessful := true
		e := ""
		if err != nil {
			pingSuccessful = false
			e = err.Error()
		}
		pings = append(pings, output.PingOutput{
			Successful: pingSuccessful,
			Error:      e,
			Time:       elapsed,
			Port:       port,
		})
	}

	o := output.Output{
		Url:           url,
		Status:        status,
		Time:          time,
		Destination:   destination,
		ContentLength: contentLength,
		IpAddress:     ipAddress,
		MXRecords:     mxRecords,
		Hostname:      hostname,
		CNAME:         cname,
		ForeignID:     defaultOptions.ForeignId,
		Certificate: output.CertificateOutput{
			Valid:      isValid,
			ValidUntil: validUntil,
			Issuer:     issuer,
		},
		Pings: pings,
	}
	formatted := template.Render(defaultOptions.Format, o, defaultOptions.NoColor)
	fmt.Println(formatted)
	return nil
}
