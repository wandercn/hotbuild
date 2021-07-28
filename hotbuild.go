/* ====================================================
#   Copyright (C) 2021  All rights reserved.
#
#   Author        : wander
#   Email         : wander@email.cn
#   File Name     : hotbuild.go
#   Last Modified : 2021-07-27 18:03
#   Describe      :
#
# ====================================================*/
package main

import (
	"log"
	"os"

	"github.com/wandercn/hotbuild/cmd"
)

func main() {
	if err := cmd.RootCmd().Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
