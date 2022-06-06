package hostinfo

import (
	"kulana/misc"
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
	misc.Check(err)
	hostinfo.IPAddress = ip.String()

	mx, mxErr := net.LookupMX(hostname)
	misc.Check(mxErr)
	var mxEntries []string
	for _, m := range mx {
		mxEntries = append(mxEntries, strings.TrimSuffix(m.Host, "."))
	}
	hostinfo.MX = mxEntries

	cname, cnameErr := net.LookupCNAME(hostname)
	misc.Check(cnameErr)
	hostinfo.CNAME = strings.TrimSuffix(cname, ".")

	nameservers, nsErr := net.LookupNS(hostname)
	misc.Check(nsErr)
	var nsEntries []string
	for _, ns := range nameservers {
		nsEntries = append(nsEntries, strings.TrimSuffix(ns.Host, "."))
	}
	hostinfo.Nameserver = nsEntries

	txt, txtErr := net.LookupTXT(hostname)
	misc.Check(txtErr)
	hostinfo.TXT = txt

	return hostinfo
}
