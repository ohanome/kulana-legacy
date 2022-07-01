![](./docs/static/img/logo.svg)

Kūlana means "status" in hawaiian[*](https://hilo.hawaii.edu/wehe/?q=kulana#w2w2-10743) and that's all this project is about. With this tool on your hand you can fetch the status of nearly every single website, including information about the HTTP status, the response time and the content length.

Possible use cases may be to monitor your hosts or to get if there are any redirections (HTTP status 3xx) or just flex on how small your website is.

> A more in-depth guide to build, install and use can be found [here](https://ohanome.github.io/kulana).

_kūlana is brought to you by [ohano](https://ohano.me) ([GitHub](https://github.com/ohanome))._

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

## 📜 License

This project is licensed under the GNU GPL v3, for more details see [license](./LICENSE).