package template

import (
	"fmt"
	"strconv"
	"strings"
)

func RenderDefault() string {
	var template []string

	if url != "" {
		template = append(template, url)
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

	t := strings.Join(template, "\t")

	return fmt.Sprint(t)
}
