package main

import (
	"encoding/json"
	"fmt"
	"time"

	common "github.com/apiheat/akamai-cli-common"
	edgegrid "github.com/apiheat/go-edgegrid"
	log "github.com/sirupsen/logrus"

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

	ack, _, err := apiClient.SiteShield.ListMap(id)
	common.ErrorCheck(err)

	if ack.Acknowledged {
		log.Info("SiteShield Map '" + id + "' is up to date")
		return nil
	}

	ackReqBy := fmt.Sprintf("%v", time.Unix(0, ack.AcknowledgeRequiredBy*int64(time.Millisecond)))
	msg := fmt.Sprintf("SiteShield map should be acknowledged till %s", ackReqBy)

	errMessage := &Message{
		ID:                         id,
		Message:                    msg,
		Acknowledged:               ack.Acknowledged,
		AcknowledgeRequiredBy:      ackReqBy,
		AcknowledgeRequiredByEpoch: ack.AcknowledgeRequiredBy,
	}

	jsonMsg, err := json.MarshalIndent(errMessage, "", "  ")
	common.ErrorCheck(err)

	return cli.NewExitError(string(jsonMsg), 2)
}

func ackMap(c *cli.Context) error {
	id := common.SetIntID(c, "Please provide Map ID")

	// Here we need to add check if we need to ack
	ack, _, err := apiClient.SiteShield.ListMap(id)
	common.ErrorCheck(err)

	if ack.Acknowledged {
		log.Info("SiteShield Map '" + id + "' is up to date")
		return nil
	}

	data, _, err := apiClient.SiteShield.AckMap(id)
	common.ErrorCheck(err)

	var arr []edgegrid.AkamaiSiteShieldMap
	arr = append(arr, *data)
	printIDs(arr)

	return nil
}
