package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

var (
	id                int
	ips               bool
	reqString, output string
)

type MapsApiResp struct {
	SiteShieldMaps []Map `json:"siteShieldMaps"`
}

type Map struct {
	Contacts     []string `json:"contacts"`
	CurrentCidrs []string `json:"currentCidrs"`
	ID           int      `json:"id"`
	Type         string   `json:"type"`
	RuleName     string   `json:"ruleName"`
}

func MapsApiRespParse(in string) (maps MapsApiResp, err error) {
	if err = json.Unmarshal([]byte(in), &maps); err != nil {
		return
	}
	return
}

func MapApiRespParse(in string) (maps Map, err error) {
	if err = json.Unmarshal([]byte(in), &maps); err != nil {
		return
	}
	return
}

func init() {
	flag.BoolVar(&ips, "only-cidrs", false, "Show only current CIDRs")
	flag.StringVar(&output, "output", "", "Output format")
	flag.IntVar(&id, "id", 0, "Siteshield map ID")
	flag.Parse()
}

func main() {
	reqString := "/siteshield/v1/maps"
	config, _ := edgegrid.Init("~/.edgerc", "default")

	if id != 0 {
		reqString = fmt.Sprintf("/siteshield/v1/maps/%d", id)
	}

	req, _ := client.NewRequest(config, "GET", reqString, nil)
	resp, _ := client.Do(config, req)

	defer resp.Body.Close()
	byt, _ := ioutil.ReadAll(resp.Body)

	if id != 0 {
		result, err := MapApiRespParse(string(byt))
		if err != nil {
			fmt.Println("error:", err)
		}
		if ips {
			switch output {
			case "apache":
				join := strings.Join(result.CurrentCidrs[:], " ")
				output := fmt.Sprintf("# Akamai SiteShield\nRequire ip %s", join)
				fmt.Println(output)
			default:
				fmt.Println(result.CurrentCidrs)
			}
		} else {
			fmt.Println(result)
		}
	} else {
		result, err := MapsApiRespParse(string(byt))
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Println(result.SiteShieldMaps)
	}

}
