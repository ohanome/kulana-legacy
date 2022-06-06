package app

import (
	"fmt"
	"kulana/hostinfo"
)

func runMX(application Application) {
	info := hostinfo.Fetch(application.Host)
	fmt.Println(info)
}
