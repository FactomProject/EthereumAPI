package EthereumAPI

import (
	"fmt"
)

//https://github.com/ethereum/wiki/wiki/JSON-RPC

//TODO: test
func SSHPost(whisper *WhisperMessage) (bool, error) {
	resp, err := Call("shh_post", []interface{}{whisper})
	if err != nil {
		return false, err
	}
	if resp.Error != nil {
		return false, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(bool), nil
}

//TODO: test
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

//TODO: test
func SSHNewIdentity() (string, error) {
	resp, err := Call("shh_newIdentity", nil)
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

//TODO: test
func SSHHasIdentity(identityAddress string) (bool, error) {
	resp, err := Call("shh_hasIdentity", []interface{}{identityAddress})
	if err != nil {
		return false, err
	}
	if resp.Error != nil {
		return false, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(bool), nil
}

//TODO: test
func SSHNewGroup() (string, error) {
	resp, err := Call("shh_newGroup", nil)
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

//TODO: test
func SSHAddToGroup(identityAddress string) (interface{}, error) {
	resp, err := Call("shh_addToGroup", []interface{}{identityAddress})
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

//TODO: test
func SSHNewFilter(filter *FilterOptions) (int64, error) {
	resp, err := Call("shh_newFilter", []interface{}{filter})
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

//TODO: test
func SSHUninstallFilter(filterID string) (bool, error) {
	resp, err := Call("shh_uninstallFilter", []interface{}{filterID})
	if err != nil {
		return false, err
	}
	if resp.Error != nil {
		return false, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(bool), nil
}

//TODO: test
func SSHGetFilterChanges(filterID string) ([]*Message, error) {
	resp, err := Call("shh_getFilterChanges", []interface{}{filterID})
	if err != nil {
		return nil, err
	}
	answer := []*Message{}
	err = MapToObject(resp.Result, &answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

//TODO: test
func SSHGetMessages(filterID string) ([]*Message, error) {
	resp, err := Call("shh_getMessages", []interface{}{filterID})
	if err != nil {
		return nil, err
	}
	answer := []*Message{}
	err = MapToObject(resp.Result, &answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}
