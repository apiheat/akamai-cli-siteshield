# Akamai CLI for SiteShield
*NOTE:* This tool is intended to be installed via the Akamai CLI package manager, which can be retrieved from the releases page of the [Akamai CLI](https://github.com/akamai/cli) tool.

### Local Install, if you choose not to use the akamai package manager
* Go 1.9.2
* go get https://github.com/partamonov/akamai-cli-siteshield
* cd $GOPATH/src/github.com/partamonov/akamai-cli-siteshield
* go build

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

## Usage
```shell
# akamai siteshield
NAME:
   akamai siteshield - A new cli application

USAGE:
   akamai siteshield [global options] command [command options] [arguments...]

VERSION:
   0.0.2

AUTHORS:
   Petr Artamonov
   Rafal Pieniazek

COMMANDS:
     list-map, lm   List SiteShield Map by `ID`
     list-maps, ls  List SiteShield Maps `ID`
     help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config FILE, -c FILE   Location of the credentials FILE (default: "/Users/partamonov/.edgerc") [$AKAMAI_EDGERC]
   --section NAME, -s NAME  NAME of section to use from credentials file (default: "default") [$AKAMAI_EDGERC_SECTION]
   --help, -h               show help
   --version, -v            print the version
```

### List Maps
```shell
NAME:
   akamai-siteshield list-maps - List SiteShield Maps

USAGE:
   akamai-siteshield list-maps --only-ids

OPTIONS:
   --only-ids  Show only SiteShield Maps IDs
```

### List Map
```shell
NAME:
   akamai-siteshield list-map - List SiteShield Map by `ID`

USAGE:
   akamai-siteshield list-map [command options] `ID`

OPTIONS:
   --output value  Output format. Supported [json and apache] (default: "json")
   --only-cidrs    Show only CIDR IP addresses
```