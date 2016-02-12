package EthereumAPI

import (
	"fmt"
)

//https://github.com/ethereum/wiki/wiki/JSON-RPC

func EthProtocolVersion() (string, error) {
	resp, err := Call("eth_protocolVersion", nil)
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

//TODO: finish
func EthSyncing() (interface{}, error) {
	resp, err := Call("eth_syncing", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func EthCoinbase() (interface{}, error) {
	resp, err := Call("eth_coinbase", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

func EthMining() (bool, error) {
	resp, err := Call("eth_mining", nil)
	if err != nil {
		return false, err
	}
	if resp.Error != nil {
		return false, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(bool), nil
}

//TODO: finish
func EthHashrate() (interface{}, error) {
	resp, err := Call("eth_hashrate", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func EthGasPrice() (interface{}, error) {
	resp, err := Call("eth_gasPrice", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func EthAccounts() (interface{}, error) {
	resp, err := Call("eth_accounts", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func EthBlockNumber() (interface{}, error) {
	resp, err := Call("eth_blockNumber", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func Eth() (interface{}, error) {
	resp, err := Call("XXXXXXXXXXXXXXXXX", nil)
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
func EthGetBalance() (interface{}, error) {
	resp, err:=Call("eth_getBalance", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func EthGetStorageAt() (interface{}, error) {
	resp, err:=Call("eth_getStorageAt", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func EthGetTransactionCount() (interface{}, error) {
	resp, err:=Call("eth_getTransactionCount", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func EthGetBlockTransactionCountByHash() (interface{}, error) {
	resp, err:=Call("eth_getBlockTransactionCountByHash", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func EthGetBlockTransactionCountByNumber() (interface{}, error) {
	resp, err:=Call("eth_getBlockTransactionCountByNumber", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func EthGetUncleCountByBlockHash() (interface{}, error) {
	resp, err:=Call("eth_getUncleCountByBlockHash", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthGetUncleCountByBlockNumber() (interface{}, error) {
	resp, err:=Call("eth_getUncleCountByBlockNumber", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthGetCode() (interface{}, error) {
	resp, err:=Call("eth_getCode", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthSign() (interface{}, error) {
	resp, err:=Call("eth_sign", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthSendTransaction() (interface{}, error) {
	resp, err:=Call("eth_sendTransaction", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthSendRawTransaction() (interface{}, error) {
	resp, err:=Call("eth_sendRawTransaction", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthCall() (interface{}, error) {
	resp, err:=Call("eth_call", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthEstimateGas() (interface{}, error) {
	resp, err:=Call("eth_estimateGas", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthGetBlockByHash() (interface{}, error) {
	resp, err:=Call("eth_getBlockByHash", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthGetBlockByNumber() (interface{}, error) {
	resp, err:=Call("eth_getBlockByNumber", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthGetTransactionByHash() (interface{}, error) {
	resp, err:=Call("eth_getTransactionByHash", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthGetTransactionByBlockHashAndIndex() (interface{}, error) {
	resp, err:=Call("eth_getTransactionByBlockHashAndIndex", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthGetTransactionByBlockNumberAndIndex() (interface{}, error) {
	resp, err:=Call("eth_getTransactionByBlockNumberAndIndex", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthGetTransactionReceipt() (interface{}, error) {
	resp, err:=Call("eth_getTransactionReceipt", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthGetUncleByBlockHashAndIndex() (interface{}, error) {
	resp, err:=Call("eth_getUncleByBlockHashAndIndex", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthGetUncleByBlockNumberAndIndex() (interface{}, error) {
	resp, err:=Call("eth_getUncleByBlockNumberAndIndex", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}
*/

func EthGetCompilers() ([]string, error) {
	resp, err := Call("eth_getCompilers", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.([]string), nil
}

/*
//TODO: finish
func EthCompileLLL() (interface{}, error) {
	resp, err:=Call("eth_compileLLL", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthCompileSolidity() (interface{}, error) {
	resp, err:=Call("eth_compileSolidity", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthCompileSerpent() (interface{}, error) {
	resp, err:=Call("eth_compileSerpent", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthNewFilter() (interface{}, error) {
	resp, err:=Call("eth_newFilter", nil)
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
func EthNewBlockFilter() (interface{}, error) {
	resp, err := Call("eth_newBlockFilter", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func EthNewPendingTransactionFilter() (interface{}, error) {
	resp, err := Call("eth_newPendingTransactionFilter", nil)
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
func EthUninstallFilter() (interface{}, error) {
	resp, err:=Call("eth_uninstallFilter", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthGetFilterChanges() (interface{}, error) {
	resp, err:=Call("eth_getFilterChanges", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthGetFilterLogs() (interface{}, error) {
	resp, err:=Call("eth_getFilterLogs", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthGetLogs() (interface{}, error) {
	resp, err:=Call("eth_getLogs", nil)
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
func EthGetWork() (interface{}, error) {
	resp, err := Call("eth_getWork", nil)
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
func EthSubmitWork() (interface{}, error) {
	resp, err:=Call("eth_submitWork", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}



//TODO: finish
func EthSubmitHashrate() (interface{}, error) {
	resp, err:=Call("eth_submitHashrate", nil)
	if err!=nil {
		return nil, err
	}
	if resp.Error!=nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

*/
