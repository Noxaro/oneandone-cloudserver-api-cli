package main

import (
	"flag"
	"fmt"
)

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
	ip, _ := api.GetPublicIp(id)
	printObject(ip)
}

func ipsDelete() {
	id := flag.Arg(2)
	ip, _ := api.GetPublicIp(id)
	ip, _ = ip.Delete()
	printObject(ip)
}
