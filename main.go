package main

import (
	"os"
	"sort"

	common "github.com/apiheat/akamai-cli-common"
	edgegrid "github.com/apiheat/go-edgegrid"
	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli"
)

var (
	id              int
	ips             bool
	apiClient       *edgegrid.Client
	appName, appVer string
)

// Constants
const (
	padding = 3
)

func main() {
	app := common.CreateNewApp(appName, "A CLI to interact with Akamai SiteShield", appVer)
	app.Flags = common.CreateFlags()

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
						cli.StringFlag{
							Name:  "output",
							Value: "json",
							Usage: "Output format. Supported ['json' and 'table']",
						},
					},
					Action: cmdlistMaps,
				},
				{
					Name:  "map",
					Usage: "List SiteShield Map by `ID`",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "output",
							Value: "json",
							Usage: "Output format. Supported ['json' and 'apache']",
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
		{
			Name:   "status",
			Usage:  "Check Status: if Acknowledge for SiteShield Map by `ID` is required. If required process will exit with exit code 2",
			Action: cmdStatus,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Before = func(c *cli.Context) error {
		var err error

		apiClient, err = common.EdgeClientInit(c.GlobalString("config"), c.GlobalString("section"), c.GlobalString("debug"))

		if err != nil {
			return cli.NewExitError(err, 1)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
