package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"
	"time"

	client "github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	common "github.com/apiheat/akamai-cli-common"
)

// MapsAPIRespParse unmarshal json
func MapsAPIRespParse(in string) (maps MapsAPIResp, err error) {
	if err = json.Unmarshal([]byte(in), &maps); err != nil {
		return
	}
	return
}

// MapAPIRespParse unmarshal json
func MapAPIRespParse(in string) (maps Map, err error) {
	if err = json.Unmarshal([]byte(in), &maps); err != nil {
		return
	}
	return
}

func isOutputFormatSupported(output string) bool {
	list := []string{"raw", "apache"}
	for _, b := range list {
		if b == output {
			return true
		}
	}
	return false
}

func printIDs(data []Map) {
	fmt.Println("SiteShield Maps:")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	fmt.Fprintln(w, fmt.Sprint("ID\tName\tAlias\tEnv\tAcknowledged\tAcknowledge Required By"))
	for _, f := range data {
		if f.AcknowledgeRequiredBy == 0 {
			fmt.Fprintln(w, fmt.Sprintf("%v\t%s\t%s\t%s\t%v\t%s",
				f.ID, f.RuleName, f.MapAlias, f.Type,
				time.Unix(0, f.AcknowledgedOn*int64(time.Millisecond)),
				"Up to date",
			))
		} else {
			fmt.Fprintln(w, fmt.Sprintf("%v\t%s\t%s\t%s\t%v\t%v",
				f.ID, f.RuleName, f.MapAlias, f.Type,
				time.Unix(0, f.AcknowledgedOn*int64(time.Millisecond)),
				time.Unix(0, f.AcknowledgeRequiredBy*int64(time.Millisecond)),
			))
		}
	}
	w.Flush()
}

func fetchData(urlPath, method string) (result string) {
	req, err := client.NewRequest(edgeConfig, method, urlPath, nil)
	common.ErrorCheck(err)

	resp, err := client.Do(edgeConfig, req)
	common.ErrorCheck(err)

	defer resp.Body.Close()
	byt, _ := ioutil.ReadAll(resp.Body)

	return string(byt)
}
