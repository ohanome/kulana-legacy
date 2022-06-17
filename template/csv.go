package template

import (
	"fmt"
	"strconv"
	"strings"
)

func RenderCSV() string {
	var template []string

	if url != "" {
		template = append(template, url)
	}

	if hostname != "" {
		template = append(template, hostname)
	}

	if port > 0 {
		template = append(template, strconv.Itoa(port))
	}

	if ipAddress != "" {
		template = append(template, ipAddress)
	}

	if status != 0 {
		template = append(template, strconv.Itoa(status))
	}

	if time != 0 {
		template = append(template, fmt.Sprintf("%f", time))
	}

	if destination != "" && status < 400 && status >= 300 {
		template = append(template, destination)
	}

	if contentLength > -1 {
		template = append(template, fmt.Sprintf("%d", contentLength))
	}

	if len(mxRecords) > 0 {
		var mx []string
		for _, m := range mxRecords {
			mx = append(mx, m)
		}

		template = append(template, fmt.Sprintf("[%s]", strings.Join(mx, " ")))
	}

	if icmpCode >= 0 {
		template = append(template, fmt.Sprintf("%d", icmpCode))
	}

	// Content will not be output for now.
	//if content != "" {
	//	template = append(template, fmt.Sprintf("\"\"\"%s\"\"\"", content))
	//}

	if foreignId != "" {
		template = append(template, fmt.Sprintf("%v", foreignId))
	}

	t := strings.Join(template, ",")

	return fmt.Sprint(t)
}
