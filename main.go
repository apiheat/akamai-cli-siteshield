package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
	"strconv"
	"strings"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"github.com/urfave/cli"
)

var (
	id                        int
	ips                       bool
	reqString, output         string
	configSection, configFile string
)

// Constants
const (
	VERSION = "0.0.1"
	URL     = "/siteshield/v1/maps"
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

// MapsAPIRespParse unmarshal json
func MapsAPIRespParse(in string) (maps MapsAPIResp, err error) {
	if err = json.Unmarshal([]byte(in), &maps); err != nil {
		return
	}
	return
}

// MapAPIRespParse unmarshal json
func MapAPIRespParse(in string) (maps Map, err error) {
	if err = json.Unmarshal([]byte(in), &maps); err != nil {
		return
	}
	return
}

func userHome() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func fetchData(configFile, configSection, urlPath string) (result string) {
	if configFile == ".edgerc" {
		configFile = path.Join(userHome(), ".edgerc")
	}

	config, err := edgegrid.Init(configFile, configSection)
	if err != nil {
		log.Fatal(err)
	}

	req, err := client.NewRequest(config, "GET", urlPath, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(config, req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	byt, _ := ioutil.ReadAll(resp.Body)

	return string(byt)
}

func main() {
	app := cli.NewApp()
	app.Name = "akamai-netstorage"
	app.Usage = "Akamai CLI"
	app.Version = VERSION
	app.Copyright = ""

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "section, s",
			Value:       "default",
			Usage:       "`NAME` of section to use from .edgerc",
			Destination: &configSection,
		},
		cli.StringFlag{
			Name:        "config, c",
			Value:       ".edgerc",
			Usage:       "Load configuration from `FILE`",
			Destination: &configFile,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "list-maps",
			Aliases: []string{"ls"},
			Usage:   "List SiteShield Maps `ID`",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "only-ids",
					Usage: "Show only IDs",
				},
			},
			Action: func(c *cli.Context) error {
				data := fetchData(configFile, configSection, URL)

				result, err := MapsAPIRespParse(data)
				if err != nil {
					log.Fatal(err)
				}

				if c.Bool("only-ids") {
					for _, f := range result.SiteShieldMaps {
						fmt.Println("SiteShield Map:")
						fmt.Printf("  ID: %v\n", f.ID)
						fmt.Printf("  Name: %s\n", f.RuleName)
						fmt.Printf("  Alias: %s\n", f.MapAlias)
						fmt.Printf("  Env: %s\n", f.Type)
					}
				} else {
					fmt.Println(result.SiteShieldMaps)
				}

				return nil
			},
		},
		{
			Name:    "list-map",
			Aliases: []string{"lm"},
			Usage:   "List SiteShield Map by `ID`",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "output",
					Value:       "json",
					Usage:       "Output format. supported json and apache]",
					Destination: &output,
				},
				cli.BoolFlag{
					Name:  "only-cidrs",
					Usage: "Show only CIDR IP addresses",
				},
			},
			Action: func(c *cli.Context) error {
				var id string
				if c.NArg() > 0 {
					id = c.Args().Get(0)
					if _, err := strconv.Atoi(id); err != nil {
						errStr := fmt.Sprintf("SiteShield Map ID should be number, you provided: %q\n", id)
						log.Fatal(errStr)
					}
				} else {
					log.Fatal("Please provide ID for map to fetch")
				}

				urlStr := fmt.Sprintf("%s/%s", URL, id)
				data := fetchData(configFile, configSection, urlStr)

				result, err := MapAPIRespParse(data)
				if err != nil {
					log.Fatal(err)
				}

				if c.Bool("only-cidrs") {
					switch output {
					case "apache":
						join := strings.Join(result.CurrentCidrs[:], " ")
						outputStr := fmt.Sprintf("# Akamai SiteShield\nRequire ip %s", join)
						fmt.Println(outputStr)
					case "json":
						fmt.Println(result.CurrentCidrs)
					}
				} else {
					fmt.Println(result)
				}

				return nil
			},
		},
	}

	app.Run(os.Args)
}
