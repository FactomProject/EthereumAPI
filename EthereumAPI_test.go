package EthereumAPI_test

import (
	. "github.com/FactomProject/EthereumAPI"
	"testing"
)

func TestTest(t *testing.T) {
	resp, err := Call("bla", nil)
	if err != nil {
		t.Errorf("Error - %v", err)
	}
	t.Logf("Resp - %v", resp)
	//t.Fail()
}

func TestEthProtocolVersion(t *testing.T) {
	ver, err := EthProtocolVersion()
	if err != nil {
		t.Errorf("Error - %v", err)
	}
	t.Logf("ver - %v", ver)
	//t.Fail()
}

func TestNetPeerCount(t *testing.T) {
	peers, err := NetPeerCount()
	if err != nil {
		t.Errorf("Error - %v", err)
	}
	t.Logf("number of peers - %v", peers)
	//t.Fail()
}

func TestEthGetTransactionCount(t *testing.T) {
	count, err := EthGetTransactionCount("0x7b2D985524C3fADc0F939342fc95edCaF7163616", "latest")
	if err != nil {
		t.Errorf("Error - %v", err)
	}
	t.Logf("# tx sent by 0x7b2D985524C3fADc0F939342fc95edCaF7163616 - %v", count)
}

func TestEthGetTransactionByHash(t *testing.T) {
	txinfo, err := EthGetTransactionByHash("0x1aee1d1473442e48507bdff08b62bd88baabe785798943ce1af2a0b05e787836")
	if err != nil {
		t.Errorf("Error - %v", err)
	}
	t.Logf("txinfo from 0x1aee1d1473442e48507bdff08b62bd88baabe785798943ce1af2a0b05e787836 - %v", txinfo)
}
