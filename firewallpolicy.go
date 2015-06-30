/*
 * Copyright 2015 1&1 Internet AG, http://1und1.de . All rights reserved. Licensed under the Apache v2 License.
 */

package main

import (
	"flag"
	"fmt"
	oao "github.com/jlusiardi/oneandone-cloudserver-api"
)

const firewallError = 2

var firewallFunctions = string2function{
	"list": handlerFunction{
		Description: "List the available firewall policies",
		Func:        firewallsList,
	},
	"create": handlerFunction{
		Arguments:   "NAME DESCRIPTION",
		Description: "Creates a new firewall policy with the given name and description. It opens port 22/TCP",
		Func:        firewallCreate,
	},
	"info": handlerFunction{
		Arguments:   "ID",
		Description: "Shows information about selected firewall policy.",
		Func:        firewallInfo,
	},
	"delete": handlerFunction{
		Arguments:   "ID",
		Description: "Deletes the selected firewall policy.",
		Func:        firewallDelete,
	},
	"addIp": handlerFunction{
		Arguments:   "ID IPID",
		Description: "Applies the selected firewall policy to the given IP.",
		Func:        firewallAddIp,
	},
	"deleteIp": handlerFunction{
		Arguments:   "ID IPID",
		Description: "Removes the selected firewall policy from the given IP.",
		Func:        firewallDeleteIp,
	},
}

func firewallsList() {
	policies, _ := api.GetFirewallPolicies()
	fmt.Printf("ID                               | Name\n")
	fmt.Printf("--------------------------------------------------------------------------------\n")
	for _, policy := range policies {
		fmt.Printf("%v | %v\n", policy.Id, policy.Name)
	}
	fmt.Printf("--------------------------------------------------------------------------------\n")
}

func firewallCreate() {
	name := flag.Arg(2)
	desc := flag.Arg(3)
	req := oao.FirewallPolicyCreateData{
		Name:        name,
		Description: desc,
		Rules: []oao.FirewallPolicyRulesCreateData{
			oao.FirewallPolicyRulesCreateData{
				Protocol: "TCP",
				PortFrom: oao.Int2Pointer(22),
				PortTo:   oao.Int2Pointer(22),
				SourceIp: "0.0.0.0",
			},
		},
	}
	policy, err := api.CreateFirewallPolicy(req)
	printErrorAndExit(err, firewallError)
	printObject(policy)
}

func firewallInfo() {
	id := flag.Arg(2)
	policy, err := api.GetFirewallPolicy(id)
	printErrorAndExit(err, firewallError)
	printObject(policy)
}

func firewallDelete() {
	id := flag.Arg(2)
	policy, err := api.GetFirewallPolicy(id)
	printErrorAndExit(err, firewallError)
	policy, err = policy.Delete()
	printErrorAndExit(err, firewallError)
	printObject(policy)
}

func firewallAddIp() {
	id := flag.Arg(2)
	ipId := flag.Arg(3)
	policy, err := api.GetFirewallPolicy(id)
	printErrorAndExit(err, firewallError)
	policy, err = policy.AddServerIp(ipId)
	printErrorAndExit(err, firewallError)
	printObject(policy)
}

func firewallDeleteIp() {
	id := flag.Arg(2)
	ipId := flag.Arg(3)
	policy, err := api.GetFirewallPolicy(id)
	printErrorAndExit(err, firewallError)
	policy, err = policy.DeleteServerIp(ipId)
	printErrorAndExit(err, firewallError)
	printObject(policy)
}
