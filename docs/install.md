---
title: 📦 Installation
---

[Homepage](index.md) > Installation

## Unix/Mac

You can run the provided script `install.sh` by executing
```shell
./install.sh
```

Optionally, you can execute the steps done by the script by yourself:

1. Build the application by running `go build`
2. Copy the built binary to `/usr/bin/kulana` to make it globally available (or `~/bin/kulana` to install it only for the active user)
3. Make sure the directory where you copied the binary to is in the $PATH variable

## Windows

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