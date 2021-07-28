/* ====================================================
#   Copyright (C) 2021  All rights reserved.
#
#   Author        : wander
#   Email         : wander@ffactory.org
#   File Name     : root.go
#   Last Modified : 2021-07-28 10:54
#   Describe      :
#
# ====================================================*/

package cmd

import "github.com/spf13/cobra"

const Cliname = "hotbuild"

func RootCmd() *cobra.Command {
	var command = &cobra.Command{
		Use:   Cliname,
		Short: "Hotbuild is a cross platform hot compilation tool",
		Long: `A cross platform hot compilation tool built by wandercn in Go.
Complete documentation is available at https://hotbuild.ffactory.org`,
		Run: func(c *cobra.Command, args []string) {
			c.HelpFunc()(c, args)
		},
	}
	command.AddCommand(runCmd)
	command.AddCommand(initconfCmd)
	command.AddCommand(versionCmd)
	command.CompletionOptions.DisableDefaultCmd = true
	return command
}
