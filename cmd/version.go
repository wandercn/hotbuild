/* ====================================================
#   Copyright (C) 2021  All rights reserved.
#
#   Author        : wander
#   Email         : wander@ffactory.org
#   File Name     : version.go
#   Last Modified : 2021-07-28 10:54
#   Describe      :
#
# ====================================================*/

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wandercn/hotbuild/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hotbuild",
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(version.Version)
	},
}
