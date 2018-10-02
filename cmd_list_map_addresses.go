package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	common "github.com/apiheat/akamai-cli-common"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func cmdAddresses(c *cli.Context) error {
	return addresses(c)
}

func addresses(c *cli.Context) error {
	id := common.SetIntID(c, "Please provide Map ID")

	data, _, err := apiClient.SiteShield.ListMap(id)
	common.ErrorCheck(err)

	current := data.CurrentCidrs
	sort.Strings(current)

	proposed := data.ProposedCidrs
	sort.Strings(proposed)

	if c.Bool("show-changes") {
		if len(proposed) == 0 {
			log.Println("There are no proposed CIDRs, your current CIDR list is up to date")
			return nil
		}
		if len(common.StringsSlicesDifference(current, proposed)) > 0 {
			fmt.Println("Removed:")
			for i := range common.StringsSlicesDifference(current, proposed) {
				color.Set(color.FgRed)
				fmt.Printf("\t-%+v\n", common.StringsSlicesDifference(current, proposed)[i])
				color.Unset()
			}
		}
		if len(common.StringsSlicesDifference(proposed, current)) > 0 {
			fmt.Println("Added:")
			for i := range common.StringsSlicesDifference(proposed, current) {
				color.Set(color.FgGreen)
				fmt.Printf("\t+%+v\n", common.StringsSlicesDifference(proposed, current)[i])
				color.Unset()
			}
		}
	} else {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
		fmt.Fprintln(w, fmt.Sprint("Current\tProposed"))

		iter := proposed
		if len(current) >= len(proposed) {
			iter = current
		}

		for i := range iter {
			cIP := ""
			if i < len(current) {
				cIP = current[i]
			}

			pIP := ""
			if i < len(proposed) {
				pIP = proposed[i]
			}

			fmt.Fprintln(w, fmt.Sprintf("%s\t%s", cIP, pIP))
		}
		w.Flush()
	}

	return nil
}
