package fetcher

import (
	"strconv"
	"strings"
)

const FilterUrl = 1
const FilterStatus = 2
const FilterTime = 4
const FilterDestination = 8
const FilterContentLength = 16

type ResponseFilter struct {
	Url           bool
	Status        bool
	Time          bool
	Destination   bool
	ContentLength bool
}

func BuildFilterFromBoolean(filterUrl bool, filterStatus bool, filterTime bool, filterDestination bool, filterContentLength bool) ResponseFilter {
	filter := ResponseFilter{
		Url:           filterUrl,
		Status:        filterStatus,
		Time:          filterTime,
		Destination:   filterDestination,
		ContentLength: filterContentLength,
	}

	return filter
}

func BuildFilterFromNumeric(settings int64) ResponseFilter {
	settingsBinary := strconv.FormatInt(settings, 2)

	filters := []bool{false, false, false, false, false}

	for index, element := range strings.Split(settingsBinary, "") {
		if element == "1" {
			filters[index] = true
		}
	}

	filter := ResponseFilter{
		Url:           filters[0],
		Status:        filters[1],
		Time:          filters[2],
		Destination:   filters[3],
		ContentLength: filters[4],
	}

	return filter
}

func FilterResponse(response Response, filter ResponseFilter) Response {
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

	return response
}
