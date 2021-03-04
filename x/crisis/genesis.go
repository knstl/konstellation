package crisis

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"

	"github.com/konstellation/konstellation/types"
)

const (
	ModuleName = crisistypes.ModuleName
)

// GenesisUpdater implements an genesis updater for the crisis module.
type GenesisUpdater struct{}

// Name returns the crisis module's name.
func (GenesisUpdater) Name() string {
	return ModuleName
}

// UpdateGenesis returns genesis state after changes as raw bytes for the crisis module.
func (gu GenesisUpdater) UpdateGenesis(cdc codec.JSONMarshaler, appState map[string]json.RawMessage) {
	var genesisState crisistypes.GenesisState
	err := cdc.UnmarshalJSON(appState[ModuleName], &genesisState)
	if err != nil {
		panic(err)
	}

	updateGenesisParams(&genesisState)

	appState[ModuleName] = cdc.MustMarshalJSON(&genesisState)
}

func updateGenesisParams(genesisState *crisistypes.GenesisState) {
	genesisState.ConstantFee = sdk.Coin{
		Denom:  types.DefaultBondDenom,
		Amount: sdk.TokensFromConsensusPower(types.DefaultConsensusPower),
	}
}
