package template

import (
	"kulana/filter"
)

const FormatJSON = "json"
const FormatCSV = "csv"

var url string
var status int
var time float64
var destination string
var contentLength int64
var ipAddress string

func Render(t string, output filter.Output) string {
	url = output.Url
	status = output.Status
	time = output.Time
	destination = output.Destination
	contentLength = output.ContentLength
	ipAddress = output.IpAddress

	switch t {
	case "json":
		return RenderJSON()

	case "csv":
		return RenderCSV()

	default:
		return RenderDefault()
	}
}
