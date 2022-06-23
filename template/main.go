package template

import (
	"kulana/command"
	"kulana/output"
)

var url string
var hostname string
var port int
var status int
var time float64
var destination string
var contentLength int64
var ipAddress string
var mxRecords []string
var icmpCode int
var content string
var foreignId string

var opt command.DefaultOptions

func Render(t string, o output.Output, opts command.DefaultOptions) string {
	opt = opts

	url = o.Url
	hostname = o.Hostname
	port = o.Port
	status = o.Status
	time = o.Time
	destination = o.Destination
	contentLength = o.ContentLength
	ipAddress = o.IpAddress
	mxRecords = o.MXRecords
	icmpCode = o.ICMPCode
	content = o.Content
	foreignId = o.ForeignID

	switch t {
	case "json":
		return RenderJSON()

	case "csv":
		return RenderCSV()

	default:
		return RenderDefault()
	}
}
