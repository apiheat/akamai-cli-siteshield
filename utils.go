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
	"github.com/urfave/cli"
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

func setID(c *cli.Context) string {
	var id string
	if c.NArg() == 0 {
		log.Fatal("Please provide ID for map")
	}

	id = c.Args().Get(0)
	verifyID(id)
	return id
}

func verifyID(id string) {
	if _, err := strconv.Atoi(id); err != nil {
		errStr := fmt.Sprintf("SiteShield Map ID should be number, you provided: %q\n", id)
		log.Fatal(errStr)
	}
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

func printJSON(str interface{}) {
	jsonRes, _ := json.MarshalIndent(str, "", "  ")
	fmt.Printf("%+v\n", string(jsonRes))
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

func errorCheck(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func fetchData(urlPath, method string) (result string) {
	req, err := client.NewRequest(edgeConfig, method, urlPath, nil)
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
