package main

import (
	"encoding/json"
	"flag"
	"fmt"
	oao "github.com/jlusiardi/oneandone-cloudserver-api"
	"os"
)

var (
	token string
	api   *oao.API
)

func setupArgsParsing() {
	flag.StringVar(&token, "token", "DEFAULT", "the API token")
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

type string2function map[string]handlerFunction

type handlerFunction func()

var entities2actions2functions = map[string]string2function{
	"servers": string2function{
		"list":      listServers,
		"listSizes": listInstanceSizes,
		"listIps":   listServersIps,
		"info":      infoServer,
		"delete":    deleteServer,
		"reboot":    rebootServer,
		"shutdown":  shutdownServer,
		"start":     startServer,
		"create":    createServer,
	},
	"appliances": string2function{
		"list": appliancesList,
		"info": applianceInfo,
	},
	"dvds": string2function{
		"list": listDvd,
		"info": dvdInfo,
	},
	"firewalls": string2function{
		"list":     firewallsList,
		"create":   firewallCreate,
		"info":     firewallInfo,
		"delete":   firewallDelete,
		"addIp":    firewallAddIp,
		"deleteIp": firewallDeleteIp,
	},
	"ips": string2function{
		"list":   ipsList,
		"info":   ipsInfo,
		"delete": ipsDelete,
	},
	"sharedstorages": string2function{
		"list": sharedstoragesList,
		"info": sharedstoragesInfo,
	},
}

func main() {
	setupArgsParsing()
	if err := extractToken(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(-1)
	}
	api = oao.New(token, "https://cloudpanel-api.1and1.com/v1")
	resource := flag.Arg(0)
	command := flag.Arg(1)
	if entity, contained := entities2actions2functions[resource]; contained {
		if fun, contained := entity[command]; contained {
			fun()
		} else {
			fmt.Printf("Unknown command '%v'!", command)
		}
	} else if resource == "help" {
		printHelp()
	} else {
		fmt.Printf("Unknown entity '%v'!", command)
	}
}

func printHelp() {
	for entityKey := range entities2actions2functions {
		fmt.Printf("%v\n", entityKey)
		for actionKey := range entities2actions2functions[entityKey] {
			fmt.Printf("	%v\n", actionKey)
		}
	}
}

func printObject(in interface{}) {
	bytes, _ := json.MarshalIndent(in, "", " ")
	fmt.Printf("%v\n", string(bytes))
}
