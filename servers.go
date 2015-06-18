package main

import (
	"flag"
	"fmt"
	oao "github.com/jlusiardi/oneandone-cloudserver-api"
)

func listServers() {
	servers, _ := api.GetServers()
	fmt.Printf("ID                               | Name\n")
	fmt.Printf("--------------------------------------------------------------------------------\n")
	for _, server := range servers {
		fmt.Printf("%v | %v\n", server.Id, server.Name)
	}
	fmt.Printf("--------------------------------------------------------------------------------\n")
}

func listInstanceSizes() {
	sizes, _ := api.GetFixedInstanceSizes()
	fmt.Printf("ID                               | Name\n")
	fmt.Printf("--------------------------------------------------------------------------------\n")
	for _, size := range sizes {
		fmt.Printf("%v | %v\n", size.Id, size.Name)
	}
	fmt.Printf("--------------------------------------------------------------------------------\n")
}
func listServersIps() {
	id := flag.Arg(2)
	server, err := api.GetServer(id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("ID                               | IP\n")
	fmt.Printf("--------------------------------------------------------------------------------\n")
	for _, ip := range server.Ips {
		fmt.Printf("%v | %v\n", ip.Id, ip.Ip)
	}
	fmt.Printf("--------------------------------------------------------------------------------\n")
}

func infoServer() {
	id := flag.Arg(2)
	server, err := api.GetServer(id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	printObject(server)
}

func deleteServer() {
	id := flag.Arg(2)
	server, err := api.GetServer(id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	server, err = server.Delete()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	printObject(server)
}

func rebootServer() {
	id := flag.Arg(2)
	server, err := api.GetServer(id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	mode := flag.Arg(3)
	server, err = server.Reboot(mode == "force")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	printObject(server)
}

func shutdownServer() {
	id := flag.Arg(2)
	server, err := api.GetServer(id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	mode := flag.Arg(3)
	server, err = server.Shutdown(mode == "force")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	api.WaitForServerState(id, "POWERED_OFF")
	printObject(server)
}

func startServer() {
	id := flag.Arg(2)
	server, err := api.GetServer(id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	server, err = server.Start()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	api.WaitForServerState(id, "POWERED_ON")
	printObject(server)
}

func createServer() {
	name := flag.Arg(2)
	appliance := flag.Arg(3)
	req := oao.ServerCreateData{
		Name:        name,
		ApplianceId: appliance,
		Hardware: oao.Hardware{
			Vcores:            1,
			CoresPerProcessor: 1,
			Ram:               1,
			Hdds: []oao.Hdd{
				oao.Hdd{
					Size:   40,
					IsMain: true,
				},
			},
		},
	}
	server, err := api.CreateServer(req)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	printObject(server)

}
