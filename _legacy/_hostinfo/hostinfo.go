package _hostinfo

import (
	"net"
	"strings"
)

type HostInfo struct {
	Hostname   string
	IPAddress  string
	Nameserver []string
	MX         []string
	TXT        []string
	CNAME      string
}

func Fetch(hostname string) HostInfo {
	hostinfo := HostInfo{}
	hostinfo.Hostname = hostname

	ip, err := net.ResolveIPAddr("ip", hostname)
	_misc.Check(err)
	hostinfo.IPAddress = ip.String()

	mx, mxErr := net.LookupMX(hostname)
	_misc.Check(mxErr)
	var mxEntries []string
	for _, m := range mx {
		mxEntries = append(mxEntries, strings.TrimSuffix(m.Host, "."))
	}
	hostinfo.MX = mxEntries

	cname, cnameErr := net.LookupCNAME(hostname)
	_misc.Check(cnameErr)
	hostinfo.CNAME = strings.TrimSuffix(cname, ".")

	nameservers, nsErr := net.LookupNS(hostname)
	_misc.Check(nsErr)
	var nsEntries []string
	for _, ns := range nameservers {
		nsEntries = append(nsEntries, strings.TrimSuffix(ns.Host, "."))
	}
	hostinfo.Nameserver = nsEntries

	txt, txtErr := net.LookupTXT(hostname)
	_misc.Check(txtErr)
	hostinfo.TXT = txt

	return hostinfo
}

func FetchAsOutput(hostname string, f _filter.OutputFilter) (_filter.Output, _filter.Output) {
	info := Fetch(hostname)

	o := _filter.Output{
		Hostname:      hostname,
		Status:        0,
		Time:          0,
		Destination:   "",
		ContentLength: 0,
		IpAddress:     "",
		MXRecords:     info.MX,
		ICMPCode:      0,
	}

	of := _filter.FilterOutput(o, f)
	return o, of
}
