/*
 * Copyright 2015 1&1 Internet AG, http://1und1.de . All rights reserved. Licensed under the Apache v2 License.
 */

package main

import (
	"flag"
	"fmt"
)

var sharedstorageFunctions = string2function{
	"list": handlerFunction{
		Arguments:   "",
		Description: "List all available shared storages.",
		Func:        sharedstoragesList,
	},
	"info": handlerFunction{
		Arguments:   "ID",
		Description: "Shows information about the selected shared storages.",
		Func:        sharedstoragesInfo,
	},
}

func sharedstoragesList() {
	storages, _ := api.GetSharedStorages()
	fmt.Printf("ID                               | Name\n")
	fmt.Printf("--------------------------------------------------------------------------------\n")
	for _, policy := range storages {
		fmt.Printf("%v | %v\n", policy.Id, policy.Name)
	}
	fmt.Printf("--------------------------------------------------------------------------------\n")
}

func sharedstoragesInfo() {
	id := flag.Arg(2)
	storage, _ := api.GetSharedStorage(id)
	printObject(storage)
}
