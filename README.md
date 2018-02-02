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
   `go build -ldflags="-s -w" -o akamai-siteshield`

### Credentials
In order to use this configuration, you need to:
* Set up your credential files as described in the [authorization](https://developer.akamai.com/introduction/Prov_Creds.html) and [credentials](https://developer.akamai.com/introduction/Conf_Client.html) sections of the getting started guide on developer.akamai.com.

Expects `default` section in .edgerc, can be changed via --section parameter

```
[netstorage]
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
   0.0.3

AUTHORS:
   Petr Artamonov
   Rafal Pieniazek

COMMANDS:
     acknowledge, ack   Acknowledge SiteShield Map by `ID`
     compare-cidrs, cc  Compare SiteShield Map Current CIDRs with Proposed by `ID`
     list-map, lm       List SiteShield Map by `ID`
     list-maps, ls      List SiteShield Maps
     help, h            Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config FILE, -c FILE   Location of the credentials FILE (default: "/Users/partamonov/.edgerc") [$AKAMAI_EDGERC]
   --section NAME, -s NAME  NAME of section to use from credentials file (default: "default") [$AKAMAI_EDGERC_SECTION]
   --help, -h               show help
   --version, -v            print the version
```

### Compare CIDRs in current and proposed list
```shell
NAME:
   akamai-siteshield compare-cidrs - Compare SiteShield Map Current CIDRs with Proposed by `ID`

USAGE:
   akamai-siteshield compare-cidrs [command options] [arguments...]

OPTIONS:
   --only-diff  Show only diff
```

### List Map by ID
```shell
NAME:
   akamai-siteshield list-map - List SiteShield Map by `ID`

USAGE:
   akamai-siteshield list-map [command options] [arguments...]

OPTIONS:
   --output value  Output format. Supported ['json' and 'apache'] (default: "json")
   --only-cidrs    Show only CIDR IP addresses
```

### List all Maps
```shell
NAME:
   akamai-siteshield list-maps - List SiteShield Maps

USAGE:
   akamai-siteshield list-maps [command options] [arguments...]

OPTIONS:
   --only-ids  Show only SiteShield Maps IDs
```