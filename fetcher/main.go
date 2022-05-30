package fetcher

import (
	"kulana/misc"
	"net/http"
)

type Response struct {
	Url           string
	Status        int
	Time          float64
	Destination   string
	ContentLength int64
}

func CreateHTTPClient() *http.Client {
	return &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}

func FetchHTTPHost(url string) Response {
	client := CreateHTTPClient()
	start := misc.MicroTime()
	resp, err := client.Get(url)
	defer client.CloseIdleConnections()
	misc.Check(err)

	statusCode := resp.StatusCode
	responseTime := (misc.MicroTime() - start) * 1000
	responseTimeRounded := misc.Round(responseTime, 0.000001)

	var destination string
	location, err := resp.Location()
	if err != nil {
		destination = url
	} else {
		destination = location.String()
	}

	contentLength := resp.ContentLength

	response := Response{
		Url:           url,
		Status:        statusCode,
		Time:          responseTimeRounded,
		Destination:   destination,
		ContentLength: contentLength,
	}

	return response
}

func FetchAndFilter(url string, filter ResponseFilter) (Response, Response) {
	response := FetchHTTPHost(url)
	filteredResponse := FilterResponse(response, filter)
	return response, filteredResponse
}
