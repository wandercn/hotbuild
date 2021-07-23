/* ====================================================
#   Copyright (C) 2021  All rights reserved.
#
#   Author        : wander
#   Email         : wander@email.cn
#   File Name     : main.go
#   Last Modified : 2021-07-23 12:10
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
|_| |_|\___/ |_|   |____/ \___/|___|_____|____/ `, Version, `built with Go`, goVersion, watch.Reset)
	watch.Start()
}
