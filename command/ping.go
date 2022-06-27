package command

import (
	"fmt"
	"kulana/filter"
	"kulana/output"
	"kulana/ping"
	"kulana/template"
	"time"
)

type PingCommand struct {
	Hostname     string `long:"hostname" alias:"host" description:"The hostname to ping" value-name:"HOSTNAME" required:"true"`
	Port         int    `short:"p" long:"port" description:"The port to ping" value-name:"PORT" required:"true" default:"0"`
	Timeout      int    `short:"t" long:"timeout" description:"The timeout in seconds" value-name:"TIMEOUT" required:"true" default:"30"`
	Ports        []int  `long:"ports" description:"The ports to ping" value-name:"PORTS"`
	DefaultPorts bool   `long:"default-ports" description:"Use the default ports"`
}

var pingCommand PingCommand

var pingCommandDescription = CommandDescription{
	CommandName:      "ping",
	ShortDescription: "Pings a given host",
	LongDescription:  "The ping command pings a given host",
	Command:          &pingCommand,
}

var defaultPorts = []int{
	21, 22, 80, 143, 443,
}

func (c *PingCommand) Execute(args []string) error {
	SetFormat()
	f := filter.Filter{
		Url:                   false,
		Status:                false,
		Time:                  true,
		Destination:           false,
		ContentLength:         false,
		IpAddress:             true,
		MXRecords:             false,
		ICMPCode:              false,
		PingSuccessful:        true,
		Hostname:              true,
		Port:                  true,
		Content:               false,
		ForeignID:             false,
		CertificateValid:      false,
		CertificateValidUntil: false,
		CertificateIssuer:     false,
	}

	of := output.Output{}

	if c.DefaultPorts {
		c.Ports = defaultPorts
	}

	if c.Port != 0 {
		c.Ports = append(c.Ports, c.Port)
	}

	if len(c.Ports) == 0 {
		for {
			if c.Port == 0 {
				_, of = ping.ICMPAsOutput(c.Hostname, c.Timeout, f)
			} else {
				_, of = ping.PortAsOutput(c.Hostname, c.Port, ping.ProtocolTCP, c.Timeout, f)
			}

			formatted := template.Render(defaultOptions.Format, of, defaultOptions.NoColor)
			fmt.Println(formatted)

			if defaultOptions.Loop {
				time.Sleep(time.Duration(defaultOptions.Delay) * time.Millisecond)
			} else {
				break
			}
		}
	} else {
		for _, port := range c.Ports {
			_, of := ping.PortAsOutput(c.Hostname, port, ping.ProtocolTCP, c.Timeout, f)
			formatted := template.Render(defaultOptions.Format, of, defaultOptions.NoColor)
			fmt.Println(formatted)
		}
	}
	return nil
}
