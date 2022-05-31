# 📊 kulana (_kūlana_)

Kūlana means "status" in hawaiian[*](https://hilo.hawaii.edu/wehe/?q=kulana#w2w2-10743) and that's all this project is about. With this tool on your hand you can fetch the status of nearly every single website, including information about the HTTP status, the response time and the content length.

Possible use cases may be to monitor your hosts or to get if there are any redirections (HTTP status 3xx) or just flex on how small your website is.

## 🛠 Build

**Prerequisites**
- [Go](https://go.dev/doc/install) version >= 1.18

To fetch the projects dependencies, run
```shell
go mod vendor
```

After that you can build the app for your platform by running
```shell
go build
```

## 📦 Installation

**Unix/Mac**

You can run the provided script `install.sh` by executing 
```shell
./install.sh
```

Optionally, you can execute the steps done by the script by yourself:

1. Build the application by running `go build`
2. Copy the built binary to `/usr/bin/kulana` to make it globally available (or `~/bin/kulana` to install it only for the active user)
3. Make sure the directory where you copied the binary to is in the $PATH variable

**Windows**

To install the app in your default go binary directory (created when installing Go on Windows) run 
```shell
go install
```

You should immediately have access to the `kulana.exe` command from the command line.

Alternatively you can build the app by yourself by running 
```shell
go build
```
and move the binary to your preferred destination.

## 🔥 Usage

This app is separated in several commands, like `help` or `status`.

For every command the schema is the following:
```shell
kulana COMMAND [...args]
```

**TOC**
- [Help](#help)
- [Status](#status)

### Help

The base usage is
```shell
kulana help
```

This will show you all available commands you can use.

Since every command provides its own help section you can safely execute 
```shell
kulana COMMAND --help
```
to get information about that command.

### Status

The base usage is 
```shell
kulana status URL
```
where `URL` is a valid URL.

**TOC**
- [Output format](#output-format)
- [Looping](#looping)
- [Redirect URLS](#redirect-urls)
- [Content length](#content-length)
- [Environment setup](#environment-setup)
- [Sending notification mails](#sending-notification-mails)

**Summary**
```text
Usage
  kulana [...args]

Possible arguments
  http...                   - The URL to request; must start with 'http'
  -h | --help               - This usage
  --json                    - Format the output as JSON
  --csv                     - Format the output as CSV
  --loop                    - Keeps sending requests
  --delay=N                 - Wait N milliseconds after each request; works only in combination with '--loop'; doesn't work with '-f'
  -f | --follow-redirect    - Sends another request if the response contains a Location header and a 3xx status code; doesn't work with '--loop'
  -l | --include-length     - Includes the content length
  --url-only                - Outputs only the URL (-l will be ignored)
  --time-only               - Outputs only the response time in milliseconds (-l will be ignored)
  --status-only             - Outputs only the HTTP status (-l will be ignored)
  -n | --notify             - Sends an email with the status code to the given email address (--notify-mail needed). The environment will be checked before, so make sure you fill in all variables in ~/.kulana/.env
  --notify-mail=MAIL        - The address to send the email to
  --check-env               - Validates that all environment configurations are setup

Examples
  kulana https://ohano.me               - To get the HTTP status and the response time of https://ohano.me
  kulana https://ohano.me --loop        - Same as above, but the request will be sent every second until the program will be stopped
  kulana https://ohano.me --loop -f     - Will result in an error message since you can't follow redirects in a loop (yet)
```

#### Output format

The following formats are available by appending the corresponding argument:

| Format | Argument |
|--------|----------|
| CSV    | `--csv`  |
| JSON   | `--json` |

The last argument will be used, meaning that if you execute 
```shell
kulana URL --json --csv
```
the output will be formatted as CSV.

If no specific format is given, the default format will be used.

##### Default

The default format separates the single values by tabs, like so
```shell
https://ohano.me        200     104.208708
```

##### JSON

The JSON format contains all keys to every value, which may be more useful to work with in other systems.

Sample output:
```json
{
  "url":"https://ohano.me",
  "status":"200",
  "time":"91.022730"
}
```

##### CSV

The CSV format does not contain any keys, nly the "raw" values.

The order is the following:
```csv
URL,HTTP status,time(,redirect URL)(,content length)
```

Sample output:
```csv
https://ohano.me,200,96.359968
```

#### Looping

You can watch a hosts response by passing the `--loop` argument. With this kulana will keep sending requests after a defined delay.

The delay is by default 1000 ms, but you can modify it with the `--delay` argument. Example:
```shell
kulana URL --loop --delay=60000
```
Here the request will be resent every 60 seconds (60.000 milliseconds).

> **Attention when using the `--json` flag**
> 
> Until now the JSON formatter doesn't know if its in a loop or not. That means the output will be formatted in an invalid structure like the following:
> ```text
> {"url":"https://ohano.me","status":"200","time":"110.463381"}
> {"url":"https://ohano.me","status":"200","time":"63.969135"}
> {"url":"https://ohano.me","status":"200","time":"77.401638"}
> ```

#### Redirect URLS

If the status code of the response is between 300 and 399 and another URL is given by the `Location` header, the destination URL will be added to the output as `"destination"` (in CSV and default format the destination URL will be the 4th value).

You can send another request by using `--follow-redirect` (or `-f`).

> Kulana can't follow redirects in a loop (yet), so using `-f` and `--loop` will result in an error.

#### Content length

You can include the length of the responses content by passing `--include-length` (or `-l`). The content length will be added to the output without any unit, just as plain byte value.

#### Environment setup

On the first start an .env file will be created under `~/.kulana` (on Windows its mostly under `C:\.kulana`). This will holds several variables for sensitive information, most importantly for any login or access like SMTP authentication.

This environment is needed for the following features (list will be longer as time passes and development continues):

| Feature | Environment prefix |
|---------|--------------------|
| Mail    | `SMTP_`            |

#### Sending notification mails

You can email the HTTP status of every request to a specified address by passing 2 additional arguments.

The first is the "switch" that turns this feature on:
```shell
kulana https://ohano.me -n
```
or alternatively
```shell
kulana https://ohano.me --notify
```

The second argument contains the address the notification should be sent to:
```shell
kulana https://ohano.me -n --notify-mail=somemail@provider.tld
```

> Make sure, you have your environment set up, see [#environment-setup](#environment-setup)

## 🔮 Planned features

- [ ] Runnable as background task (like docker with commands like `start` and `stop`)
- [ ] Crawling functionality
- [ ] Sending result to an API
- [ ] Global configuration (via configuration file)
- [ ] Port ping option

## ⭐️ Usage examples

Fetch the time and HTTP-Status of the host `https://ohano.me`:
```shell
kulana https://ohano.me
```

Watch the status of the host `https://ohano.me`:
```shell
kulana https://ohano.me --loop
```

Watch the status of the host `https://ohano.me` but wait 60 seconds between every request:
```shell
kulana https://ohano.me --loop --delay=60000
```

Fetch the status of the host `https://ohano.me` as JSON:
```shell
kulana https://ohano.me --json
```

## 📜 License

This project is licensed under the GNU GPL v3, for more details see [license](./LICENSE).