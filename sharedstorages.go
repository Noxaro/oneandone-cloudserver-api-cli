package main

import (
	"flag"
	"fmt"
)

func sharedstoragesList() {
	storages, err := api.GetSharedStorages()
	fmt.Printf("%v\n", err)
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
