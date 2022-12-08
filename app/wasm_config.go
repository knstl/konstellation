package app

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
)

const (
	// DefaultDarcInstanceCost is initially set the same as in wasmd
	DefaultDarcInstanceCost uint64 = 60_000
	// DefaultDarcCompileCost set to a large number for testing
	DefaultDarcCompileCost uint64 = 3
)

// JunoGasRegisterConfig is defaults plus a custom compile amount
func DarcGasRegisterConfig() wasmkeeper.WasmGasRegisterConfig {
	gasConfig := wasmkeeper.DefaultGasRegisterConfig()
	gasConfig.InstanceCost = DefaultDarcInstanceCost
	gasConfig.CompileCost = DefaultDarcCompileCost

	return gasConfig
}

func NewDarcWasmGasRegister() wasmkeeper.WasmGasRegister {
	return wasmkeeper.NewWasmGasRegister(DarcGasRegisterConfig())
}
