/*
   Copyright (c) 2021 ffactory.org
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
