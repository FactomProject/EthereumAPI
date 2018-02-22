package EthereumAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//https://etherscan.io/apis
//https://testnet.etherscan.io/apis

//https://api.etherscan.io/api?module=account&action=txlist&address=0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae&startblock=0&endblock=99999999&sort=asc&apikey=YourApiKeyToken
//https://testnet.etherscan.io/api?module=account&action=txlist&address=0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae&startblock=0&endblock=99999999&sort=asc&apikey=YourApiKeyToken

var EtherscanTestNet bool = true
var EtherscanTestNetName string = "ropsten"
var EtherscanAPIKeyToken string = "UninitializedApiKey"

func CallRest(url string, dst interface{}) error {
	postGet := "GET"

	req, err := http.NewRequest(postGet, url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, dst)
	if err != nil {
		return err
	}

	return nil
}

type EtherscanResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type EtherscanTransaction struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	TransactionIndex  string `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	IsError           string `json:"isError"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	Confirmations     string `json:"confirmations"`
}

func (e *EtherscanTransaction) JSONByte() ([]byte, error) {
	return EncodeJSON(e)
}

func (e *EtherscanTransaction) JSONString() (string, error) {
	return EncodeJSONString(e)
}

func (e *EtherscanTransaction) JSONBuffer(b *bytes.Buffer) error {
	return EncodeJSONToBuffer(e, b)
}

func (e *EtherscanTransaction) String() string {
	str, _ := e.JSONString()
	return str
}

func EtherscanTxList(address string) ([]*EtherscanTransaction, error) {
	prefix := "api"
	if EtherscanTestNet == true {
		prefix = EtherscanTestNetName
	}

	url := fmt.Sprintf("https://%v.etherscan.io/api?module=account&action=txlist&address=%v&startblock=0&endblock=99999999&sort=asc&apikey=%v", prefix, address, EtherscanAPIKeyToken)
	dst := new(EtherscanResponse)
	dst.Result = []*EtherscanTransaction{}
	err := CallRest(url, dst)
	if err != nil {
		return nil, err
	}
	list := []*EtherscanTransaction{}
	err = MapToObject(dst.Result, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func EtherscanTxListWithStartBlock(address string, startBlock int64) ([]*EtherscanTransaction, error) {
	prefix := "api"
	if EtherscanTestNet == true {
		prefix = EtherscanTestNetName
	}
	url := fmt.Sprintf("https://%v.etherscan.io/api?module=account&action=txlist&address=%v&startblock=%v&endblock=99999999&sort=asc&apikey=%v", prefix, address, startBlock, EtherscanAPIKeyToken)
	dst := new(EtherscanResponse)
	dst.Result = []*EtherscanTransaction{}
	err := CallRest(url, dst)
	if err != nil {
		return nil, err
	}
	list := []*EtherscanTransaction{}
	err = MapToObject(dst.Result, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
