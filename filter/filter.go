package filter

import (
	"kulana/misc"
	"kulana/options"
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

func FromOptions() Filter {
	o, _, _ := options.Parse()
	return Filter{
		Url:           misc.StringInSlice("url", o.Include),
		Status:        misc.StringInSlice("status", o.Include),
		Time:          misc.StringInSlice("time", o.Include),
		Destination:   misc.StringInSlice("destination", o.Include),
		ContentLength: misc.StringInSlice("content_length", o.Include),
		IpAddress:     misc.StringInSlice("ip_address", o.Include),
		MXRecords:     misc.StringInSlice("mx_records", o.Include),
		ICMPCode:      misc.StringInSlice("icmp_code", o.Include),
		Hostname:      misc.StringInSlice("hostname", o.Include),
		Port:          misc.StringInSlice("port", o.Include),
		Content:       misc.StringInSlice("content", o.Include),
		ForeignID:     misc.StringInSlice("foreign_id", o.Include),
	}
}
