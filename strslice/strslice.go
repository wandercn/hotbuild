/* ====================================================
#   Copyright (C) 2021  All rights reserved.
#
#   Author        : wander
#   Email         : wander@email.cn
#   File Name     : strslice.go
#   Last Modified : 2021-07-02 15:31
#   Describe      :
#
# ====================================================*/

package strslice

func IsStrInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
