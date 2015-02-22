// Copyright 2015 Geoff Bourne
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.package main

package command

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/itzg/esctl/common"
)

// CreateNodesCommand provides a command to get the cluster name and the nodes listing of that cluster
func CreateNodesCommand() cli.Command {
	return cli.Command{
		Name:   "nodes",
		Usage:  "get basic info about the Elasticsearch node",
		Action: handleNodesCommand,
	}
}

func handleNodesCommand(c *cli.Context) {
	conn := common.CreateConnection(c)
	nodeInfo, err := conn.AllNodesInfo()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Cluster : %v\n", nodeInfo.ClusterName)
	for _, node := range nodeInfo.Nodes {
		fmt.Printf("\nNode : %v\n", node.Name)
		fmt.Printf("   %17s : %v\n", "Version", node.Version)
		fmt.Printf("   %17s : %v (%v)\n", "Host/IP", node.Host, node.Ip)
		fmt.Printf("   %17s : %v\n", "Transport Address", node.TransportAddress)
	}
}
