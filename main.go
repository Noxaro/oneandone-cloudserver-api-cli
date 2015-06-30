/*
 * Copyright 2015 1&1 Internet AG, http://1und1.de . All rights reserved. Licensed under the Apache v2 License.
 */

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	oao "github.com/jlusiardi/oneandone-cloudserver-api"
	"os"
	"sort"
)

var (
	token    string
	endpoint string
	api      *oao.API
)

func setupArgsParsing() {
	flag.StringVar(&token, "token", "DEFAULT", "the API token")
	flag.StringVar(&endpoint, "endpoint", "DEFAULT", "the API endpoint")
}

func extractToken() error {
	flag.Parse()
	if token != "DEFAULT" {
		return nil
	}
	token = os.Getenv("TOKEN")
	if token != "" {
		return nil
	}
	return fmt.Errorf("No token found, use either command line option -token or environment variable TOKEN")
}

func extractEndpoint() error {
	flag.Parse()
	if endpoint != "DEFAULT" {
		return nil
	}
	endpoint = os.Getenv("ENDPOINT")
	if endpoint != "" {
		return nil
	}
	return fmt.Errorf("No endpoint found, use either command line option -endpoint or environment variable ENDPOINT")
}

type string2function map[string]handlerFunction

type handlerFunction struct {
	Arguments   string
	Description string
	Func        func()
}

var entities2actions2functions = map[string]string2function{
	"servers":        serverFunctions,
	"appliances":     applianceFunctions,
	"dvds":           dvdFunctions,
	"firewalls":      firewallFunctions,
	"ips":            ipFunctions,
	"sharedstorages": sharedstorageFunctions,
}

func printErrorAndExit(err error, code int) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(code)
	}
}

func main() {
	setupArgsParsing()
	if err := extractEndpoint(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(-1)
	}
	if err := extractToken(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(-1)
	}
	api = oao.New(token, endpoint)
	resource := flag.Arg(0)
	command := flag.Arg(1)
	if entity, contained := entities2actions2functions[resource]; contained {
		if fun, contained := entity[command]; contained {
			fun.Func()
		} else {
			fmt.Printf("Unknown command '%v'!\n\n", command)
			printHelp()
		}
	} else if resource == "help" {
		printHelp()
	} else {
		fmt.Printf("Unknown entity '%v'!\n\n", resource)
		printHelp()
	}
}

func sortResourceKeys(in map[string]string2function) []string {
	var keys []string
	for k := range in {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func sortActionKeys(in string2function) []string {
	var keys []string
	for k := range in {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func printHelp() {
	fmt.Printf("-endpoint or environment variable ENDPOINT: set the API endpoint that will be used\n")
	fmt.Printf("-token or environment variable TOKEN: set the API token that will be used\n")
	fmt.Printf("\n")
	for _, entityKey := range sortResourceKeys(entities2actions2functions) {
		fmt.Printf("%v\n", entityKey)
		for _, actionKey := range sortActionKeys(entities2actions2functions[entityKey]) {
			fun := entities2actions2functions[entityKey][actionKey]
			fmt.Printf("	%v %v\n", actionKey, fun.Arguments)
			fmt.Printf("		%v\n", fun.Description)
		}
	}
}

func printObject(in interface{}) {
	bytes, _ := json.MarshalIndent(in, "", " ")
	fmt.Printf("%v\n", string(bytes))
}
