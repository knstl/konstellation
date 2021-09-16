package oracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/keeper"
	"github.com/konstellation/konstellation/x/oracle/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the adminAddr
	if err := k.SetAllowedAddressesInternal(ctx, genState.AllowedAddresses); err != nil {
		panic(err)
	}
	//	// this line is used by starport scaffolding # genesis/module/init
	//	// Set all the params
	//	for _, elem := range genState.ParamsList {
	//		k.SetParams(ctx, *elem)
	//	}
	//
	//	// Set params count
	//	k.SetParamsCount(ctx, genState.ParamsCount)
	//
	//
	//	// Set adminAddr count
	//	k.SetAdminAddrCount(ctx, genState.AdminAddrCount)
	//
	//	// Set all the exchangeRate
	//	for _, elem := range genState.ExchangeRateList {
	//		k.SetExchangeRate(ctx, *elem)
	//	}
	//
	//	// Set exchangeRate count
	//	k.SetExchangeRateCount(ctx, genState.ExchangeRateCount)
	//
	//	// this line is used by starport scaffolding # ibc/genesis/init
	//*/
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// Get all adminAddr
	adminAddrList := k.GetAllowedAddresses(ctx)
	for _, elem := range adminAddrList {
		elem := elem
		genesis.AllowedAddresses = append(genesis.AllowedAddresses, &elem)
	}
	/*
		// this line is used by starport scaffolding # genesis/module/export
		// Get all params
		paramsList := k.GetAllParams(ctx)
		for _, elem := range paramsList {
			elem := elem
			genesis.ParamsList = append(genesis.ParamsList, &elem)
		}

		// Set the current count
		genesis.ParamsCount = k.GetParamsCount(ctx)


		// Set the current count
		genesis.AdminAddrCount = k.GetAdminAddrCount(ctx)

		// Get all exchangeRate
		exchangeRateList := k.GetAllExchangeRate(ctx)
		for _, elem := range exchangeRateList {
			elem := elem
			genesis.ExchangeRateList = append(genesis.ExchangeRateList, &elem)
		}

		// Set the current count
		genesis.ExchangeRateCount = k.GetExchangeRateCount(ctx)

		// this line is used by starport scaffolding # ibc/genesis/export
	*/
	return genesis
}
