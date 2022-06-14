package _template

const FormatJSON = "json"
const FormatCSV = "csv"

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

func Render(t string, output _filter.Output) string {
	url = output.Url
	hostname = output.Hostname
	port = output.Port
	status = output.Status
	time = output.Time
	destination = output.Destination
	contentLength = output.ContentLength
	ipAddress = output.IpAddress
	mxRecords = output.MXRecords
	icmpCode = output.ICMPCode

	switch t {
	case "json":
		return RenderJSON()

	case "csv":
		return RenderCSV()

	default:
		return RenderDefault()
	}
}
