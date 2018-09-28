package main

import (
	"encoding/json"
	"fmt"
	"time"

	common "github.com/apiheat/akamai-cli-common"
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
	id := common.SetIntID(c, "Please provide Map ID")

	urlStr := fmt.Sprintf("%s/%s", URL, id)
	data := fetchData(urlStr, "GET")

	result, err := MapAPIRespParse(data)
	common.ErrorCheck(err)

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
	common.ErrorCheck(err)

	return cli.NewExitError(string(jsonMsg), 2)
}

func ackMap(c *cli.Context) error {
	id := common.SetIntID(c, "Please provide Map ID")

	// Here we need to add check if we need to ack

	urlStr := fmt.Sprintf("%s/%s/acknowledge", URL, id)
	data := fetchData(urlStr, "POST")

	result, err := MapAPIRespParse(data)
	common.ErrorCheck(err)

	var arr []Map
	arr = append(arr, result)
	printIDs(arr)

	return nil
}
