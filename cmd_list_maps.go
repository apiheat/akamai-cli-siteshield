package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	common "github.com/apiheat/akamai-cli-common"
	edgegrid "github.com/apiheat/go-edgegrid"

	"github.com/urfave/cli"
)

func cmdlistMaps(c *cli.Context) error {
	return listMaps(c)
}

func listMaps(c *cli.Context) error {
	data, response, err := apiClient.SiteShield.ListMaps()
	common.ErrorCheck(err)

	switch c.String("output") {
	case "json":
		common.PrintJSON(response.Body)
	case "table":
		printIDs(data.SiteShieldMaps)
	}

	return nil
}

func printIDs(data []edgegrid.AkamaiSiteShieldMap) {
	fmt.Println("SiteShield Maps:")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	fmt.Fprintln(w, fmt.Sprint("ID\tName\tAlias\tEnv\tAcknowledged\tAcknowledge Required By"))
	for _, f := range data {
		if f.AcknowledgeRequiredBy == 0 {
			fmt.Fprintln(w, fmt.Sprintf("%v\t%s\t%s\t%s\t%v\t%s",
				f.ID, f.RuleName, f.MapAlias, f.Type,
				time.Unix(0, f.AcknowledgedOn*int64(time.Millisecond)),
				"Up to date",
			))
		} else {
			fmt.Fprintln(w, fmt.Sprintf("%v\t%s\t%s\t%s\t%v\t%v",
				f.ID, f.RuleName, f.MapAlias, f.Type,
				time.Unix(0, f.AcknowledgedOn*int64(time.Millisecond)),
				time.Unix(0, f.AcknowledgeRequiredBy*int64(time.Millisecond)),
			))
		}
	}
	w.Flush()
}
