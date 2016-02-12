// Copyright 2016 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package EthereumAPI

import (
	"encoding/hex"
	"fmt"
	"strconv"
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

func QuantityToInt(q string) (int64, error) {
	//Remove "0x" if it exists
	if len(q) > 1 {
		if q[0] == '0' && q[1] == 'x' {
			q = q[2:]
		}
	}

	//Making sure the bytes are of even length
	if len(q)%2 == 1 {
		q = "0" + q
	}

	return strconv.ParseInt(q, 16, 64)
}

func HexToData(b []byte) string {
	return fmt.Sprintf("0x%x", b)
}

func DataToHex(data string) ([]byte, error) {
	//Remove "0x" if it exists
	if len(data) > 1 {
		if data[0] == '0' && data[1] == 'x' {
			data = data[2:]
		}
	}
	return hex.DecodeString(data)
}
