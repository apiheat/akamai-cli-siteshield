package main

import (
	common "github.com/apiheat/akamai-cli-common"
	"github.com/urfave/cli"
)

func cmdlistMaps(c *cli.Context) error {
	return listMaps(c)
}

func listMaps(c *cli.Context) error {
	data := fetchData(URL, "GET")

	result, err := MapsAPIRespParse(data)
	common.ErrorCheck(err)

	if c.Bool("raw") {
		common.OutputJSON(result.SiteShieldMaps)
	} else {
		printIDs(result.SiteShieldMaps)
	}

	return nil
}
