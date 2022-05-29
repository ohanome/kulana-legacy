package template

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

	t := "{" + strings.Join(template, ",") + "}"

	return fmt.Sprint(t)
}
