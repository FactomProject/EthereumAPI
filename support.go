// Copyright 2016 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package EthereumAPI

import (
	"fmt"
)

func IntToQuantity(i int64) string {
	hex := fmt.Sprintf("%x", i)
	index := -1
	for i := range hex {
		if hex[i] == byte('0') {
			index = i
		} else {
			break
		}
	}
	if index > -1 {
		hex = hex[index:]
	}
	if len(hex) == 0 {
		hex = "0"
	}
	hex = "0x" + hex
	return hex
}
