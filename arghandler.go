package main

import (
	"kulana/misc"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func processArgs() {
	if len(os.Args) == 1 {
		usage()
		os.Exit(1)
	}

	if len(os.Args) >= 1 {
		for _, arg := range os.Args[1:] {
			switch arg {
			case "--help":
			case "-h":
				usage()
				os.Exit(0)

			case "--json":
				format = "json"
				break

			case "--csv":
				format = "csv"
				break

			case "--loop":
				loop = true
				if followRedirect {
					misc.Die("Cannot follow redirects in a loop")
				}
				break

			case "--include-length":
			case "-l":
				getContentLength = true
				break

			case "--follow-redirect":
			case "-f":
				followRedirect = true
				if loop {
					misc.Die("Cannot follow redirects in a loop")
				}
				break

			default:
				delayMatch, _ := regexp.Match(`^--delay=\d.*$`, []byte(arg))
				if delayMatch {
					d := strings.ReplaceAll(arg, "--delay=", "")
					dInt, _ := strconv.ParseInt(d, 10, 64)
					delay = dInt * 1000000
				}

				urlMatch, _ := regexp.Match(`^http(s|)://\w.*\.\w{2,3}$`, []byte(arg))
				if urlMatch {
					url = arg
				}
				break
			}
		}

		if url == "" {
			misc.Die("No URL given.")
		}
	}
}
