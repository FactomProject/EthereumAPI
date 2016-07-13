package EthereumAPI_test

import (
	"testing"

	. "github.com/FactomProject/EthereumAPI"
)

func TestEtherscanTxList(t *testing.T) {
	resp, err := EtherscanTxList("0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae")
	t.Errorf("%v, %v", resp, err)
}
