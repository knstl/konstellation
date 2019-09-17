package staking

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/konstellation/konstellation/coin"
)

func InitGenesis(cdc *codec.Codec, appState map[string]json.RawMessage) {
	var genesisState staking.GenesisState
	err := cdc.UnmarshalJSON(appState[staking.ModuleName], &genesisState)
	if err != nil {
		panic(err)
	}

	updateBondDenom(&genesisState)

	appState[staking.ModuleName] = cdc.MustMarshalJSON(genesisState)
}

func updateBondDenom(genesisState *staking.GenesisState) {
	genesisState.Params.BondDenom = coin.DefaultBondDenom
}
