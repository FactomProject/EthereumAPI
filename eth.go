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

func EthSyncing() (*EthSyncingResponse, error) {
	resp, err := Call("eth_syncing", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	answer := new(EthSyncingResponse)
	syncing, ok := resp.Result.(bool)
	if ok == true {
		answer.Syncing = syncing
		return answer, nil
	}
	err = MapToObject(resp.Result, answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

func EthCoinbase() (string, error) {
	resp, err := Call("eth_coinbase", nil)
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
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

func EthHashrate() (int64, error) {
	resp, err := Call("eth_hashrate", nil)
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

func EthGasPrice() (int64, error) {
	resp, err := Call("eth_gasPrice", nil)
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

func EthAccounts() ([]string, error) {
	resp, err := Call("eth_accounts", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	answer := []string{}
	err = MapToObject(resp.Result, &answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

func EthBlockNumber() (int64, error) {
	resp, err := Call("eth_blockNumber", nil)
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

//TODO: test
func EthGetBalance(address string, blockNumberOrTag string) (int64, error) {
	resp, err := Call("eth_getBalance", []string{address, blockNumberOrTag})
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

//TODO: test
func EthGetStorageAt(address, positionInTheStorage, blockNumberOrTag string) (string, error) {
	resp, err := Call("eth_getStorageAt", []string{address, positionInTheStorage, blockNumberOrTag})
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

//TODO: test
func EthGetTransactionCount(address, blockNumberOrTag string) (int64, error) {
	resp, err := Call("eth_getTransactionCount", []string{address, blockNumberOrTag})
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

//TODO: test
func EthGetBlockTransactionCountByHash(blockHash string) (int64, error) {
	resp, err := Call("eth_getBlockTransactionCountByHash", []string{blockHash})
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

//TODO: test
func EthGetBlockTransactionCountByNumber(blockNumberOrTag string) (int64, error) {
	resp, err := Call("eth_getBlockTransactionCountByNumber", []string{blockNumberOrTag})
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

//TODO: test
func EthGetUncleCountByBlockHash(blockHash string) (int64, error) {
	resp, err := Call("eth_getUncleCountByBlockHash", []string{blockHash})
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

//TODO: test
func EthGetUncleCountByBlockNumber(blockNumberOrTag string) (int64, error) {
	resp, err := Call("eth_getUncleCountByBlockNumber", []string{blockNumberOrTag})
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

//TODO: test
func EthGetCode(address, blockNumberOrTag string) (string, error) {
	resp, err := Call("eth_getCode", []string{address, blockNumberOrTag})
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

//TODO: test
func EthSign(address, hashOfDataToSign string) (string, error) {
	resp, err := Call("eth_sign", []string{address, hashOfDataToSign})
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

//TODO: test
func EthSendTransaction(tx *TransactionObject) (string, error) {
	resp, err := Call("eth_sendTransaction", []interface{}{tx})
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

//TODO: test
func EthSendRawTransaction(txData string) (string, error) {
	resp, err := Call("eth_sendRawTransaction", []string{txData})
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

//TODO: test
func EthCall(tx *TransactionObject, blockNumberOrTag string) (string, error) {
	resp, err := Call("eth_call", []interface{}{tx, blockNumberOrTag})
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

//TODO: test
func EthEstimateGas(tx *TransactionObject, blockNumberOrTag string) (int64, error) {
	resp, err := Call("eth_estimateGas", []interface{}{tx, blockNumberOrTag})
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

//TODO: test
func EthGetBlockByHash(blockHash string, fullTransaction bool) (*BlockObject, error) {
	resp, err := Call("eth_getBlockByHash", []interface{}{blockHash, fullTransaction})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	answer := new(BlockObject)
	err = MapToObject(resp.Result, answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

//TODO: test
func EthGetBlockByNumber(blockNumberOrTag string, fullTransaction bool) (*BlockObject, error) {
	resp, err := Call("eth_getBlockByNumber", []interface{}{blockNumberOrTag, fullTransaction})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	answer := new(BlockObject)
	err = MapToObject(resp.Result, answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

//TODO: test
func EthGetTransactionByHash(txHash string) (*TransactionObject, error) {
	resp, err := Call("eth_getTransactionByHash", []interface{}{txHash})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	answer := new(TransactionObject)
	err = MapToObject(resp.Result, answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

//TODO: test
func EthGetTransactionByBlockHashAndIndex(blockHash, txIndex string) (*TransactionObject, error) {
	resp, err := Call("eth_getTransactionByBlockHashAndIndex", []interface{}{blockHash, txIndex})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	answer := new(TransactionObject)
	err = MapToObject(resp.Result, answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

//TODO: test
func EthGetTransactionByBlockNumberAndIndex(blockNumberOrTag string, txIndex string) (*TransactionObject, error) {
	resp, err := Call("eth_getTransactionByBlockNumberAndIndex", []interface{}{blockNumberOrTag, txIndex})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	answer := new(TransactionObject)
	err = MapToObject(resp.Result, answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

//TODO: test
func EthGetTransactionReceipt(txHash string) (*TransactionReceipt, error) {
	resp, err := Call("eth_getTransactionReceipt", []interface{}{txHash})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	answer := new(TransactionReceipt)
	err = MapToObject(resp.Result, answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

//TODO: test
func EthGetUncleByBlockHashAndIndex(blockHash, uncleIndex string) (*BlockObject, error) {
	resp, err := Call("eth_getUncleByBlockHashAndIndex", []interface{}{blockHash, uncleIndex})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	answer := new(BlockObject)
	err = MapToObject(resp.Result, answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

//TODO: test
func EthGetUncleByBlockNumberAndIndex(blockNumberOrTag string, uncleIndex string) (*BlockObject, error) {
	resp, err := Call("eth_getUncleByBlockNumberAndIndex", []interface{}{blockNumberOrTag, uncleIndex})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	answer := new(BlockObject)
	err = MapToObject(resp.Result, answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

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

//TODO: test
func EthCompileLLL(sourceCode string) (string, error) {
	resp, err := Call("eth_compileLLL", []interface{}{sourceCode})
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

//TODO: test
func EthCompileSolidity(sourceCode string) (string, error) {
	resp, err := Call("eth_compileSolidity", []interface{}{sourceCode})
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

//TODO: test
func EthCompileSerpent(sourceCode string) (string, error) {
	resp, err := Call("eth_compileSerpent", []interface{}{sourceCode})
	if err != nil {
		return "", err
	}
	if resp.Error != nil {
		return "", fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(string), nil
}

//TODO: test
func EthNewFilter(filterOptions *FilterOptions) (int64, error) {
	resp, err := Call("eth_newFilter", []interface{}{filterOptions})
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

//TODO: test
func EthNewBlockFilter() (int64, error) {
	resp, err := Call("eth_newBlockFilter", nil)
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

//TODO: test
func EthNewPendingTransactionFilter() (int64, error) {
	resp, err := Call("eth_newPendingTransactionFilter", nil)
	if err != nil {
		return 0, err
	}
	if resp.Error != nil {
		return 0, fmt.Errorf(resp.Error.Message)
	}
	return ParseQuantity(resp.Result.(string))
}

//TODO: test
func EthUninstallFilter(filterID string) (bool, error) {
	resp, err := Call("eth_uninstallFilter", []interface{}{filterID})
	if err != nil {
		return false, err
	}
	if resp.Error != nil {
		return false, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(bool), nil
}

//TODO: test
func EthGetFilterChanges(filterID string) (*LogObject, error) {
	resp, err := Call("eth_getFilterChanges", []interface{}{filterID})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	answer := new(LogObject)
	err = MapToObject(resp.Result, answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

//TODO: test
func EthGetFilterLogs(filterID string) (*LogObject, error) {
	resp, err := Call("eth_getFilterLogs", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	answer := new(LogObject)
	err = MapToObject(resp.Result, answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

//TODO: test
func EthGetLogs(filter []*FilterOptions) (*LogObject, error) {
	resp, err := Call("eth_getLogs", filter)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	answer := new(LogObject)
	err = MapToObject(resp.Result, answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

//TODO: test
func EthGetWork() ([]string, error) {
	resp, err := Call("eth_getWork", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	answer := []string{}
	err = MapToObject(resp.Result, &answer)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

//TODO: test
func EthSubmitWork(work []string) (bool, error) {
	resp, err := Call("eth_submitWork", work)
	if err != nil {
		return false, err
	}
	if resp.Error != nil {
		return false, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(bool), nil
}

//TODO: test
func EthSubmitHashrate(hashrate []string) (bool, error) {
	resp, err := Call("eth_submitHashrate", hashrate)
	if err != nil {
		return false, err
	}
	if resp.Error != nil {
		return false, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result.(bool), nil
}
