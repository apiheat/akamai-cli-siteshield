package main

import (
	"os"
	"sort"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

var (
	id                                  int
	ips, colorOn                        bool
	reqString, output, version, appName string
	configSection, configFile           string
	edgeConfig                          edgegrid.Config
)

// Constants
const (
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

	appName = "akamai-siteshield"
	if inCLI {
		appName = "akamai siteshield"
	}

	app := cli.NewApp()
	app.Name = appName
	app.HelpName = appName
	app.Usage = "A CLI to interact with Akamai SiteShield"
	app.Version = version
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
		cli.BoolFlag{
			Name:        "no-color",
			Usage:       "Disable color output",
			Destination: &colorOn,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "List SiteShield objects",
			Subcommands: []cli.Command{
				{
					Name:  "maps",
					Usage: "List SiteShield Maps",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name:  "raw",
							Usage: "Show raw data of SiteShield Maps",
						},
					},
					Action: cmdlistMaps,
				},
				{
					Name:  "map",
					Usage: "List SiteShield Map by `ID`",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "output",
							Value:       "raw",
							Usage:       "Output format. Supported ['json==raw' and 'apache']",
							Destination: &output,
						},
						cli.BoolFlag{
							Name:  "only-addresses",
							Usage: "Show only Map addresses.",
						},
					},
					Action: cmdlistMap,
					Subcommands: []cli.Command{
						{
							Name:  "addresses",
							Usage: "List SiteShield Map Current and Proposed Addresses",
							Flags: []cli.Flag{
								cli.BoolFlag{
									Name:  "show-changes",
									Usage: "Show only changes",
								},
							},
							Action: cmdAddresses,
						},
					},
				},
			},
		},
		{
			Name:    "acknowledge",
			Aliases: []string{"ack"},
			Usage:   "Acknowledge SiteShield Map by `ID`",
			Action:  cmdAck,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Before = func(c *cli.Context) error {
		if c.Bool("no-color") {
			color.NoColor = true
		}

		config(configFile, configSection)

		return nil
	}

	app.Run(os.Args)
}
