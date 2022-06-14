package _filter

var defaultAll = OutputFilter{
	Url:           true,
	Status:        true,
	Time:          true,
	Destination:   true,
	ContentLength: true,
	IpAddress:     true,
	MXRecords:     true,
	ICMPCode:      true,
	Hostname:      true,
	Port:          true,
	Content:       true,
	ForeignID:     true,
}

var defaultStatus = OutputFilter{
	Url:           true,
	Status:        true,
	Time:          true,
	Destination:   true,
	ContentLength: false,
	IpAddress:     false,
	MXRecords:     false,
	ICMPCode:      false,
	Hostname:      false,
	Port:          false,
	Content:       false,
	ForeignID:     true,
}

var defaultPing = OutputFilter{
	Url:           false,
	Status:        false,
	Time:          true,
	Destination:   false,
	ContentLength: false,
	IpAddress:     true,
	MXRecords:     false,
	ICMPCode:      true,
	Hostname:      true,
	Port:          true,
	Content:       false,
	ForeignID:     true,
}

var defaultMx = OutputFilter{
	Url:           false,
	Status:        false,
	Time:          false,
	Destination:   false,
	ContentLength: false,
	IpAddress:     false,
	MXRecords:     true,
	ICMPCode:      false,
	Hostname:      true,
	Port:          false,
	Content:       false,
	ForeignID:     true,
}

func GetDefault(command string) OutputFilter {
	switch command {
	case "status":
		return defaultStatus
	case "ping":
		return defaultPing
	case "mx":
		return defaultMx
	default:
		return defaultAll
	}
}
