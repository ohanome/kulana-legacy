package misc

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"time"
)

func MicroTime() float64 {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	micSeconds := float64(now.Nanosecond()) / 1000000000
	return float64(now.Unix()) + micSeconds
}

func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Die(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func Ternary(condition bool, trueValue any, falseValue any) any {
	if condition {
		return trueValue
	}

	return falseValue
}

func GetNLChar() string {
	nl := "\n"

	if //goland:noinspection ALL
	runtime.GOOS == "windows" {
		nl = "\r\n"
	}

	return nl
}

func Usage() {
	fmt.Println("kulana v" + Version + GetNLChar() +
		"" + GetNLChar() +
		"A tool to request any HTTP host and get its status code, response time and other information." + GetNLChar() +
		"The return value will always contain the called URL, the HTTP status code of the response and the response time in milliseconds." + GetNLChar() +
		"" + GetNLChar() +
		"Usage" + GetNLChar() +
		"  kulana [...args]" + GetNLChar() +
		"" + GetNLChar() +
		"Possible arguments" + GetNLChar() +
		"  http...                   - The URL to request; must start with 'http'" + GetNLChar() +
		"  -h | --help               - This usage" + GetNLChar() +
		"  --json                    - Format the output as JSON" + GetNLChar() +
		"  --csv                     - Format the output as CSV" + GetNLChar() +
		"  --loop                    - Keeps sending requests" + GetNLChar() +
		"  --delay=n                 - Wait n milliseconds after each request; works only in combination with '--loop'; doesn't work with '-f'" + GetNLChar() +
		"  -f | --follow-redirect    - Sends another request if the response contains a Location header and a 3xx status code; doesn't work with '--loop'" + GetNLChar() +
		"  -l | --include-length     - Includes the content length" + GetNLChar() +
		"  --url-only                - Outputs only the URL (-l will be ignored)" + GetNLChar() +
		"  --time-only               - Outputs only the response time in milliseconds (-l will be ignored)" + GetNLChar() +
		"  --status-only             - Outputs only the HTTP status (-l will be ignored)" + GetNLChar() +
		"" + GetNLChar() +
		"Examples" + GetNLChar() +
		"  kulana https://ohano.me               - To get the HTTP status and the response time of https://ohano.me" + GetNLChar() +
		"  kulana https://ohano.me --loop        - Same as above, but the request will be sent every second until the program will be stopped" + GetNLChar() +
		"  kulana https://ohano.me --loop -f     - Will result in an error message since you can't follow redirects in a loop (yet)")
}
