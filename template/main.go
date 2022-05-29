package template

var url string
var status int
var time float64
var destination string
var contentLength int64

func Render(t string, _url string, _status int, _time float64, _destination string, _contentLength int64) string {
	url = _url
	status = _status
	time = _time
	destination = _destination
	contentLength = _contentLength

	switch t {
	case "json":
		return RenderJSON()

	case "csv":
		return RenderCSV()

	default:
		return RenderDefault()
	}
}
