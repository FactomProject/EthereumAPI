package EthereumAPI

import (
	"fmt"
)

//https://github.com/ethereum/go-ethereum/wiki/Management-APIs#txpool

//TODO: finish
func TxPoolContent() (interface{}, error) {
	resp, err := Call("txpool_content", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func TxPoolInspect() (interface{}, error) {
	resp, err := Call("txpool_inspect", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func TxPoolStatus() (interface{}, error) {
	resp, err := Call("txpool_status", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}
