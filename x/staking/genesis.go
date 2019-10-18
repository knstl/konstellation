package staking

import (
	"encoding/json"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/konstellation/konstellation/types"
)

const (
	ModuleName    = staking.ModuleName
	UnbondingTime = 60 * 10 * time.Second
)

// GenesisUpdater implements an genesis updater for the staking module.
type GenesisUpdater struct{}

// Name returns the staking module's name.
func (GenesisUpdater) Name() string {
	return ModuleName
}

// UpdateGenesis returns genesis state after changes as raw bytes for the staking module.
func (GenesisUpdater) UpdateGenesis(cdc *codec.Codec, appState map[string]json.RawMessage) {
	var genesisState staking.GenesisState
	err := cdc.UnmarshalJSON(appState[ModuleName], &genesisState)
	if err != nil {
		panic(err)
	}

	updateGenesisParams(&genesisState)

	appState[ModuleName] = cdc.MustMarshalJSON(genesisState)
}

func updateGenesisParams(genesisState *staking.GenesisState) {
	genesisState.Params.BondDenom = types.DefaultBondDenom
	genesisState.Params.UnbondingTime = UnbondingTime
}
