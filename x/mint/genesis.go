package mint

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/mint"

	"github.com/konstellation/konstellation/types"
)

const (
	ModuleName = mint.ModuleName
)

// GenesisUpdater implements an genesis updater for the mint module.
type GenesisUpdater struct{}

// Name returns the mint module's name
func (GenesisUpdater) Name() string {
	return ModuleName
}

// UpdateGenesis returns genesis state after changes as raw bytes for the mint module.
func (GenesisUpdater) UpdateGenesis(cdc *codec.Codec, appState map[string]json.RawMessage) {
	var genesisState mint.GenesisState
	err := cdc.UnmarshalJSON(appState[ModuleName], &genesisState)
	if err != nil {
		panic(err)
	}

	updateGenesisParams(&genesisState)

	appState[ModuleName] = cdc.MustMarshalJSON(genesisState)
}

func updateGenesisParams(genesisState *mint.GenesisState) {
	genesisState.Params.MintDenom = types.DefaultBondDenom
}
