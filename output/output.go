package output

import "kulana/filter"

type Output struct {
	Url            string   `json:"url,omitempty"`
	Status         int      `json:"status,omitempty"`
	Time           float64  `json:"time,omitempty"`
	Destination    string   `json:"destination,omitempty"`
	ContentLength  int64    `json:"content_length,omitempty"`
	IpAddress      string   `json:"ip_address,omitempty"`
	MXRecords      []string `json:"mx_records,omitempty"`
	ICMPCode       int      `json:"icmp_code,omitempty"`
	PingSuccessful int      `json:"ping_successful,omitempty"`
	PingError      string   `json:"ping_error,omitempty"`
	Hostname       string   `json:"hostname,omitempty"`
	CNAME          string   `json:"cname,omitempty"`
	Port           int      `json:"port,omitempty"`

	// The fetched page content.
	Content string `json:"content,omitempty"`

	// ForeignID is being used by foreign systems only.
	// For example, let's say there's a Drupal system where the output needs to be sent to and let's assume there are
	// two entity types implemented which are used to manage the contents and its relations:
	// - Output; an entity which holds the data represented by this struct and
	// - Page; an entity which is used to build the relationship between Outputs for the same URL (1:M Page:Output)
	// In this case the ForeignID is the ID of the parent Page entity.
	ForeignID   string            `json:"foreign_id,omitempty"`
	Certificate CertificateOutput `json:"certificate,omitempty"`
	Pings       []PingOutput      `json:"pings,omitempty"`
}

type CertificateOutput struct {
	// Instead of using a bool here, we will determine the validity by the following ints
	// - -1: not checked
	// - 0: not valid
	// - 1: valid
	Valid      bool   `json:"valid,omitempty"`
	ValidUntil string `json:"valid_until,omitempty"`
	Issuer     string `json:"issuer,omitempty"`
}

type PingOutput struct {
	Successful bool    `json:"successful"`
	Error      string  `json:"error"`
	Time       float64 `json:"time"`
	Port       int     `json:"port"`
}

func (o Output) Filter(f filter.Filter) Output {
	if !f.Url {
		o.Url = ""
	}

	if !f.Status {
		o.Status = 0
	}

	if !f.Time {
		o.Time = 0
	}

	if !f.Destination {
		o.Destination = ""
	}

	if !f.ContentLength {
		o.ContentLength = 0
	}

	if !f.IpAddress {
		o.IpAddress = ""
	}

	if !f.MXRecords {
		o.MXRecords = []string{}
	}

	if !f.ICMPCode {
		o.ICMPCode = 0
	}

	if !f.PingSuccessful {
		o.PingSuccessful = 0
	}

	if !f.PingError {
		o.PingError = ""
	}

	if !f.CNAME {
		o.CNAME = ""
	}

	if !f.Hostname {
		o.Hostname = ""
	}

	if !f.Port {
		o.Port = 0
	}

	if !f.Content {
		o.Content = ""
	}

	if !f.ForeignID {
		o.ForeignID = ""
	}

	if !f.Certificate.Valid {
		o.Certificate.Valid = false
	}

	if !f.Certificate.ValidUntil {
		o.Certificate.ValidUntil = ""
	}

	if !f.Certificate.Issuer {
		o.Certificate.Issuer = ""
	}

	return o
}

func (o Output) filterUrl(f filter.Filter) Output {
	if !f.Url {
		o.Url = ""
	}
	return o
}
