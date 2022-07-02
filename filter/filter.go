package filter

type Filter struct {
	Url            bool
	Status         bool
	Time           bool
	Destination    bool
	ContentLength  bool
	IpAddress      bool
	MXRecords      bool
	ICMPCode       bool
	PingSuccessful bool
	PingError      bool
	Hostname       bool
	CNAME          bool
	Port           bool
	Content        bool
	ForeignID      bool
	Certificate    CertificateFilter
}

type CertificateFilter struct {
	Valid      bool
	ValidUntil bool
	Issuer     bool
}
