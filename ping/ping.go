package _ping

import (
	"fmt"
	"kulana/filter"
	"kulana/hostinfo"
	"kulana/l"
	"kulana/output"
	"net"
	"os"
	"regexp"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

const Timeout = 30
const ProtocolTCP = "tcp"
const ProtocolUDP = "udp"
const DefaultProtocol = ProtocolUDP
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
		info := hostinfo.Fetch(host)
		ipaddress = info.IPAddress
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

func PortAsOutput(host string, port int, protocol string, timeout int, f filter.Filter) (output.Output, output.Output) {
	t, ip, err := Port(host, port, protocol, timeout)
	if err != nil {
		l.Emergency(err.Error())
	}

	o := output.Output{
		Hostname:      host,
		Port:          port,
		Status:        0,
		Time:          t,
		Destination:   "",
		ContentLength: 0,
		IpAddress:     ip,
		ICMPCode:      -1,
	}

	filtered := o.Filter(f)

	return o, filtered
}

func ICMP(target string, timeout int) (string, int, bool, float64, error) {
	ip, err := net.ResolveIPAddr("ip4", target)
	if err != nil {
		return "", 0, false, 0, err
	}
	conn, err := icmp.ListenPacket("udp4", "0.0.0.0")
	if err != nil {
		return "", 0, false, 0, err
	}
	defer func(conn *icmp.PacketConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte(""),
		},
	}
	msgBytes, err := msg.Marshal(nil)
	if err != nil {
		return "", 0, false, 0, err
	}

	// Write the message to the listening connection
	s := time.Now()
	if _, err := conn.WriteTo(msgBytes, &net.UDPAddr{IP: net.ParseIP(ip.String())}); err != nil {
		return "", 0, false, 0, err
	}

	err = conn.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
	if err != nil {
		return "", 0, false, 0, err
	}
	reply := make([]byte, 1500)
	n, _, err := conn.ReadFrom(reply)
	e := time.Now()
	elapsed := float64(e.Sub(s)) / float64(time.Millisecond)

	if err != nil {
		return "", 0, false, 0, err
	}
	parsedReply, err := icmp.ParseMessage(1, reply[:n])

	if err != nil {
		return "", 0, false, 0, err
	}

	ok := false
	if parsedReply.Code == 0 {
		ok = true
	}

	return ip.String(), parsedReply.Code, ok, elapsed, nil
}

func ICMPAsOutput(target string, timeout int, f filter.Filter) (output.Output, output.Output) {
	ip, icmpCode, _, elapsed, err := ICMP(target, timeout)
	if err != nil {
		l.Emergency(err.Error())
	}

	o := output.Output{
		Hostname:      target,
		Status:        0,
		Time:          elapsed,
		Destination:   "",
		ContentLength: 0,
		IpAddress:     ip,
		ICMPCode:      icmpCode,
	}

	of := o.Filter(f)
	return o, of
}
