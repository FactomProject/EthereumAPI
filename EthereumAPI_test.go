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
