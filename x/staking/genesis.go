package staking

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/konstellation/konstellation/types"
	"time"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

const (
	ModuleName    = stakingtypes.ModuleName
	UnbondingTime = 60 * 10 * time.Second
)

// GenesisUpdater implements an genesis updater for the staking module.
type GenesisUpdater struct{}

// Name returns the staking module's name.
func (GenesisUpdater) Name() string {
	return ModuleName
}

// UpdateGenesis returns genesis state after changes as raw bytes for the staking module.
func (GenesisUpdater) UpdateGenesis(cdc codec.JSONMarshaler, appState map[string]json.RawMessage) {
	var genesisState stakingtypes.GenesisState
	err := cdc.UnmarshalJSON(appState[ModuleName], &genesisState)
	if err != nil {
		panic(err)
	}

	updateGenesisParams(&genesisState)

	appState[ModuleName] = cdc.MustMarshalJSON(&genesisState)
}

func updateGenesisParams(genesisState *stakingtypes.GenesisState) {
	genesisState.Params.BondDenom = types.DefaultBondDenom
	genesisState.Params.UnbondingTime = UnbondingTime
}
