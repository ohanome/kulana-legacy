package _template

import (
	"fmt"
	"strconv"
	"strings"
)

func RenderJSON() string {
	var template []string

	if url != "" {
		template = append(template, "\"url\":\""+url+"\"")
	}

	if hostname != "" {
		template = append(template, "\"hostname\":\""+hostname+"\"")
	}

	if port > 0 {
		template = append(template, "\"port\":\""+strconv.Itoa(port)+"\"")
	}

	if ipAddress != "" {
		template = append(template, "\"ip\":\""+ipAddress+"\"")
	}

	if status != 0 {
		template = append(template, "\"status\":\""+strconv.Itoa(status)+"\"")
	}

	if time != 0 {
		template = append(template, "\"time\":\""+fmt.Sprintf("%f", time)+"\"")
	}

	if destination != "" && status < 400 && status >= 300 {
		template = append(template, "\"destination\":\""+destination+"\"")
	}

	if contentLength > -1 {
		template = append(template, "\"length\":\""+fmt.Sprintf("%d", contentLength)+"\"")
	}

	if len(mxRecords) > 0 {
		var mx []string
		for _, m := range mxRecords {
			mx = append(mx, m)
		}

		template = append(template, fmt.Sprintf("\"mx_records\":[\"%s\"]", strings.Join(mx, "\",\"")))
	}

	if icmpCode >= 0 {
		template = append(template, fmt.Sprintf("\"icmp_code\":%d", icmpCode))
	}

	t := "{" + strings.Join(template, ",") + "}"

	return fmt.Sprint(t)
}
