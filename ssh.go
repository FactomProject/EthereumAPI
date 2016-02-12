package EthereumAPI

import (
	"fmt"
)

//https://github.com/ethereum/wiki/wiki/JSON-RPC

/*

//TODO: finish
func SSHPost() (interface{}, error) {
	resp, err:=Call("shh_post", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}
*/

//TODO: finish
func SSHVersion() (string, error) {
	resp, err := Call("shh_version", nil)
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

//TODO: finish
func SSHNewIdentity() (interface{}, error) {
	resp, err := Call("shh_newIdentity", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

/*
//TODO: finish
func SSHHasIdentity() (interface{}, error) {
	resp, err:=Call("shh_hasIdentity", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}
*/

//TODO: finish
func SSHNewGroup() (interface{}, error) {
	resp, err := Call("shh_newGroup", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

/*
//TODO: finish
func SSHAddToGroup() (interface{}, error) {
	resp, err:=Call("shh_addToGroup", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}


//TODO: finish
func SSHNewFilter() (interface{}, error) {
	resp, err:=Call("shh_newFilter", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}


//TODO: finish
func SSHUninstallFilter() (interface{}, error) {
	resp, err:=Call("shh_uninstallFilter", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}


//TODO: finish
func SSHGetFilterChanges() (interface{}, error) {
	resp, err:=Call("shh_getFilterChanges", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}


//TODO: finish
func SSHGetMessages() (interface{}, error) {
	resp, err:=Call("shh_getMessages", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

*/
