package EthereumAPI_test

import (
	"bytes"
	. "github.com/FactomProject/EthereumAPI"
	"testing"
)

func TestIntToQuantity(t *testing.T) {
	var i int64 = 0
	hex := IntToQuantity(i)
	if hex != "0x0" {
		t.Errorf("Wrong Quantity for %v - %v", i, hex)
	}

	i = 1
	hex = IntToQuantity(i)
	if hex != "0x1" {
		t.Errorf("Wrong Quantity for %v - %v", i, hex)
	}

	i = 65
	hex = IntToQuantity(i)
	if hex != "0x41" {
		t.Errorf("Wrong Quantity for %v - %v", i, hex)
	}

	i = 1024
	hex = IntToQuantity(i)
	if hex != "0x400" {
		t.Errorf("Wrong Quantity for %v - %v", i, hex)
	}
}

func TestQuantityToInt(t *testing.T) {
	var i int64 = 0
	i2, err := QuantityToInt(IntToQuantity(i))
	if i != i2 {
		t.Errorf("Wrong Quantity for %v - %v", i, i2)
	}

	i = 1
	i2, err = QuantityToInt(IntToQuantity(i))
	if err != nil {
		t.Errorf("Err - %v", err)
	}
	if i != i2 {
		t.Errorf("Wrong Quantity for %v - %v", i, i2)
	}

	i = 65
	i2, err = QuantityToInt(IntToQuantity(i))
	if err != nil {
		t.Errorf("Err - %v", err)
	}
	if i != i2 {
		t.Errorf("Wrong Quantity for %v - %v", i, i2)
	}

	i = 1024
	i2, err = QuantityToInt(IntToQuantity(i))
	if err != nil {
		t.Errorf("Err - %v", err)
	}
	if i != i2 {
		t.Errorf("Wrong Quantity for %v - %v", i, i2)
	}
}

func TestHexToData(t *testing.T) {
	var i []byte = []byte("A")
	hex := HexToData(i)
	if hex != "0x41" {
		t.Errorf("Wrong Data for %x - %v", i, hex)
	}

	i = []byte{0x00, 0x42, 0x00}
	hex = HexToData(i)
	if hex != "0x004200" {
		t.Errorf("Wrong Data for %x - %v", i, hex)
	}

	i = []byte("")
	hex = HexToData(i)
	if hex != "0x" {
		t.Errorf("Wrong Data for %x - %v", i, hex)
	}
}

func TestDataToHex(t *testing.T) {
	var i []byte = []byte("ABBA")
	hex, err := DataToHex(HexToData(i))
	if err != nil {
		t.Errorf("Err - %v", err)
	}
	if bytes.Compare(i, hex) != 0 {
		t.Errorf("%x != %x", i, hex)
	}

	i = []byte("")
	hex, err = DataToHex(HexToData(i))
	if err != nil {
		t.Errorf("Err - %v", err)
	}
	if bytes.Compare(i, hex) != 0 {
		t.Errorf("%x != %x", i, hex)
	}

	i = []byte("A")
	hex, err = DataToHex(HexToData(i))
	if err != nil {
		t.Errorf("Err - %v", err)
	}
	if bytes.Compare(i, hex) != 0 {
		t.Errorf("%x != %x", i, hex)
	}

	i = []byte{0x00, 0x42, 0x00}
	hex, err = DataToHex(HexToData(i))
	if err != nil {
		t.Errorf("Err - %v", err)
	}
	if bytes.Compare(i, hex) != 0 {
		t.Errorf("%x != %x", i, hex)
	}
}
