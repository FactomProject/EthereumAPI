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

func NetPeerCount() (*Quantity, error) {
	resp, err := Call("net_peerCount", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	answer := new(Quantity)
	err = MapToObject(resp.Result, answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}
