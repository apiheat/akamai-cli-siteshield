package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/urfave/cli"
)

func cmdlistMaps(c *cli.Context) error {
	return listMaps(c)
}

func cmdlistMap(c *cli.Context) error {
	return listMap(c)
}

func listMaps(c *cli.Context) error {
	data := fetchData(URL)

	result, err := MapsAPIRespParse(data)
	errorCheck(err)

	if c.Bool("only-ids") {
		printIDs(result.SiteShieldMaps)
	} else {
		fmt.Println(result.SiteShieldMaps)
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
	data := fetchData(urlStr)

	result, err := MapAPIRespParse(data)
	errorCheck(err)

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
}
