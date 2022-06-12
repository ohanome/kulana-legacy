package fetcher

import (
	"io"
	"io/ioutil"
	"kulana/filter"
	"kulana/misc"
	"net/http"
)

func CreateHTTPClient() *http.Client {
	return &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}

func FetchHTTPHost(url string, foreignId string) filter.Output {
	client := CreateHTTPClient()
	start := misc.MicroTime()
	resp, err := client.Get(url)
	end := misc.MicroTime()
	defer client.CloseIdleConnections()
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	misc.Check(err)

	statusCode := resp.StatusCode
	responseTime := (end - start) * 1000
	responseTimeRounded := misc.Round(responseTime, 0.000001)

	var destination string
	location, err := resp.Location()
	if err != nil {
		destination = url
	} else {
		destination = location.String()
	}

	contentLength := resp.ContentLength
	body, err := ioutil.ReadAll(resp.Body)

	response := filter.Output{
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

func FetchAndFilter(url string, f filter.OutputFilter, foreignId string) (filter.Output, filter.Output) {
	response := FetchHTTPHost(url, foreignId)
	filteredResponse := filter.FilterOutput(response, f)
	return response, filteredResponse
}
