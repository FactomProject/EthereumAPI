package EthereumAPI

import (
	"fmt"
)

//https://github.com/ethereum/go-ethereum/wiki/Management-APIs#admin

//TODO: finish
func AdminAddPeer(url string) (interface{}, error) {
	resp, err := Call("admin_addPeer", []interface{}{url})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func AdminDatadir() (interface{}, error) {
	resp, err := Call("admin_datadir", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func AdminNodeInfo() (interface{}, error) {
	resp, err := Call("admin_nodeInfo", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func AdminPeers() (interface{}, error) {
	resp, err := Call("admin_peers", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func AdminSetSolc(path string) (interface{}, error) {
	resp, err := Call("admin_setSolc", []interface{}{path})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func AdminStartRPC(host string, port int64, cors string, apis string) (interface{}, error) {
	resp, err := Call("admin_startRPC", []interface{}{host, port, cors, apis})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func AdminStartWS(host string, port int64, cors string, apis string) (interface{}, error) {
	resp, err := Call("admin_startWS", []interface{}{host, port, cors, apis})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func AdminStopRPC() (interface{}, error) {
	resp, err := Call("admin_stopRPC", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func AdminStopWS() (interface{}, error) {
	resp, err := Call("admin_stopWS", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}
