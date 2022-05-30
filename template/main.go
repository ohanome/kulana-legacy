package template

import "kulana/fetcher"

var url string
var status int
var time float64
var destination string
var contentLength int64

func Render(t string, response fetcher.Response) string {
	url = response.Url
	status = response.Status
	time = response.Time
	destination = response.Destination
	contentLength = response.ContentLength

	switch t {
	case "json":
		return RenderJSON()

	case "csv":
		return RenderCSV()

	default:
		return RenderDefault()
	}
}
