package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func cmdAck(c *cli.Context) error {
	return ackMap(c)
}

func ackMap(c *cli.Context) error {
	id := setID(c)

	// Here we need to add check if we need to ack

	urlStr := fmt.Sprintf("%s/%s/acknowledge", URL, id)
	data := fetchData(urlStr, "POST")

	result, err := MapAPIRespParse(data)
	errorCheck(err)

	var arr []Map
	arr = append(arr, result)
	printIDs(arr)

	return nil
}
