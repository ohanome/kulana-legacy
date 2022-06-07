package misc

import (
	"fmt"
)

func Usage(command string) {
	switch command {
	case "help":
		usageHelp()
		break

	case "status":
		usageStatus()
		break

	case "config":
		usageConfig()
		break

	case "ping":
		usagePing()
		break

	case "mx":
		usageMX()
		break
	}

}

func usageHelp() {
	fmt.Println("kulana v" + Version + GetNLChar() +
		"" + GetNLChar() +
		"A tool for working with hosts and their responses." + GetNLChar() +
		"" + GetNLChar() +
		"Usage" + GetNLChar() +
		"  kulana [command]" + GetNLChar() +
		"" + GetNLChar() +
		"Possible commands" + GetNLChar() +
		"  help                      - This help" + GetNLChar() +
		"  status                    - The status tool" + GetNLChar() +
		"  config                    - Alter the configuration" + GetNLChar() +
		"" + GetNLChar() +
		"There are arguments which are globally available but will not effect the output of every command." + GetNLChar() +
		"Below is a list of all available arguments, followed by a table where you can see which argument is being used in which command." + GetNLChar() +
		"" + GetNLChar() +
		"Possible arguments" + GetNLChar() +
		"  -h | --help               - This usage" + GetNLChar() +
		"  --json                    - Format the output as JSON" + GetNLChar() +
		"  --csv                     - Format the output as CSV" + GetNLChar() +
		"  --loop                    - Keeps sending requests" + GetNLChar() +
		"  --delay=N                 - Wait N milliseconds after each request; works only in combination with '--loop'; doesn't work with '-f'" + GetNLChar() +
		"  -f | --follow-redirect    - Sends another request if the response contains a Location header and a 3xx status code; doesn't work with '--loop'" + GetNLChar() +
		"  -l | --include-length     - Includes the content length" + GetNLChar() +
		"  --url-only                - Outputs only the URL (-l will be ignored)" + GetNLChar() +
		"  --time-only               - Outputs only the response time in milliseconds (-l will be ignored)" + GetNLChar() +
		"  --status-only             - Outputs only the HTTP status (-l will be ignored)" + GetNLChar() +
		"  -n | --notify             - Sends an email with the status code to the given email address (--notify-mail needed). The environment will be checked before, so make sure you fill in all variables in ~/.kulana/.env" + GetNLChar() +
		"  --notify-mail=MAIL        - The address to send the email to" + GetNLChar() +
		"  --check-env               - Validates that all environment configurations are setup" + GetNLChar() +
		"" + GetNLChar() +
		"Argument support by command" + GetNLChar() +
		"|---------------------------|-----------|-----------|-----------|-----------|" + GetNLChar() +
		"| Argument                  | Status    | Ping      | MX Lookup | Config    |" + GetNLChar() +
		"|---------------------------|-----------|-----------|-----------|-----------|" + GetNLChar() +
		"| -h | --help               | supported | supported | supported | supported |" + GetNLChar() +
		"| --json                    | supported | supported | supported |           |" + GetNLChar() +
		"| --csv                     | supported | supported | supported |           |" + GetNLChar() +
		"| --loop                    | supported | supported |           |           |" + GetNLChar() +
		"| --delay=N                 | supported | supported |           |           |" + GetNLChar() +
		"| -f | --follow-redirect    | supported |           |           |           |" + GetNLChar() +
		"| -l | --include-length     | supported |           |           |           |" + GetNLChar() +
		"| --url-only                | supported |           |           |           |" + GetNLChar() +
		"| --time-only               | supported |           |           |           |" + GetNLChar() +
		"| --status-only             | supported |           |           |           |" + GetNLChar() +
		"| -n | --notify             | supported |           |           |           |" + GetNLChar() +
		"| --notify-mail=MAIL        | supported |           |           |           |" + GetNLChar() +
		"| --check-env               | supported | supported | supported | supported |" + GetNLChar() +
		"|---------------------------|-----------|-----------|-----------|-----------|" + GetNLChar() +
		"" + GetNLChar() +
		"For more detailed help type" + GetNLChar() +
		"  kulana [command] --help" + GetNLChar() +
		"Every single command provides its own help section with instructions and examples")
}

func usageStatus() {
	fmt.Println("kulana v" + Version + GetNLChar() +
		"" + GetNLChar() +
		"A tool to request any HTTP host and get its status code, response time and other information." + GetNLChar() +
		"The return value will always contain the called URL, the HTTP status code of the response and the response time in milliseconds." + GetNLChar() +
		"" + GetNLChar() +
		"Usage" + GetNLChar() +
		"  kulana status [...args]" + GetNLChar() +
		"" + GetNLChar() +
		"Possible arguments (+ globals, see kulana help for more details)" + GetNLChar() +
		"  http...                   - The URL to request; must start with 'http'" + GetNLChar() +
		"" + GetNLChar() +
		"Examples" + GetNLChar() +
		"  kulana status https://ohano.me               - To get the HTTP status and the response time of https://ohano.me" + GetNLChar() +
		"  kulana status https://ohano.me --loop        - Same as above, but the request will be sent every second until the program will be stopped" + GetNLChar() +
		"  kulana status https://ohano.me --loop -f     - Will result in an error message since you can't follow redirects in a loop (yet)")
}

func usageConfig() {
	fmt.Println("kulana v" + Version + GetNLChar() +
		"" + GetNLChar() +
		"Edit the configuration from the CLI." + GetNLChar() +
		"" + GetNLChar() +
		"Usage" + GetNLChar() +
		"  kulana config [get|set] ([key] ([value])|--all)" + GetNLChar() +
		"" + GetNLChar() +
		"Possible configs to reach" + GetNLChar() +
		"  mail                      - Mail configuration" + GetNLChar() +
		"  mail.status_codes         - The status codes where emails should be sent" + GetNLChar() +
		"" + GetNLChar() +
		"Examples" + GetNLChar() +
		"  kulana config get mail                       - Prints all configurations for mails" + GetNLChar() +
		"  kulana config set mail.status_codes 500      - Sets the values of mails.status_codes to \"500\"")
}

func usagePing() {
	fmt.Println("kulana v" + Version + GetNLChar() +
		"" + GetNLChar() +
		"Pings a host by sending an ICMP packet or checks if a specific port is open." + GetNLChar() +
		"" + GetNLChar() +
		"Usage" + GetNLChar() +
		"  kulana ping HOST [PORT] [...args]" + GetNLChar() +
		"" + GetNLChar() +
		"Possible arguments (+ globals, see kulana help for more details)" + GetNLChar() +
		"  HOST                      - Any hostname or IP (v4) address" + GetNLChar() +
		"  PORT                      - The port to ping, must be an integer" + GetNLChar() +
		"" + GetNLChar() +
		"Examples" + GetNLChar() +
		"  kulana ping ohano.me                         - Sends an ICMP packet to check if the given host is alive" + GetNLChar() +
		"  kulana ping ohano.me 25565 --json            - Sends a ping to ohano.me:25565 and formats its output as JSON. Will end with an error if port is not open")
}

func usageMX() {
	fmt.Println("kulana v" + Version + GetNLChar() +
		"" + GetNLChar() +
		"Fetches the MX records for a given domain." + GetNLChar() +
		"" + GetNLChar() +
		"Usage" + GetNLChar() +
		"  kulana mx HOST [...args]" + GetNLChar() +
		"" + GetNLChar() +
		"Possible arguments (+ globals, see kulana help for more details)" + GetNLChar() +
		"  HOST                      - Any hostname or IP (v4) address" + GetNLChar() +
		"" + GetNLChar() +
		"Examples" + GetNLChar() +
		"  kulana mx ohano.me                           - Returns an array of MX records")
}
