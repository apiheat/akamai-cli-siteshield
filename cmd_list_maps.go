package main

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli"
)

func cmdlistMaps(c *cli.Context) error {
	return listMaps(c)
}

func listMaps(c *cli.Context) error {
	data := fetchData(URL, "GET")

	result, err := MapsAPIRespParse(data)
	errorCheck(err)

	if c.Bool("raw") {
		jsonRes, _ := json.MarshalIndent(result.SiteShieldMaps, "", "  ")
		fmt.Printf("%+v\n", string(jsonRes))
	} else {
		printIDs(result.SiteShieldMaps)
	}

	return nil
}
