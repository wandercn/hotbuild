/*
   Copyright (c) 2021-present ffactory.org
   hotbuild is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
            http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
   EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
   MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/
package cmd

import (
	"github.com/flylog/colorstyle"
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
	colorstyle.New().ColorGreen().Println(`
 _   _  ___ _____   ____  _   _ ___ _     ____
| | | |/ _ \_   _| | __ )| | | |_ _| |   |  _ \
| |_| | | | || |   |  _ \| | | || || |   | | | |
|  _  | |_| || |   | |_) | |_| || || |___| |_| |
|_| |_|\___/ |_|   |____/ \___/|___|_____|____/ `, version.Version, ` built with `, version.GoVersion)
	watch.Start()
}
