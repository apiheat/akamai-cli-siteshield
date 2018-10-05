# Akamai CLI for SiteShield

The Akamai SiteShield Kit is a set of go libraries that wraps Akamai's {OPEN} APIs to help simplify common SiteShield related tasks.

Should you miss something we *gladly accept patches* :)

CLI uses custom [Akamai API client](https://github.com/apiheat/go-edgegrid)

## Configuration & Installation

### Credentials

Set up your credential files as described in the [authorization](https://developer.akamai.com/introduction/Prov_Creds.html) and [credentials](https://developer.akamai.com/introduction/Conf_Client.html) sections of the getting started guide on developer.akamai.com.

Tools expect proper format of sections in edgerc file which example is shown below

*NOTE:* Default file location is *~/.edgerc*

```
[default]
client_secret = XXXXXXXXXXXX
host = XXXXXXXXXXXX
access_token = XXXXXXXXXXXX
client_token = XXXXXXXXXXXX
```

In order to change section which is being actively used you can

* change it via `--config parameter` of the tool itself
* change it via env variable `export AKAMAI_EDGERC_CONFIG=/Users/jsmitsh/.edgerc`

In order to change section which is being actively used you can

* change it via `--section parameter` of the tool itself
* change it via env variable `export AKAMAI_EDGERC_SECTION=mycustomsection`

> *NOTE:* Make sure your API client do have appropriate scopes enabled

### Installation

The tool can be used as a stand-alone binary or in conjuction with [Akamai CLI](https://developer.akamai.com/cli).

#### Akamai-cli ( recommended )

Execute the following from console

```shell
> akamai install https://github.com/apiheat/akamai-cli-reporting
```

#### Stand-alone

As part of automated releases/builds you can download latest version from the project release page

## Usage

```shell
NAME:
   akamai-cli-siteshield - A CLI to interact with Akamai SiteShield

USAGE:
   akamai-cli-siteshield [global options] command [command options] [arguments...]

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
   --config FILE, -c FILE   Location of the credentials FILE (default: "/Users/USER_NAME/.edgerc") [$AKAMAI_EDGERC_CONFIG]
   --debug value            Debug Level [$AKAMAI_EDGERC_DEBUGLEVEL]
   --section NAME, -s NAME  NAME of section to use from credentials file (default: "default") [$AKAMAI_EDGERC_SECTION]
   --help, -h               show help
   --version, -v            print the version
```

## Development

In order to develop the tool with us do the following:

1. Fork repository
1. Clone it to your folder ( within *GO* path )
1. Ensure you can restore dependencies by running

   ```shell
   dep ensure
   ```

1. Make necessary changes
1. Make sure solution builds properly ( feel free to add tests )

   ```shell
   go build -ldflags="-s -w -X main.appVer=1.2.3 -X main.appName=$(basename `pwd`)"
   ```
