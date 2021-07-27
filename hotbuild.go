/* ====================================================
#   Copyright (C) 2021  All rights reserved.
#
#   Author        : wander
#   Email         : wander@email.cn
#   File Name     : main.go
#   Last Modified : 2021-07-27 15:54
#   Describe      :
#
# ====================================================*/
package main

import (
	"fmt"

	"github.com/wandercn/hotbuild/watch"
)

func main() {
	fmt.Println(watch.Green, `
 _   _  ___ _____   ____  _   _ ___ _     ____
| | | |/ _ \_   _| | __ )| | | |_ _| |   |  _ \
| |_| | | | || |   |  _ \| | | || || |   | | | |
|  _  | |_| || |   | |_) | |_| || || |___| |_| |
|_| |_|\___/ |_|   |____/ \___/|___|_____|____/ `, Version, `built with Go`, GoVersion, watch.Reset)
	watch.Start()
}
