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

// Package command is where all of the esctl commands are implemented
package command

import (
	"github.com/codegangsta/cli"
)

// CreateCommands is used to populate the app.Commands
func CreateCommands() []cli.Command {
	return []cli.Command{
		CreateNodesCommand(),
	}
}
