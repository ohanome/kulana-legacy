package filter

import (
	"kulana/misc"
)

type Filter struct {
	Url           bool
	Status        bool
	Time          bool
	Destination   bool
	ContentLength bool
	IpAddress     bool
	MXRecords     bool
	ICMPCode      bool
	Hostname      bool
	Port          bool
	Content       bool
	ForeignID     bool
}

func FromOptions(includes []string) Filter {
	return Filter{
		Url:           misc.StringInSlice("url", includes),
		Status:        misc.StringInSlice("status", includes),
		Time:          misc.StringInSlice("time", includes),
		Destination:   misc.StringInSlice("destination", includes),
		ContentLength: misc.StringInSlice("content_length", includes),
		IpAddress:     misc.StringInSlice("ip_address", includes),
		MXRecords:     misc.StringInSlice("mx_records", includes),
		ICMPCode:      misc.StringInSlice("icmp_code", includes),
		Hostname:      misc.StringInSlice("hostname", includes),
		Port:          misc.StringInSlice("port", includes),
		Content:       misc.StringInSlice("content", includes),
		ForeignID:     misc.StringInSlice("foreign_id", includes),
	}
}
