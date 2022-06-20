package hostinfo

import (
	"kulana/filter"
	"kulana/l"
	"kulana/output"
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
	if err != nil {
		l.Emergency(err.Error())
	}
	hostinfo.IPAddress = ip.String()

	mx, mxErr := net.LookupMX(hostname)
	if mxErr != nil {
		l.Emergency(mxErr.Error())
	}
	var mxEntries []string
	for _, m := range mx {
		mxEntries = append(mxEntries, strings.TrimSuffix(m.Host, "."))
	}
	hostinfo.MX = mxEntries

	cname, cnameErr := net.LookupCNAME(hostname)
	if cnameErr != nil {
		l.Emergency(cnameErr.Error())
	}
	hostinfo.CNAME = strings.TrimSuffix(cname, ".")

	nameservers, nsErr := net.LookupNS(hostname)
	if nsErr != nil {
		l.Emergency(nsErr.Error())
	}
	var nsEntries []string
	for _, ns := range nameservers {
		nsEntries = append(nsEntries, strings.TrimSuffix(ns.Host, "."))
	}
	hostinfo.Nameserver = nsEntries

	txt, txtErr := net.LookupTXT(hostname)
	if txtErr != nil {
		l.Emergency(txtErr.Error())
	}
	hostinfo.TXT = txt

	return hostinfo
}

func FetchAsOutput(hostname string, f filter.Filter) (output.Output, output.Output) {
	info := Fetch(hostname)

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
