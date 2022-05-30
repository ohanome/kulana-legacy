package main

import (
	"fmt"
	"kulana/fetcher"
	"kulana/misc"
	"kulana/setup"
	"kulana/template"
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
var only = ""

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
	client := fetcher.CreateClient()
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

	response := fetcher.Response{
		Url:           url,
		Status:        statusCode,
		Time:          responseTimeRounded,
		Destination:   destination,
		ContentLength: contentLength,
	}

	switch only {
	case "url":
		response.Status = 0
		response.Time = 0
		response.Destination = ""
		response.ContentLength = -1
		break

	case "status":
		response.Url = ""
		response.Time = 0
		response.Destination = ""
		response.ContentLength = -1
		break

	case "time":
		response.Url = ""
		response.Status = 0
		response.Destination = ""
		response.ContentLength = -1
		break
	}

	t := template.Render(format, response)
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
