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
		if !opt.NoColor {
			url = format.Blue + url + format.Reset
		}
		template = append(template, url)
	}

	if hostname != "" {
		template = append(template, hostname)
	}

	if port > 0 {
		p := fmt.Sprintf("%d", port)
		if !opt.NoColor {
			p = format.Purple + strconv.Itoa(port) + format.Reset
		}
		template = append(template, p)
	}

	if ipAddress != "" {
		template = append(template, ipAddress)
	}

	if status != 0 {
		s := fmt.Sprintf("%d", status)
		if !opt.NoColor {
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

	if time != 0 {
		t := fmt.Sprintf("%f", time)
		if !opt.NoColor {
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

	if foreignId != "" {
		template = append(template, fmt.Sprintf("%v", foreignId))
	}

	t := strings.Join(template, "\t")

	return fmt.Sprint(t)
}
