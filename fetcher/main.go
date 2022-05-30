package fetcher

import "net/http"

type Response struct {
	Url           string
	Status        int
	Time          float64
	Destination   string
	ContentLength int64
}

func CreateClient() *http.Client {
	return &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}
