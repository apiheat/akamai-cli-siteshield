package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/urfave/cli"
)

func cmdlistMaps(c *cli.Context) error {
	return listMaps(c)
}

func cmdlistMap(c *cli.Context) error {
	return listMap(c)
}

func cmdCompareCidr(c *cli.Context) error {
	return compareCidr(c)
}
func cmdAck(c *cli.Context) error {
	return ackMap(c)
}

func listMaps(c *cli.Context) error {
	data := fetchData(URL, "GET")

	result, err := MapsAPIRespParse(data)
	errorCheck(err)

	if c.Bool("only-ids") {
		printIDs(result.SiteShieldMaps)
	} else {
		jsonRes, _ := json.MarshalIndent(result.SiteShieldMaps, "", "  ")
		fmt.Printf("%+v\n", string(jsonRes))
	}

	return nil
}

func listMap(c *cli.Context) error {
	var id string
	if c.NArg() > 0 {
		id = c.Args().Get(0)
		verifyID(id)
	} else {
		log.Fatal("Please provide ID for map")
	}

	urlStr := fmt.Sprintf("%s/%s", URL, id)
	data := fetchData(urlStr, "GET")

	result, err := MapAPIRespParse(data)
	errorCheck(err)

	if c.Bool("only-cidrs") {
		sorted := result.CurrentCidrs
		sort.Strings(sorted)

		switch output {
		case "apache":
			join := strings.Join(sorted[:], " ")
			outputStr := fmt.Sprintf("# Akamai SiteShield\nRequire ip %s", join)
			fmt.Println(outputStr)
		case "json":
			jsonRes, _ := json.MarshalIndent(sorted, "", "  ")
			fmt.Printf("%+v\n", string(jsonRes))
		}
	} else {
		jsonRes, _ := json.MarshalIndent(result, "", "  ")
		fmt.Printf("%+v\n", string(jsonRes))
	}

	return nil
}

func ackMap(c *cli.Context) error {
	var id string
	if c.NArg() > 0 {
		id = c.Args().Get(0)
		verifyID(id)
	} else {
		log.Fatal("Please provide ID for map")
	}

	urlStr := fmt.Sprintf("%s/%s/acknowledge", URL, id)
	data := fetchData(urlStr, "POST")

	result, err := MapAPIRespParse(data)
	errorCheck(err)

	var arr []Map
	arr = append(arr, result)
	printIDs(arr)

	return nil
}

func compareCidr(c *cli.Context) error {
	var id string
	if c.NArg() > 0 {
		id = c.Args().Get(0)
		verifyID(id)
	} else {
		log.Fatal("Please provide ID for map")
	}

	urlStr := fmt.Sprintf("%s/%s", URL, id)
	data := fetchData(urlStr, "GET")

	result, err := MapAPIRespParse(data)
	errorCheck(err)

	current := result.CurrentCidrs
	sort.Strings(current)

	proposed := result.ProposedCidrs
	sort.Strings(proposed)

	if c.Bool("only-diff") {
		if len(proposed) == 0 {
			log.Println("There are no proposed CIDRs, your current CIDR list is up to date")
			return nil
		}
		if len(difference(current, proposed)) > 0 {
			fmt.Println("Removed:")
			for i := range difference(current, proposed) {
				fmt.Printf("\t%+v\n", difference(current, proposed)[i])
			}
		}
		if len(difference(proposed, current)) > 0 {
			fmt.Println("Added:")
			for i := range difference(proposed, current) {
				fmt.Printf("\t%+v\n", difference(proposed, current)[i])
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
