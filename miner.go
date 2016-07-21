package EthereumAPI

import (
	"fmt"
)

//https://github.com/ethereum/go-ethereum/wiki/Management-APIs#miner

//TODO: finish
func MinerHashrate() (interface{}, error) {
	resp, err := Call("miner_hashrate", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func MinerMakeDAG(number int64) (interface{}, error) {
	resp, err := Call("miner_makeDAG", []interface{}{number})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func MinerSetExtra(extra string) (interface{}, error) {
	resp, err := Call("miner_setExtra", []interface{}{extra})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func MinerSetGasPrice(number int64) (interface{}, error) {
	resp, err := Call("miner_setGasPrice", []interface{}{number})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func MinerStart(threads int64) (interface{}, error) {
	resp, err := Call("miner_start", []interface{}{threads})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func MinerStartAutoDAG() (interface{}, error) {
	resp, err := Call("miner_startAutoDAG", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func MinerStop() (interface{}, error) {
	resp, err := Call("miner_stop", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func MinerStopAutoDAG() (interface{}, error) {
	resp, err := Call("miner_stopAutoDAG", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}
