package main

import (
	"fmt"
	"log"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"github.com/fatih/color"
	"github.com/go-ini/ini"
)

func config(configFile, configSection string) edgegrid.Config {
	cfg, err := ini.Load(configFile)
	if err != nil {
		color.Set(color.FgRed)
		log.Printf("You have problem with configuration file. For help, please run '%s --help'\n", appName)
		errorMessage := fmt.Sprintf("'%s' does not exist. Please run '%s --config Your_Configuration_File'...", configFile, appName)
		log.Fatal(errorMessage)
		color.Unset()
	}

	_, err = cfg.GetSection(configSection)
	if err != nil {
		color.Set(color.FgRed)
		log.Printf("You have problem with configuration file. For help, please run '%s --help'\n", appName)
		errorMessage := fmt.Sprintf("Section '%s' does not exist in %s. Please run '%s --section Your_Section_Name' ...", configSection, configFile, appName)
		log.Fatal(errorMessage)
		color.Unset()
	}

	config, err := edgegrid.Init(configFile, configSection)
	errorCheck(err)

	return config
}
