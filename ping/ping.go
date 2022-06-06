package ping

import (
	"fmt"
	"kulana/filter"
	"kulana/misc"
	"net"
	"regexp"
	"time"
)

const Timeout = 30
const ProtocolTCP = "tcp"
const ProtocolUDP = "udp"
const DefaultProtocol = ProtocolTCP
const DefaultPort = 80

func Port(host string, port int, protocol string, timeout int) (float64, string, error) {
	if protocol != ProtocolTCP && protocol != ProtocolUDP {
		protocol = DefaultProtocol
	}

	if timeout < 0 {
		timeout = Timeout
	}

	if port < 1 || port > 65000 {
		port = DefaultPort
	}

	address := fmt.Sprintf("%s:%d", host, port)

	duration := time.Duration(timeout) * time.Second

	start := time.Now()
	conn, err := net.DialTimeout(protocol, address, duration)
	end := time.Now()

	ipaddress := host
	hostIPMatch, _ := regexp.Match(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`, []byte(host))
	if !hostIPMatch {
		ipaddress = conn.RemoteAddr().String()
	}

	if err != nil {
		return 0.0, "", err
	}

	err = conn.Close()
	if err != nil {
		return 0.0, "", err
	}

	elapsed := float64(end.Sub(start)) / float64(time.Millisecond)
	return elapsed, ipaddress, nil
}

func Host(host string, protocol string, timeout int) (float64, string, error) {
	return Port(host, DefaultPort, protocol, timeout)
}

func PortAsOutput(host string, port int, protocol string, timeout int, f filter.OutputFilter) (filter.Output, filter.Output) {
	t, ip, err := Port(host, port, protocol, timeout)
	misc.Check(err)

	o := filter.Output{
		Url:           host,
		Status:        0,
		Time:          t,
		Destination:   "",
		ContentLength: 0,
		IpAddress:     ip,
	}

	filtered := filter.FilterOutput(o, f)

	return o, filtered
}
