package main

import (
	"os"
	"sort"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

var (
	id                        int
	ips                       bool
	reqString, output         string
	configSection, configFile string
	edgeConfig                edgegrid.Config
)

// Constants
const (
	VERSION = "0.0.3"
	URL     = "/siteshield/v1/maps"
	padding = 3
)

// MapsAPIResp response struct
type MapsAPIResp struct {
	SiteShieldMaps []Map `json:"siteShieldMaps"`
}

// Map struct
type Map struct {
	AcknowledgeRequiredBy int64    `json:"acknowledgeRequiredBy"`
	Acknowledged          bool     `json:"acknowledged"`
	AcknowledgedBy        string   `json:"acknowledgedBy"`
	AcknowledgedOn        int64    `json:"acknowledgedOn"`
	Contacts              []string `json:"contacts"`
	CurrentCidrs          []string `json:"currentCidrs"`
	ID                    int      `json:"id"`
	LatestTicketID        int      `json:"latestTicketId"`
	MapAlias              string   `json:"mapAlias"`
	McmMapRuleID          int      `json:"mcmMapRuleId"`
	ProposedCidrs         []string `json:"proposedCidrs"`
	RuleName              string   `json:"ruleName"`
	Service               string   `json:"service"`
	Shared                bool     `json:"shared"`
	Type                  string   `json:"type"`
}

func main() {
	_, inCLI := os.LookupEnv("AKAMAI_CLI")

	appName := "akamai-siteshield"
	if inCLI {
		appName = "akamai siteshield"
	}

	app := cli.NewApp()
	app.Name = appName
	app.HelpName = appName
	app.Version = VERSION
	app.Copyright = ""
	app.Authors = []cli.Author{
		{
			Name: "Petr Artamonov",
		},
		{
			Name: "Rafal Pieniazek",
		},
	}

	dir, _ := homedir.Dir()
	dir += string(os.PathSeparator) + ".edgerc"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "section, s",
			Value:       "default",
			Usage:       "`NAME` of section to use from credentials file",
			Destination: &configSection,
			EnvVar:      "AKAMAI_EDGERC_SECTION",
		},
		cli.StringFlag{
			Name:        "config, c",
			Value:       dir,
			Usage:       "Location of the credentials `FILE`",
			Destination: &configFile,
			EnvVar:      "AKAMAI_EDGERC",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "list-maps",
			Aliases: []string{"ls"},
			Usage:   "List SiteShield Maps",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "only-ids",
					Usage: "Show only SiteShield Maps IDs",
				},
			},
			Action: cmdlistMaps,
		},
		{
			Name:    "list-map",
			Aliases: []string{"lm"},
			Usage:   "List SiteShield Map by `ID`",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "output",
					Value:       "json",
					Usage:       "Output format. Supported ['json' and 'apache']",
					Destination: &output,
				},
				cli.BoolFlag{
					Name:  "only-cidrs",
					Usage: "Show only CIDR IP addresses",
				},
			},
			Action: cmdlistMap,
		},
		{
			Name:    "compare-cidrs",
			Aliases: []string{"cc"},
			Usage:   "Compare SiteShield Map Current CIDRs with Proposed by `ID`",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "only-diff",
					Usage: "Show only diff",
				},
			},
			Action: cmdCompareCidr,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Before = func(c *cli.Context) error {
		edgeConfig = config(configFile, configSection)
		return nil
	}

	app.Run(os.Args)
}
