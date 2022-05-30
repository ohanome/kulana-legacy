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

_tbd_

## 🔮 Usage

The base usage is 
```shell
kulana URL
```
where `URL` is a valid URL.

**Summary**
```shell
Usage
  kulana [...args]

Possible arguments
  http...                   - The URL to request; must start with 'http'
  -h | --help               - This usage
  --json                    - Format the output as JSON
  --csv                     - Format the output as CSV
  --loop                    - Keeps sending requests
  --delay=n                 - Wait n milliseconds after each request; works only in combination with '--loop'; doesn't work with '-f'
  -f | --follow-redirect    - Sends another request if the response contains a Location header and a 3xx status code; doesn't work with '--loop'
  -l | --include-length     - Includes the content length

Examples
  kulana https://ohano.me               - To get the HTTP status and the response time of https://ohano.me
  kulana https://ohano.me --loop        - Same as above, but the request will be sent every second until the program will be stopped
  kulana https://ohano.me --loop -f     - Will result in an error message since you can't follow redirects in a loop (yet)
```

### Output format

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

#### JSON

The JSON format contains all keys to every value, which may be more useful to work with in other systems.

Sample output:
```json
{
  "url":"https://ohano.me",
  "status":"200",
  "time":"91.022730"
}
```

#### CSV

The CSV format does not contain any keys, nly the "raw" values.

The order is the following:
```csv
URL,HTTP status,time(,redirect URL)(,content length)
```

Sample output:
```csv
https://ohano.me,200,96.359968
```

### Looping

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

### Redirect URLS

If the status code of the response is between 300 and 399 and another URL is given by the `Location` header, the destination URL will be added to the output as `"destination"` (in CSV and default format the destination URL will be the 4th value).

You can send another request by using `--follow-redirect` (or `-f`).

> Kulana can't follow redirects in a loop (yet), so using `-f` and `--loop` will result in an error.

### Content length

You can include the length of the responses content by passing `--include-length` (or `-l`). The content length will be added to the output without any unit, just as plain byte value.

## 📜 License

This project is licensed under the GNU GPL v3, for more details see [license](./LICENSE).