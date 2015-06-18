package main

import (
	"flag"
	"fmt"
	oao "github.com/jlusiardi/oneandone-cloudserver-api"
)

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
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	printObject(policy)
}

func firewallInfo() {
	id := flag.Arg(2)
	policy, err := api.GetFirewallPolicy(id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	printObject(policy)
}

func firewallDelete() {
	id := flag.Arg(2)
	policy, err := api.GetFirewallPolicy(id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	policy, err = policy.Delete()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	printObject(policy)
}

func firewallAddIp() {
	id := flag.Arg(2)
	ipId := flag.Arg(3)
	policy, err := api.GetFirewallPolicy(id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	policy, err = policy.AddServerIp(ipId)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	printObject(policy)
}

func firewallDeleteIp() {
	id := flag.Arg(2)
	ipId := flag.Arg(3)
	policy, err := api.GetFirewallPolicy(id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	policy, err = policy.DeleteServerIp(ipId)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	printObject(policy)
}
