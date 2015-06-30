/*
 * Copyright 2015 1&1 Internet AG, http://1und1.de . All rights reserved. Licensed under the Apache v2 License.
 */

package main

import (
	"flag"
	"fmt"
	oao "github.com/jlusiardi/oneandone-cloudserver-api"
)

const serversError = 5

var serverFunctions = string2function{
	"list": handlerFunction{
		Arguments:   "",
		Description: "List the available servers.",
		Func:        listServers,
	},
	"listSizes": handlerFunction{
		Arguments:   "",
		Description: "Show the available sizes for fixed instances.",
		Func:        listInstanceSizes,
	},
	"listIps": handlerFunction{
		Arguments:   "ID",
		Description: "Shows the list of all IPs of the selected server.",
		Func:        listServersIps,
	},
	"info": handlerFunction{
		Arguments:   "ID",
		Description: "Shows information about the selected server.",
		Func:        infoServer,
	},
	"delete": handlerFunction{
		Arguments:   "ID",
		Description: "Deletes the selected server.",
		Func:        deleteServer,
	},
	"reboot": handlerFunction{
		Arguments:   "ID",
		Description: "Reboots the selected server.",
		Func:        rebootServer,
	},
	"shutdown": handlerFunction{
		Arguments:   "ID",
		Description: "Shutdown the selected server.",
		Func:        shutdownServer,
	},
	"start": handlerFunction{
		Arguments:   "ID",
		Description: "Start the selected server.",
		Func:        startServer,
	},
	"create": handlerFunction{
		Arguments:   "NAME APPLIANCEID",
		Description: "Create a new server with given name and based on the given appliance.",
		Func:        createServer,
	},
	"rename": handlerFunction{
		Arguments:   "NAME DESCRIPTION",
		Description: "Update the name and description of a server.",
		Func:        renameServer,
	},
}

func renameServer() {
	id := flag.Arg(2)
	name := flag.Arg(3)
	desc := flag.Arg(4)
	server, err := api.GetServer(id)
	printErrorAndExit(err, serversError)
	data := oao.ServerRenameData{
		Name:        name,
		Description: desc,
	}
	server, err = server.RenameServer(data)
	printErrorAndExit(err, serversError)
	printObject(server)
}

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
	printErrorAndExit(err, serversError)
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
	printErrorAndExit(err, serversError)
	printObject(server)
}

func deleteServer() {
	id := flag.Arg(2)
	server, err := api.GetServer(id)
	printErrorAndExit(err, serversError)
	server, err = server.Delete()
	printErrorAndExit(err, serversError)
	printObject(server)
}

func rebootServer() {
	id := flag.Arg(2)
	server, err := api.GetServer(id)
	printErrorAndExit(err, serversError)
	mode := flag.Arg(3)
	server, err = server.Reboot(mode == "force")
	printErrorAndExit(err, serversError)
	printObject(server)
}

func shutdownServer() {
	id := flag.Arg(2)
	server, err := api.GetServer(id)
	printErrorAndExit(err, serversError)
	mode := flag.Arg(3)
	server, err = server.Shutdown(mode == "force")
	printErrorAndExit(err, serversError)
	api.WaitForServerState(id, "POWERED_OFF")
	server, err = api.GetServer(id)
	printErrorAndExit(err, serversError)
	printObject(server)
}

func startServer() {
	id := flag.Arg(2)
	server, err := api.GetServer(id)
	printErrorAndExit(err, serversError)
	server, err = server.Start()
	printErrorAndExit(err, serversError)
	api.WaitForServerState(id, "POWERED_ON")
	server, err = api.GetServer(id)
	printErrorAndExit(err, serversError)
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
	printErrorAndExit(err, serversError)
	printObject(server)

}
