package EthereumAPI

import (
	"fmt"
)

//https://github.com/ethereum/wiki/wiki/JSON-RPC

//TODO: finish
func DBPutString(data []string) (bool, error) {
	resp, err := Call("db_putString", data)
	if err != nil {
		return false, err
	}
	if resp.Error != nil {
		return false, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(bool), nil
}

//TODO: finish
func DBGetString(data []string) (string, error) {
	resp, err := Call("db_getString", data)
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

//TODO: finish
func DBPutHex(data []string) (bool, error) {
	resp, err := Call("db_putHex", data)
	if err != nil {
		return false, err
	}
	if resp.Error != nil {
		return false, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(bool), nil
}

//TODO: finish
func DBGetHex(data []string) (string, error) {
	resp, err := Call("db_getHex", data)
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}
