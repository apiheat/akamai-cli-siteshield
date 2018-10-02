package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	common "github.com/apiheat/akamai-cli-common"
	"github.com/urfave/cli"
)

func cmdlistMap(c *cli.Context) error {
	return listMap(c)
}

func listMap(c *cli.Context) error {
	id := common.SetIntID(c, "Please provide Map ID")

	data, response, err := apiClient.SiteShield.ListMap(id)
	common.ErrorCheck(err)

	sorted := data.CurrentCidrs
	sort.Strings(sorted)

	if !isOutputFormatSupported(c.String("output")) {
		log.Fatalf("Output you provided `%s` is not supported. We support 'json' and 'apache'\n", c.String("output"))
	}

	switch c.String("output") {
	case "apache":
		if c.Bool("only-addresses") {
			join := strings.Join(sorted[:], " ")
			outputStr := fmt.Sprintf("# Akamai SiteShield\nRequire ip %s", join)
			fmt.Println(outputStr)
		} else {
			log.Println("'Apache' output format can be used only to show addresses")
			log.Printf("Please run '%s list map --only-addresses --output apache %s'\n", appName, id)
		}
	case "json":
		if c.Bool("only-addresses") {
			common.OutputJSON(sorted)
		} else {
			common.PrintJSON(response.Body)
		}
	}

	return nil
}

func isOutputFormatSupported(output string) bool {
	list := []string{"json", "apache"}
	for _, b := range list {
		if b == output {
			return true
		}
	}
	return false
}
