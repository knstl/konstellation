package staking

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
)

func InitGenesis(cdc *codec.Codec, appState map[string]json.RawMessage) {
	var genesisState genaccounts.GenesisState
	err := cdc.UnmarshalJSON(appState[genaccounts.ModuleName], &genesisState)
	if err != nil {
		panic(err)
	}

	appState[genaccounts.ModuleName] = cdc.MustMarshalJSON(genesisState)
}
