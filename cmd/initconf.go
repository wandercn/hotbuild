/* ====================================================
#   Copyright (C) 2021  All rights reserved.
#
#   Author        : wander
#   Email         : wander@ffactory.org
#   File Name     : initconf.go
#   Last Modified : 2021-07-28 10:54
#   Describe      :
#
# ====================================================*/

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
