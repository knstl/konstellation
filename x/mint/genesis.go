package mint

import (
	"encoding/json"
	"github.com/konstellation/konstellation/types"

	"github.com/cosmos/cosmos-sdk/codec"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
)

const (
	ModuleName = minttypes.ModuleName
)

// GenesisUpdater implements an genesis updater for the mint module.
type GenesisUpdater struct{}

// Name returns the mint module's name
func (GenesisUpdater) Name() string {
	return ModuleName
}

// UpdateGenesis returns genesis state after changes as raw bytes for the mint module.
func (GenesisUpdater) UpdateGenesis(cdc codec.JSONMarshaler, appState map[string]json.RawMessage) {
	var genesisState minttypes.GenesisState
	err := cdc.UnmarshalJSON(appState[ModuleName], &genesisState)
	if err != nil {
		panic(err)
	}

	updateGenesisParams(&genesisState)

	appState[ModuleName] = cdc.MustMarshalJSON(&genesisState)
}

func updateGenesisParams(genesisState *minttypes.GenesisState) {
	genesisState.Params.MintDenom = types.DefaultBondDenom
}
