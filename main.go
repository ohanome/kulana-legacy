package main

import (
	"fmt"
	"kulana/misc"
	"kulana/setup"
	"kulana/template"
	"net/http"
	"os"
	"runtime"
	"time"
)

var format = ""
var nl = "\n"
var loop = false
var version = "0.0.1"
var url = ""
var getContentLength = false
var followRedirect = false

// 1000 ms as default
var delay int64 = 1000 * 1000000

func detectOS() {
	if //goland:noinspection ALL
	runtime.GOOS == "windows" {
		nl = "\r\n"
	}
}

func pingURL() {
	start := misc.MicroTime()
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
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

	var contentLength int64 = -1
	if getContentLength {
		contentLength = resp.ContentLength
	}

	t := template.Render(format, url, statusCode, responseTimeRounded, destination, contentLength)
	fmt.Print(t + nl)

	if followRedirect && !loop && statusCode < 400 && statusCode >= 300 {
		url = destination
		pingURL()
	}
}

// Main entrypoint
func main() {
	detectOS()
	processArgs()
	setup.EnsureEnvironmentIsReady()

	if !loop {
		pingURL()
		os.Exit(0)
	}

	for {
		pingURL()
		time.Sleep(time.Duration(delay))
	}
}
