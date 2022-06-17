package output

import "kulana/filter"

type Output struct {
	Url           string
	Status        int
	Time          float64
	Destination   string
	ContentLength int64
	IpAddress     string
	MXRecords     []string
	ICMPCode      int
	Hostname      string
	Port          int

	// The fetched page content.
	Content string

	// ForeignID is being used by foreign systems only.
	// For example, let's say there's a Drupal system where the output needs to be sent to and let's assume there are
	// two entity types implemented which are used to manage the contents and its relations:
	// - Output; an entity which holds the data represented by this struct and
	// - Page; an entity which is used to build the relationship between Outputs for the same URL (1:M Page:Output)
	// In this case the ForeignID is the ID of the parent Page entity.
	ForeignID string
}

func (o Output) Filter(f filter.Filter) Output {
	if !f.Url {
		o.Url = ""
	}

	if !f.Status {
		o.Status = 0
	}

	if !f.Time {
		o.Time = -1
	}

	if !f.Destination {
		o.Destination = ""
	}

	if !f.ContentLength {
		o.ContentLength = -1
	}

	if !f.IpAddress {
		o.IpAddress = ""
	}

	if !f.MXRecords {
		o.MXRecords = []string{}
	}

	if !f.ICMPCode {
		o.ICMPCode = -1
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

	return o
}

func (o Output) filterUrl(f filter.Filter) Output {
	if !f.Url {
		o.Url = ""
	}
	return o
}
