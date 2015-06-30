/*
 * Copyright 2015 1&1 Internet AG, http://1und1.de . All rights reserved. Licensed under the Apache v2 License.
 */

package main

import (
	"flag"
	"fmt"
)

const ipError = 3

var ipFunctions = string2function{
	"list": handlerFunction{
		Description: "List all available IPs.",
		Func:        ipsList,
	},
	"info": handlerFunction{
		Arguments:   "ID",
		Description: "Shows information about the selected IP.",
		Func:        ipsInfo,
	},
	"delete": handlerFunction{
		Arguments:   "ID",
		Description: "Deletes the selected IP.",
		Func:        ipsDelete,
	},
}

func ipsList() {
	ips, _ := api.GetPublicIps()
	fmt.Printf("ID                               | Name\n")
	fmt.Printf("--------------------------------------------------------------------------------\n")
	for _, ip := range ips {
		fmt.Printf("%v | %v\n", ip.Id, ip.IpAddress)
	}
	fmt.Printf("--------------------------------------------------------------------------------\n")
}

func ipsInfo() {
	id := flag.Arg(2)
	ip, err := api.GetPublicIp(id)
	printErrorAndExit(err, ipError)
	printObject(ip)
}

func ipsDelete() {
	id := flag.Arg(2)
	ip, err := api.GetPublicIp(id)
	printErrorAndExit(err, ipError)
	ip, err = ip.Delete()
	printErrorAndExit(err, ipError)
	printObject(ip)
}
