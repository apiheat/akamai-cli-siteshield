package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func cmdAddresses(c *cli.Context) error {
	return addresses(c)
}

func addresses(c *cli.Context) error {
	id := setID(c)

	urlStr := fmt.Sprintf("%s/%s", URL, id)
	data := fetchData(urlStr, "GET")

	result, err := MapAPIRespParse(data)
	errorCheck(err)

	current := result.CurrentCidrs
	sort.Strings(current)

	proposed := result.ProposedCidrs
	sort.Strings(proposed)

	if c.Bool("show-changes") {
		if len(proposed) == 0 {
			log.Println("There are no proposed CIDRs, your current CIDR list is up to date")
			return nil
		}
		if len(difference(current, proposed)) > 0 {
			fmt.Println("Removed:")
			for i := range difference(current, proposed) {
				color.Set(color.FgRed)
				fmt.Printf("\t-%+v\n", difference(current, proposed)[i])
				color.Unset()
			}
		}
		if len(difference(proposed, current)) > 0 {
			fmt.Println("Added:")
			for i := range difference(proposed, current) {
				color.Set(color.FgGreen)
				fmt.Printf("\t+%+v\n", difference(proposed, current)[i])
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
