/* ====================================================
#   Copyright (C) 2021  All rights reserved.
#
#   Author        : wander
#   Email         : wander@ffactory.org
#   File Name     : run.go
#   Last Modified : 2021-07-28 10:54
#   Describe      :
#
# ====================================================*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wandercn/hotbuild/version"
	"github.com/wandercn/hotbuild/watch"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start app by hotbuid",
	Run: func(c *cobra.Command, args []string) {
		run()
	},
}

func run() {
	fmt.Println(watch.Green, `
 _   _  ___ _____   ____  _   _ ___ _     ____
| | | |/ _ \_   _| | __ )| | | |_ _| |   |  _ \
| |_| | | | || |   |  _ \| | | || || |   | | | |
|  _  | |_| || |   | |_) | |_| || || |___| |_| |
|_| |_|\___/ |_|   |____/ \___/|___|_____|____/ `, version.Version, `built with Go`, version.GoVersion, watch.Reset)
	watch.Start()
}
