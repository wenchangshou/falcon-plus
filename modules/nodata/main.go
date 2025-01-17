// Copyright 2017 Xiaomi, Inc.
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
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wenchangshou/falcon-plus/modules/nodata/collector"
	"github.com/wenchangshou/falcon-plus/modules/nodata/config"
	"github.com/wenchangshou/falcon-plus/modules/nodata/g"
	"github.com/wenchangshou/falcon-plus/modules/nodata/http"
	"github.com/wenchangshou/falcon-plus/modules/nodata/judge"
)

func main() {
	g.BinaryName = BinaryName
	g.Version = Version
	g.GitCommit = GitCommit

	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	versionGit := flag.Bool("vg", false, "show version")
	flag.Parse()

	if *version {
		fmt.Printf("Open-Falcon %s version %s, build %s\n", BinaryName, Version, GitCommit)
		os.Exit(0)
	}
	if *versionGit {
		fmt.Printf("Open-Falcon %s version %s, build %s\n", BinaryName, Version, GitCommit)
		os.Exit(0)
	}

	// global config
	g.ParseConfig(*cfg)
	// proc
	g.StartProc()

	// config
	config.Start()
	// collector
	collector.Start()
	// judge
	judge.Start()

	// http
	http.Start()

	select {}
}
