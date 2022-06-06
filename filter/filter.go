package filter

import (
	"strconv"
	"strings"
)

const Url = 1
const Status = 2
const Time = 4
const Destination = 8
const ContentLength = 16
const IpAddress = 32

type OutputFilter struct {
	Url           bool
	Status        bool
	Time          bool
	Destination   bool
	ContentLength bool
	IpAddress     bool
}

type Output struct {
	Url           string
	Status        int
	Time          float64
	Destination   string
	ContentLength int64
	IpAddress     string
}

func BuildFilterFromBoolean(filterUrl bool, filterStatus bool, filterTime bool, filterDestination bool, filterContentLength bool, filterIpAddress bool) OutputFilter {
	filter := OutputFilter{
		Url:           filterUrl,
		Status:        filterStatus,
		Time:          filterTime,
		Destination:   filterDestination,
		ContentLength: filterContentLength,
		IpAddress:     filterIpAddress,
	}

	return filter
}

func BuildFilterFromNumeric(settings int64) OutputFilter {
	settingsBinary := strconv.FormatInt(settings, 2)

	filters := []bool{false, false, false, false, false, false}

	for index, element := range strings.Split(settingsBinary, "") {
		if element == "1" {
			filters[index] = true
		}
	}

	filter := OutputFilter{
		Url:           filters[0],
		Status:        filters[1],
		Time:          filters[2],
		Destination:   filters[3],
		ContentLength: filters[4],
		IpAddress:     filters[5],
	}

	return filter
}

func FilterOutput(response Output, filter OutputFilter) Output {
	if !filter.Url {
		response.Url = ""
	}

	if !filter.Status {
		response.Status = 0
	}

	if !filter.Time {
		response.Time = 0
	}

	if !filter.Destination {
		response.Destination = ""
	}

	if !filter.ContentLength {
		response.ContentLength = -1
	}

	if !filter.IpAddress {
		response.IpAddress = ""
	}

	return response
}
