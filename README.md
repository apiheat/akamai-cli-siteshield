# Akamai CLI for SiteShield
*NOTE:* This tool is intended to be installed via the Akamai CLI package manager, which can be retrieved from the releases page of the [Akamai CLI](https://github.com/akamai/cli) tool.

### Local Install, if you choose not to use the akamai package manager
If you want to compile it from source, you will need Go 1.9 or later, and the [Glide](https://glide.sh) package manager installed:
1. Fetch the package:
   `go get https://github.com/partamonov/akamai-cli-siteshield`
1. Change to the package directory:
   `cd $GOPATH/src/github.com/partamonov/akamai-cli-siteshield`
1. Install dependencies using Glide:
   `glide install`
1. Compile the binary:
   `go build -ldflags="-s -w -X main.version=X.X.X" -o akamai-siteshield`

### Credentials
In order to use this configuration, you need to:
* Set up your credential files as described in the [authorization](https://developer.akamai.com/introduction/Prov_Creds.html) and [credentials](https://developer.akamai.com/introduction/Conf_Client.html) sections of the getting started guide on developer.akamai.com.

Expects `default` section in .edgerc, can be changed via --section parameter

```
[default]
client_secret = XXXXXXXXXXXX
host = XXXXXXXXXXXX
access_token = XXXXXXXXXXXX
client_token = XXXXXXXXXXXX
```

## Overview
The Akamai SiteShield Kit is a set of go libraries that wraps Akamai's {OPEN} APIs to help simplify common siteshield tasks.

## Main Command Usage
```shell
NAME:
   akamai-siteshield - A CLI to interact with Akamai SiteShield

USAGE:
   akamai-siteshield [global options] command [command options] [arguments...]

VERSION:
   X.X.X

AUTHORS:
   Petr Artamonov
   Rafal Pieniazek

COMMANDS:
     acknowledge, ack  Acknowledge SiteShield Map by `ID`
     list, ls          List SiteShield objects
     status            Check Status: if Acknowledge for SiteShield Map by `ID` is required. If required process will exit with exit code 2
     help, h           Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config FILE, -c FILE   Location of the credentials FILE (default: "/Users/partamonov/.edgerc") [$AKAMAI_EDGERC_CONFIG]
   --debug value            Debug Level [$AKAMAI_EDGERC_DEBUGLEVEL]
   --section NAME, -s NAME  NAME of section to use from credentials file (default: "default") [$AKAMAI_EDGERC_SECTION]
   --help, -h               show help
   --version, -v            print the version
```

### Acknowledge command
```shell
NAME:
   akamai-siteshield acknowledge - Acknowledge SiteShield Map by `ID`

USAGE:
   akamai-siteshield acknowledge [arguments...]
```

### List commands
```shell
NAME:
   akamai-siteshield list - List SiteShield objects

USAGE:
   akamai-siteshield list command [command options] [arguments...]

COMMANDS:
     maps  List SiteShield Maps
     map   List SiteShield Map by `ID`

OPTIONS:
   --help, -h  show help
```

### List all Maps
```shell
NAME:
   akamai-siteshield list maps - List SiteShield Maps

USAGE:
   akamai-siteshield list maps [command options] [arguments...]

OPTIONS:
   --raw  Show raw data of SiteShield Maps
```

### List Map by ID
```shell
NAME:
   akamai-siteshield list map - List SiteShield Map by `ID`

USAGE:
   akamai-siteshield list map command [command options] [arguments...]

COMMANDS:
     addresses  List SiteShield Map Current and Proposed Addresses

OPTIONS:
   --output value    Output format. Supported ['json' and 'apache'] (default: "raw")
   --only-addresses  Show only Map addresses.
   --help, -h        show help
```

### List Map current and proposed IP adresses
```shell
NAME:
   akamai-siteshield list map addresses - List SiteShield Map Current and Proposed Addresses

USAGE:
   akamai-siteshield list map addresses [command options] [arguments...]

OPTIONS:
   --show-changes  Show only changes
```