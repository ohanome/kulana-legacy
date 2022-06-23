package main

import (
	"fmt"
	_ "kulana/command"
	"kulana/l"
	_ "kulana/l"
	"kulana/misc"
	"kulana/setup"
	"log"
	"os"
	"strings"
)

var commands = []string{"status"}

// Main entrypoint
func main() {
	misc.Welcome()
	setup.EnsureEnvironmentIsReady()

	l.Debug(3, fmt.Sprintf("Program arguments: %#v", os.Args))

	if len(os.Args) == 1 {
		log.Println("")
		l.Error(fmt.Sprintf("Missing command. Try one of the following: %s", strings.Join(commands, ", ")))
		log.Println("")
		os.Exit(1)
	}

	if len(os.Args) > 1 && !misc.StringInSlice(os.Args[1], commands) {
		l.Error(fmt.Sprintf("Unknown command '%s', try one of the following: %s", os.Args[1], strings.Join(commands, ", ")))
	}

	//command := os.Args[1]
	//switch command {
	//case "status":
	//	if o.Url == "" {
	//		l.Emergency("URL cannot be empty.")
	//	}
	//	l.Notice(fmt.Sprintf("Issued command 'status' with URL '%s', format '%s', in loop: %v", o.Url, o.Format, o.Loop))
	//	for {
	//		out := fetcher.FetchHTTPHost(o.Url, o.ForeignId)
	//		l.Debug(3, fmt.Sprintf("Output: %#v", out))
	//		f := filter.FromOptions(o)
	//		// Apply defaults
	//		f.Url = true
	//		f.Status = true
	//		f.Time = true
	//		out = out.Filter(f)
	//		l.Debug(2, fmt.Sprintf("Output (filtered): %#v", out))
	//		//formatted := template.Render(o.Format, out)
	//		//l.Debug(1, fmt.Sprintf("Output (formatted): '%s'", formatted))
	//		//l.Notice(fmt.Sprintf("Result: %s", formatted))
	//		//
	//		//fmt.Println(formatted)
	//		if !o.Loop {
	//			break
	//		}
	//		delay := o.Delay
	//		if delay < 1000 {
	//			delay = 1000
	//		}
	//		time.Sleep(time.Duration(delay) * time.Millisecond)
	//	}
	//	break
	//}
}
