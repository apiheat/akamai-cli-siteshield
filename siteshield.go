package main

  import (
    "fmt"
		"io/ioutil"
		"encoding/json"

    "github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
    "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
  )

  type MapsApiResp struct {
    SiteShieldMaps []Map `json:"siteShieldMaps"`
	}

	type Map struct {
		AcknowledgeRequiredBy int64 `json:"acknowledgeRequiredBy"`
		Contacts []string `json:"contacts"`
		CurrentCidrs []string `json:"currentCidrs"`
		ID int `json:"id"`
		Type string `json:"type"`
		RuleName string `json:"ruleName"`
	}

	func MapsApiRespParse(in string) (maps MapsApiResp, err error) {
		if err = json.Unmarshal([]byte(in), &maps); err != nil {
			return
		}
		return
	}

  func main() {
		config, _ := edgegrid.Init("~/.edgerc", "default")

    req, _ := client.NewRequest(config, "GET", "/siteshield/v1/maps", nil)
    resp, _ := client.Do(config, req)

    defer resp.Body.Close()
    byt, _ := ioutil.ReadAll(resp.Body)

    result, err := MapsApiRespParse(string(byt))
		if err != nil {
			fmt.Println("error:", err)
		}

		id := result.SiteShieldMaps[0].CurrentCidrs
		fmt.Println(id)

		// m := fmt.Sprintf("/siteshield/v1/maps/%d", t)
		// fmt.Println(m)
		// req, _ = client.NewRequest(config, "GET", m, nil)
    // resp, _ = client.Do(config, req)

    // defer resp.Body.Close()
    // byt, _ = ioutil.ReadAll(resp.Body)

    // fmt.Println(string(byt))

	}
