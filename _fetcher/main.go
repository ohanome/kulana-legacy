package _fetcher

import (
	"io"
	"io/ioutil"
	"kulana/_filter"
	"kulana/_misc"
	"net/http"
)

func CreateHTTPClient() *http.Client {
	return &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}

func FetchHTTPHost(url string, foreignId string) _filter.Output {
	client := CreateHTTPClient()
	start := _misc.MicroTime()
	resp, err := client.Get(url)
	end := _misc.MicroTime()
	defer client.CloseIdleConnections()
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	_misc.Check(err)

	statusCode := resp.StatusCode
	responseTime := (end - start) * 1000
	responseTimeRounded := _misc.Round(responseTime, 0.000001)

	var destination string
	location, err := resp.Location()
	if err != nil {
		destination = url
	} else {
		destination = location.String()
	}

	contentLength := resp.ContentLength
	body, err := ioutil.ReadAll(resp.Body)

	response := _filter.Output{
		Url:           url,
		Status:        statusCode,
		Time:          responseTimeRounded,
		Destination:   destination,
		ContentLength: contentLength,
		Content:       string(body),
		ForeignID:     foreignId,
	}

	return response
}

func FetchAndFilter(url string, f _filter.OutputFilter, foreignId string) (_filter.Output, _filter.Output) {
	response := FetchHTTPHost(url, foreignId)
	filteredResponse := _filter.FilterOutput(response, f)
	return response, filteredResponse
}
