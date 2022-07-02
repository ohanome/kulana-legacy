package template

import (
	"fmt"
	"kulana/format"
	"strconv"
	"strings"
)

func RenderDefault() string {
	var template []string

	if url != "" {
		if !noColor {
			url = format.Blue + url + format.Reset
		}
		template = append(template, url)
	}

	if hostname != "" {
		template = append(template, hostname)
	}

	if port > 0 {
		p := fmt.Sprintf("%d", port)
		if !noColor {
			p = format.Purple + strconv.Itoa(port) + format.Reset
		}
		template = append(template, p)
	}

	if ipAddress != "" {
		template = append(template, ipAddress)
	}

	if status != 0 {
		s := fmt.Sprintf("%d", status)
		if !noColor {
			if status < 300 {
				s = format.Green + strconv.Itoa(status) + format.Reset
			} else if status < 400 {
				s = format.Yellow + strconv.Itoa(status) + format.Reset
			} else {
				s = format.Red + strconv.Itoa(status) + format.Reset
			}
		}
		template = append(template, s)
	}

	if time > -1 {
		t := fmt.Sprintf("%f", time)
		if !noColor {
			if time < 200 {
				t = format.Green + fmt.Sprintf("%f", time) + format.Reset
			} else if time < 1000 {
				t = format.Yellow + fmt.Sprintf("%f", time) + format.Reset
			} else {
				t = format.Red + fmt.Sprintf("%f", time) + format.Reset
			}
		}
		template = append(template, t)
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

		template = append(template, fmt.Sprintf("[%s]", strings.Join(mx, ", ")))
	}

	if icmpCode >= 0 {
		template = append(template, fmt.Sprintf("%d", icmpCode))
	}

	if pingSuccessful > -1 {
		success := format.Red + "Closed" + format.Reset
		if pingSuccessful == 1 {
			success = format.Green + "Open" + format.Reset
		}
		template = append(template, success)
	}

	if pingError != "" {
		template = append(template, fmt.Sprintf("Ping error: '%s'", pingError))
	}

	if cname != "" {
		template = append(template, cname)
	}

	if foreignId != "" {
		template = append(template, fmt.Sprintf("%v", foreignId))
	}

	if certificateValid == true {
		template = append(template, "Certificate Valid")
	}

	if certificateValidUntil != "" {
		template = append(template, fmt.Sprintf("Certificate Valid Until: %v", certificateValidUntil))
	}

	if certificateIssuer != "" {
		template = append(template, fmt.Sprintf("Certificate Issuer: %v", certificateIssuer))
	}

	t := strings.Join(template, "\t")

	return fmt.Sprint(t)
}
