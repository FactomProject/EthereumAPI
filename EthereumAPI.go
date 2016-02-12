package EthereumAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	/*
		C++: http://localhost:8545
		Go: http://localhost:8545
		Py: http://localhost:4000
	*/

	server = "localhost:8545"
)

//https://github.com/ethereum/wiki/wiki/JSON-RPC

func Call(method string, params interface{}) (*JSON2Response, error) {
	j := NewJSON2RequestBlank()
	j.Method = method
	j.Params = params
	j.ID = 1

	postGet := "POST"

	address := fmt.Sprintf("http://%s/", server)

	data, err := j.JSONString()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(postGet, address, strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	jResp := new(JSON2Response)

	err = json.Unmarshal(body, jResp)
	if err != nil {
		return nil, err
	}

	return jResp, nil
}
