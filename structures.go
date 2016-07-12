package EthereumAPI

type EthSyncingResponse struct {
	Syncing       bool   `json:"syncing,omitempty"`
	StartingBlock string `json:"startingBlock,omitempty"`
	CurrentBlock  string `json:"currentBlock,omitempty"`
	HighestBlock  string `json:"highestBlock,omitempty"`
}

type TransactionObject struct {
	Hash             string `json:"hash,omitempty"`
	Nonce            string `json:"nonce,omitempty"`
	BlockHash        string `json:"blockHash,omitempty"`
	BlockNumber      string `json:"blockNumber,omitempty"`
	TransactionIndex string `json:"transactionIndex,omitempty"`

	From     string `json:"from"`
	To       string `json:"to"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Value    string `json:"value"`
	Data     string `json:"data,omitempty"`
	Input    string `json:"input,omitempty"`
}

type BlockObject struct {
	Number           string `json:"number"`
	Hash             string `json:"hash"`
	ParentHash       string `json:"parentHash"`
	Nonce            string `json:"nonce"`
	Sha3Uncles       string `json:"sha3Uncles"`
	LogsBloom        string `json:"logsBloom"`
	TransactionsRoot string `json:"transactionsRoot"`
	StateRoot        string `json:"stateRoot"`
	Miner            string `json:"miner"`
	Difficulty       string `json:"difficulty"`
	TotalDifficulty  string `json:"totalDifficulty"`
	ExtraData        string `json:"extraData"`
	Size             string `json:"size"`
	GasLimit         string `json:"gasLimit"`
	GasUsed          string `json:"gasUsed"`
	Timestamp        string `json:"timestamp"`
	//TODO: handle both full transactions and the hashes
	Transactions []interface{} `json:"transactions"`
	Uncles       []string      `json:"uncles"`
}

type TransactionReceipt struct {
	TransactionHash   string        `json:"transactionHash"`
	TransactionIndex  string        `json:"transactionIndex"`
	BlockHash         string        `json:"blockHash"`
	BlockNumber       string        `json:"blockNumber"`
	CumulativeGasUsed string        `json:"cumulativeGasUsed"`
	GasUsed           string        `json:"gasUsed"`
	ContractAddress   string        `json:"contractAddress"`
	Logs              []interface{} `json:"logs"`
}

type FilterOptions struct {
	FromBlock string `json:"fromBlock,omitempty"`
	ToBlock   string `json:"toBlock,omitempty"`
	Address   string `json:"address,omitempty"`
	Topics    string `json:"topics,omitempty"`
}

type LogObject struct {
	Type             string   `json:"type"`
	logIndex         string   `json:"logIndex"`
	TransactionIndex string   `json:"transactionIndex"`
	TransactionHash  string   `json:"transactionHash"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      string   `json:"blockNumber"`
	Address          string   `json:"address"`
	Data             string   `json:"data"`
	Topics           []string `json:"topics"`
}
