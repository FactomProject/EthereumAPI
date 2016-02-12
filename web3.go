package EthereumAPI

import (
	"fmt"
)

//https://github.com/ethereum/wiki/wiki/JSON-RPC

func Web3ClientVersion() (string, error) {
	resp, err := Call("web3_clientVersion", nil)
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

func Web3Sha3(data []byte) ([]byte, error) {
	resp, err := Call("web3_sha3", []string{HexToData(data)})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return DataToHex(resp.Result.(string))
}
