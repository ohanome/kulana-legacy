---
title: 🔥 Usage
---

[Homepage](../index.md) > Usage

## General usage

The overall usage is following the schema
```shell
$ kulana COMMAND [...ARGS]
```

> ⚠️ **Attention**: Every prompt shown in this documentation assumes you have installed kulana globally. If that's not the case you may replace `kulana` by `./kulana` in the commands.

Kulana is separated into the following commands (usage docs are linked):

| Command             | Description                                                                       |
|---------------------|-----------------------------------------------------------------------------------|
| [Status](status.md) | Fetches the HTTP status and corresponding information for the given URL           |
| [Ping](ping.md)     | Pings the given host to see if its alive or if a specified port is open           |
| [MX Lookup](mx.md)  | Fetches all MX records for the given domain                                       |
| [Config](config.md) | Manipulates the global configuration which can contain things like mail templates |

## Arguments

All arguments are globally available, but not every command makes use of given arguments.

- [Misc](#misc)
  - [Help](#help)
- [Formatting the output](#formatting-the-output)
  - [Default](#json)
  - [JSON](#json)
  - [CSV](#csv)
- [Controlling the behavior](#controlling-the-behavior)

### Misc

#### Help

Argument:
```shell
--help
```
Alternative(s):
```shell
-h
```

**Description**

Displays the help section of the selected command.

**Example(s)**

Display the usage/help section for the ping command
```shell
$ kulana ping --help
```

**Support**

| Status | Ping | MX Lookup | Config |
|--------|------|-----------|--------|
| ✅      | ✅    | ✅         | ✅      |

### Formatting the output

#### Default

Argument: *none*

**Description**

This format is the default format that is being used if no other format is specified.

The format works like [CSV](#csv) but instead of joining all values with commas tabs (`\t`) are being used.

**Example(s)**

Formats the status output with the default formatter.
```shell
$ kulana status https://ohano.me
```

This produces output like the following:
```text
https://ohano.me        200     41.421521
```

**Support**

| Status | Ping | MX Lookup | Config |
|--------|------|-----------|--------|
| ✅      | ✅    | ✅         |        |

#### JSON

Argument:
```shell
--json
```

**Description**

Formats the output as JSON.

The JSON format is the only format where you get keys to the corresponding values.

> Should not be used when using `--loop` since this will produce invalid JSON like the following:
> ```text
> {"url":"https://ohano.me","status":200,"time":42.9144}
> {"url":"https://ohano.me","status":200,"time":41.0313}
> {"url":"https://ohano.me","status":200,"time":42.3749}
> ```

**Example(s)**

Formats the status output as JSON.
```shell
$ kulana status https://ohano.me --json
```

This produces output like the following:
```text
{"url":"https://ohano.me","status":200,"time":42.9144}
```

**Support**

| Status | Ping | MX Lookup | Config |
|--------|------|-----------|--------|
| ✅      | ✅    | ✅         |        |

#### CSV

Argument:
```shell
--csv
```

**Description**

Formats the output as CSV.

**Example(s)**

Formats the status output as JSON.
```shell
$ kulana status https://ohano.me --csv
```

Sample output:
```text
https://ohano.me,200,31.892892
```

**Support**

| Status | Ping | MX Lookup | Config |
|--------|------|-----------|--------|
| ✅      | ✅    | ✅         |        |

### Controlling the behavior

#### Looping

#### Delay