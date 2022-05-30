package main

import "fmt"

func usage() {
	fmt.Println("kulana v" + version + nl +
		"" + nl +
		"A tool to request any HTTP host and get its status code, response time and other information." + nl +
		"The return value will always contain the called URL, the HTTP status code of the response and the response time in milliseconds." + nl +
		"" + nl +
		"Usage" + nl +
		"  kulana [...args]" + nl +
		"" + nl +
		"Possible arguments" + nl +
		"  http...                   - The URL to request; must start with 'http'" + nl +
		"  -h | --help               - This usage" + nl +
		"  --json                    - Format the output as JSON" + nl +
		"  --csv                     - Format the output as CSV" + nl +
		"  --loop                    - Keeps sending requests" + nl +
		"  --delay=n                 - Wait n milliseconds after each request; works only in combination with '--loop'; doesn't work with '-f'" + nl +
		"  -f | --follow-redirect    - Sends another request if the response contains a Location header and a 3xx status code; doesn't work with '--loop'" + nl +
		"  -l | --include-length     - Includes the content length" + nl +
		"  --url-only                - Outputs only the URL (-l will be ignored)" + nl +
		"  --time-only               - Outputs only the response time in milliseconds (-l will be ignored)" + nl +
		"  --status-only             - Outputs only the HTTP status (-l will be ignored)" + nl +
		"" + nl +
		"Examples" + nl +
		"  kulana https://ohano.me               - To get the HTTP status and the response time of https://ohano.me" + nl +
		"  kulana https://ohano.me --loop        - Same as above, but the request will be sent every second until the program will be stopped" + nl +
		"  kulana https://ohano.me --loop -f     - Will result in an error message since you can't follow redirects in a loop (yet)")
}
