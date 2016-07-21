package EthereumAPI

import (
	"fmt"
)

//https://github.com/ethereum/go-ethereum/wiki/Management-APIs#debug

//TODO: finish
func DebugBacktraceAt(location string) (interface{}, error) {
	resp, err := Call("debug_backtraceAt", []interface{}{location})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func DebugBlockProfile(file string, seconds int64) (interface{}, error) {
	resp, err := Call("debug_blockProfile", []interface{}{file, seconds})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func DebugCPUProfile(file string, seconds int64) (interface{}, error) {
	resp, err := Call("debug_cpuProfile", []interface{}{file, seconds})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func DebugDumpBlock(number uint64) (interface{}, error) {
	resp, err := Call("debug_dumpBlock", []interface{}{number})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func DebugGCStats() (interface{}, error) {
	resp, err := Call("debug_gcStats", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func DebugGetBlockRLP(number uint64) (interface{}, error) {
	resp, err := Call("debug_getBlockRlp", []interface{}{number})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func DebugGoTrace(file string, seconds int64) (interface{}, error) {
	resp, err := Call("debug_goTrace", []interface{}{file, seconds})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func DebugMemStats() (interface{}, error) {
	resp, err := Call("debug_memStats", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func DebugSeedHash(number uint64) (interface{}, error) {
	resp, err := Call("debug_seedHash", []interface{}{number})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func DebugSetHead(number uint64) (interface{}, error) {
	resp, err := Call("debug_setHead", []interface{}{number})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func DebugSetBlockProfileRate(rate int64) (interface{}, error) {
	resp, err := Call("debug_setBlockProfileRate", []interface{}{rate})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func DebugStacks() (interface{}, error) {
	resp, err := Call("debug_stacks", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func DebugStartCPUProfile(file string) (interface{}, error) {
	resp, err := Call("debug_startCPUProfile", []interface{}{file})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func DebugStartGoTrace(file string) (interface{}, error) {
	resp, err := Call("debug_startGoTrace", []interface{}{file})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func DebugStopCPUProfile() (interface{}, error) {
	resp, err := Call("debug_stopCPUProfile", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func DebugStopGoTrace() (interface{}, error) {
	resp, err := Call("debug_stopGoTrace", nil)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
//TODO: add config?
func DebugTraceBlock(blockRlp string) (interface{}, error) {
	resp, err := Call("debug_traceBlock", []interface{}{blockRlp})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
//TODO: add config?
func DebugTraceBlockByNumber(number uint64) (interface{}, error) {
	resp, err := Call("debug_traceBlockByNumber", []interface{}{number})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
//TODO: add config?
func DebugTraceBlockByHash(hash string) (interface{}, error) {
	resp, err := Call("debug_traceBlockByHash", []interface{}{hash})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
//TODO: add config?
func DebugTraceBlockFromFile(file string) (interface{}, error) {
	resp, err := Call("debug_traceBlockFromFile", []interface{}{file})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
//TODO: add options?
func DebugTraceTransaction(txHash string) (interface{}, error) {
	resp, err := Call("debug_traceTransaction", []interface{}{txHash})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func DebugVerbosity(number int64) (interface{}, error) {
	resp, err := Call("debug_verbosity", []interface{}{number})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func DebugVModule(pattern string) (interface{}, error) {
	resp, err := Call("debug_vmodule", []interface{}{pattern})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func DebugWriteBlockProfile(file string) (interface{}, error) {
	resp, err := Call("debug_writeBlockProfile", []interface{}{file})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}

//TODO: finish
func DebugWriteMemProfile(file string) (interface{}, error) {
	resp, err := Call("debug_writeMemProfile", []interface{}{file})
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf(resp.Error.Message)
	}
	return resp.Result, nil
}
