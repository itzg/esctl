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

/*
Package common provides stuff for the esctl commands to use.
*/
package common

import (
	"github.com/codegangsta/cli"
)

// ArgHost is the ES connection host command-line argument
const ArgHost = "host"

// ArgPort is the ES port
const ArgPort = "port"

// CreateGlobalFlags is used to populate the global/top-level flags of a cli app
func CreateGlobalFlags() []cli.Flag {

	return []cli.Flag{
		cli.StringSliceFlag{Name: ArgHost, Value: &cli.StringSlice{},
			Usage: "Hostnames (or IPs) of one or more Elasticsearch nodes in the cluster. Default is localhost"},
		cli.StringFlag{Name: ArgPort, Value: "9200", Usage: "HTTP client port of the Elasticsearch node"},
	}
}
