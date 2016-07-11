package EthereumAPI

type TransactionObject struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Value    string `json:"value"`
	Data     string `json:"data"`
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

type EthSyncingResponse struct {
	Syncing       bool   `json:"syncing,omitempty"`
	StartingBlock string `json:"startingBlock,omitempty"`
	CurrentBlock  string `json:"currentBlock,omitempty"`
	HighestBlock  string `json:"highestBlock,omitempty"`
}
