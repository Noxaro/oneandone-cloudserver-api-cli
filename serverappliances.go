package main

import (
	"flag"
	"fmt"
)

func applianceInfo() {
	id := flag.Arg(2)
	server, _ := api.GetServerAppliance(id)
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
