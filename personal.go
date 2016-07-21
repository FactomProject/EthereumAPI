package EthereumAPI

import (
	"fmt"
)

//https://github.com/ethereum/go-ethereum/wiki/Management-APIs#personal

//TODO: finish
func PersonalImportRawKey(keyData, passphrase string) (interface{}, error) {
	resp, err := Call("personal_importRawKey", []interface{}{keyData, passphrase})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func PersonalListAccounts() (interface{}, error) {
	resp, err := Call("personal_listAccounts", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func PersonalLockAccount(address string) (interface{}, error) {
	resp, err := Call("personal_lockAccount", []interface{}{address})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func PersonalNewAccount(passphrase string) (interface{}, error) {
	resp, err := Call("personal_newAccount", []interface{}{passphrase})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func PersonalUnlockAccount(address, passphrase string, duration int64) (interface{}, error) {
	resp, err := Call("personal_unlockAccount", []interface{}{address, passphrase, duration})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func PersonalSignAndSendTransaction(tx *TransactionObject, passphrase string) (string, error) {
	resp, err := Call("personal_signAndSendTransaction", []interface{}{tx, passphrase})
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}
