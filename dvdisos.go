package main

import (
	"flag"
	"fmt"
)

func listDvd() {
	dvdisos, err := api.GetDvdIsos()
	fmt.Printf("%v\n", err)
	fmt.Printf("ID                               | Name\n")
	fmt.Printf("--------------------------------------------------------------------------------\n")
	for _, dvdiso := range dvdisos {
		fmt.Printf("%v | %v\n", dvdiso.Id, dvdiso.Name)
	}
	fmt.Printf("--------------------------------------------------------------------------------\n")
}

func dvdInfo() {
	id := flag.Arg(2)
	dvdiso, _ := api.GetDvdIso(id)
	printObject(dvdiso)
}
