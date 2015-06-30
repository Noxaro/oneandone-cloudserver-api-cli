/*
 * Copyright 2015 1&1 Internet AG, http://1und1.de . All rights reserved. Licensed under the Apache v2 License.
 */

package main

import (
	"flag"
	"fmt"
)

const dvdError = 1

var dvdFunctions = string2function{
	"list": handlerFunction{
		Arguments:   "",
		Description: "List all available DVD isos",
		Func:        listDvd,
	},
	"info": handlerFunction{
		Arguments:   "ID",
		Description: "Shows information about the selected DVD iso.",
		Func:        dvdInfo,
	},
}

func listDvd() {
	dvdisos, _ := api.GetDvdIsos()
	fmt.Printf("ID                               | Name\n")
	fmt.Printf("--------------------------------------------------------------------------------\n")
	for _, dvdiso := range dvdisos {
		fmt.Printf("%v | %v\n", dvdiso.Id, dvdiso.Name)
	}
	fmt.Printf("--------------------------------------------------------------------------------\n")
}

func dvdInfo() {
	id := flag.Arg(2)
	dvdiso, err := api.GetDvdIso(id)
	printErrorAndExit(err, dvdError)
	printObject(dvdiso)
}
