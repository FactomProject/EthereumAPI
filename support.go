// Copyright 2016 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package EthereumAPI

import (
	"encoding/hex"
	"fmt"
	"github.com/tonnerre/golang-go.crypto/sha3"
	"strconv"
)

//https://github.com/ethereum/wiki/wiki/Ethereum-Contract-ABI

func IntToQuantity(i int64) string {
	return "0x" + IntToQuantityWithoutPrefix(i)
}

func IntToQuantityWithoutPrefix(i int64) string {
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
	return "0x" + HexToDataWithoutPrefix(b)
}

func HexToDataWithoutPrefix(b []byte) string {
	return fmt.Sprintf("%x", b)
}

func HexToPaddedData(b []byte) string {
	return "0x" + HexToDataWithoutPrefix(b)
}

func HexToPaddedDataWithoutPrefix(b []byte) string {
	l := len(b)
	data := IntToData(int64(l))
	data += fmt.Sprintf("%x", b)
	if l%32 != 0 {
		rest := make([]byte, 32-l%32)
		data += fmt.Sprintf("%x", rest)
	}
	return data
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

func StringToMethodID(method string) string {
	h := sha3.NewKeccak256()
	h.Write([]byte(method))
	var digest [32]byte
	h.Sum(digest[:0])
	return fmt.Sprintf("%x", digest[:4])
}

func IntToData(i int64) string {
	return fmt.Sprintf("%064x", i)
}

func StringToData(str string) string {
	return HexToPaddedData([]byte(str))
}

func StringToDataWithoutPrefix(str string) string {
	return HexToPaddedDataWithoutPrefix([]byte(str))
}
