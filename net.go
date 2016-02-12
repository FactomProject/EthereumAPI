package EthereumAPI

import (
	"fmt"
)

//https://github.com/ethereum/wiki/wiki/JSON-RPC

func NetVersion() (string, error) {
	resp, err := Call("net_version", nil)
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

func NetListening() (bool, error) {
	resp, err := Call("net_listening", nil)
	if err != nil {
		return false, err
	}
	if resp.Error != nil {
		return false, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(bool), nil
}

func NetPeerCount() (int64, error) {
	resp, err := Call("net_peerCount", nil)
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return QuantityToInt(resp.Result.(string))
}
