/* ====================================================
#   Copyright (C) 2021  All rights reserved.
#
#   Author        : wander
#   Email         : wander@email.cn
#   File Name     : main.go
#   Last Modified : 2021-07-03 08:08
#   Describe      :
#
# ====================================================*/
package main

import (
	"fmt"

	"github.com/wandercn/hotbuild/watch"
)

func main() {
	fmt.Printf(`
 _   _  ___ _____   ____  _   _ ___ _     ____
| | | |/ _ \_   _| | __ )| | | |_ _| |   |  _ \
| |_| | | | || |   |  _ \| | | || || |   | | | |
|  _  | |_| || |   | |_) | |_| || || |___| |_| |
|_| |_|\___/ |_|   |____/ \___/|___|_____|____/ %s, built with Go %s

`, hotbuildVersion, goVersion)
	watch.Start()
}
