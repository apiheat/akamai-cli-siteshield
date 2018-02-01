package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	client "github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
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

func verifyID(id string) {
	if _, err := strconv.Atoi(id); err != nil {
		errStr := fmt.Sprintf("SiteShield Map ID should be number, you provided: %q\n", id)
		log.Fatal(errStr)
	}
}

func printIDs(data []Map) {
	fmt.Println("SiteShield Maps:")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	fmt.Fprintln(w, fmt.Sprint("ID\tName\tAlias\tEnv\tAcknowledged\tAcknowledge Required By"))
	for _, f := range data {
		fmt.Fprintln(w, fmt.Sprintf("%v\t%s\t%s\t%s\t%v\t%v",
			f.ID, f.RuleName, f.MapAlias, f.Type,
			time.Unix(0, f.AcknowledgedOn*int64(time.Millisecond)),
			time.Unix(0, f.AcknowledgeRequiredBy*int64(time.Millisecond)),
		))
	}
	w.Flush()
}

func errorCheck(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func fetchData(urlPath string) (result string) {
	req, err := client.NewRequest(edgeConfig, "GET", urlPath, nil)
	errorCheck(err)

	resp, err := client.Do(edgeConfig, req)
	errorCheck(err)

	defer resp.Body.Close()
	byt, _ := ioutil.ReadAll(resp.Body)

	return string(byt)
}

func difference(slice1 []string, slice2 []string) []string {
	var diff []string
	for _, s1 := range slice1 {
		found := false
		for _, s2 := range slice2 {
			if s1 == s2 {
				found = true
				break
			}
		}
		// String not found. We add it to return slice
		if !found {
			diff = append(diff, s1)
		}
	}

	return diff
}
