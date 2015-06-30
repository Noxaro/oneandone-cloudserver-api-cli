/*
 * Copyright 2015 1&1 Internet AG, http://1und1.de . All rights reserved. Licensed under the Apache v2 License.
 */

package main

import (
	"flag"
	"fmt"
)

const applianceError = 4

var applianceFunctions = string2function{
	"list": handlerFunction{
		Arguments:   "",
		Description: "List all available appliances.",
		Func:        appliancesList,
	},
	"info": handlerFunction{
		Arguments:   "ID",
		Description: "Shows information about the selected appliance.",
		Func:        applianceInfo,
	},
}

func applianceInfo() {
	id := flag.Arg(2)
	server, err := api.GetServerAppliance(id)
	printErrorAndExit(err, applianceError)
	printObject(server)
}

func appliancesList() {
	serverappliances, _ := api.GetServerAppliances()
	fmt.Printf("ID                               | Name\n")
	fmt.Printf("--------------------------------------------------------------------------------\n")
	for _, server := range serverappliances {
		fmt.Printf("%v | %v\n", server.Id, server.Name)
	}
	fmt.Printf("--------------------------------------------------------------------------------\n")
}
