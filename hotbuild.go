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
