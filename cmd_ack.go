package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/urfave/cli"
)

func cmdAck(c *cli.Context) error {
	return ackMap(c)
}

func cmdStatus(c *cli.Context) error {
	return statusMap(c)
}

type Message struct {
	AcknowledgeRequiredBy      string `json:"acknowledgeRequiredBy"`
	Acknowledged               bool   `json:"acknowledged"`
	ID                         string `json:"id"`
	Message                    string `json:"message"`
	AcknowledgeRequiredByEpoch int64  `json:"acknowledgeRequiredByEpoch"`
}

func statusMap(c *cli.Context) error {
	id := setID(c)

	urlStr := fmt.Sprintf("%s/%s", URL, id)
	data := fetchData(urlStr, "GET")

	result, err := MapAPIRespParse(data)
	errorCheck(err)

	if result.Acknowledged {
		return nil
	}

	ackReqBy := fmt.Sprintf("%v", time.Unix(0, result.AcknowledgeRequiredBy*int64(time.Millisecond)))
	msg := fmt.Sprintf("SiteShield map should be acknowledged till %s", ackReqBy)

	errMessage := &Message{
		ID:                         id,
		Message:                    msg,
		Acknowledged:               result.Acknowledged,
		AcknowledgeRequiredBy:      ackReqBy,
		AcknowledgeRequiredByEpoch: result.AcknowledgeRequiredBy,
	}

	jsonMsg, err := json.MarshalIndent(errMessage, "", "  ")
	errorCheck(err)

	//fmt.Println(string(jsonMsg))
	return cli.NewExitError(string(jsonMsg), 2)
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
