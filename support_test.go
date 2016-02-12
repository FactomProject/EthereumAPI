package EthereumAPI_test

import (
	. "github.com/FactomProject/EthereumAPI"
	"testing"
)

func TestIntToQuantity(t *testing.T) {
	var i int64 = 0
	hex := IntToQuantity(i)
	if hex != "0x0" {
		t.Error("Wrong Quantity for %v - %v", i, hex)
	}

	i = 1
	hex = IntToQuantity(i)
	if hex != "0x1" {
		t.Error("Wrong Quantity for %v - %v", i, hex)
	}

	i = 65
	hex = IntToQuantity(i)
	if hex != "0x41" {
		t.Error("Wrong Quantity for %v - %v", i, hex)
	}

	i = 1024
	hex = IntToQuantity(i)
	if hex != "0x400" {
		t.Error("Wrong Quantity for %v - %v", i, hex)
	}
}
