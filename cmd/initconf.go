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
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/wandercn/hotbuild/config"
)

var initconfCmd = &cobra.Command{
	Use:   "initconf",
	Short: "Initialize .hotbuild.toml config of Hotbuild",
	Run: func(c *cobra.Command, args []string) {
		if err := config.InitConf(); err != nil {
			log.Fatalf("InitConf failed: %v", err)
			os.Exit(1)
		}
	},
}
