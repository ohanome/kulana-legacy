package hostinfo

import (
	"crypto/tls"
	"fmt"
	"kulana/filter"
	"kulana/output"
	"net"
	"regexp"
	"strings"
	"time"
)

type HostInfo struct {
	Hostname   string
	IPAddress  string
	Nameserver []string
	MX         []string
	TXT        []string
	CNAME      string
}

func CheckCertificate(url string) (bool, time.Time, string) {
	hostname := URLToHostname(url)
	conn, err := tls.Dial("tcp", hostname+":443", nil)
	if err != nil {
		// Server doesn't support TLS
		return false, time.Time{}, ""
	}

	err = conn.VerifyHostname(hostname)
	if err != nil {
		return false, time.Time{}, ""
	}

	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	issuer := conn.ConnectionState().PeerCertificates[0].Issuer
	issuerString := fmt.Sprintf("%v", issuer)
	return true, expiry, issuerString
}

func URLToHostname(url string) string {
	urlMatch, err := regexp.Match(`://`, []byte(url))
	if err != nil {
		panic(err)
	}
	if urlMatch {
		url = strings.Split(url, "//")[1]
		url = strings.Split(url, "/")[0]
	}

	return url
}

func ResolveIPAddress(hostname string) string {
	hostname = URLToHostname(hostname)
	ip, err := net.ResolveIPAddr("ip", hostname)
	if err != nil {
		panic(err)
	}
	return ip.String()
}

func LookupMX(hostname string) []string {
	hostname = URLToHostname(hostname)
	mx, mxErr := net.LookupMX(hostname)
	if mxErr != nil {
		panic(mxErr)
	}
	var mxEntries []string
	for _, m := range mx {
		mxEntries = append(mxEntries, strings.TrimSuffix(m.Host, "."))
	}
	return mxEntries
}

func LookupCNAME(hostname string) string {
	hostname = URLToHostname(hostname)
	cname, cnameErr := net.LookupCNAME(hostname)
	if cnameErr != nil {
		panic(cnameErr)
	}
	return strings.TrimSuffix(cname, ".")
}

func LookupNameserver(hostname string) []string {
	hostname = URLToHostname(hostname)
	nameservers, nsErr := net.LookupNS(hostname)
	if nsErr != nil {
		panic(nsErr)
	}
	var nsEntries []string
	for _, ns := range nameservers {
		nsEntries = append(nsEntries, strings.TrimSuffix(ns.Host, "."))
	}
	return nsEntries
}

func LookupTXT(hostname string) []string {
	hostname = URLToHostname(hostname)
	txt, txtErr := net.LookupTXT(hostname)
	if txtErr != nil {
		panic(txtErr)
	}
	return txt
}

func FetchAll(hostname string) HostInfo {
	hostname = URLToHostname(hostname)
	hostinfo := HostInfo{}
	hostinfo.Hostname = hostname
	hostinfo.IPAddress = ResolveIPAddress(hostname)
	hostinfo.MX = LookupMX(hostname)
	hostinfo.CNAME = LookupCNAME(hostname)
	hostinfo.Nameserver = LookupNameserver(hostname)
	hostinfo.TXT = LookupTXT(hostname)

	return hostinfo
}

func FetchAsOutput(hostname string, f filter.Filter) (output.Output, output.Output) {
	info := FetchAll(hostname)

	o := output.Output{
		Hostname:      hostname,
		Status:        0,
		Time:          0,
		Destination:   "",
		ContentLength: 0,
		IpAddress:     "",
		MXRecords:     info.MX,
		ICMPCode:      0,
	}

	of := o.Filter(f)
	return o, of
}
