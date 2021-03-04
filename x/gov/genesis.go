package gov

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/konstellation/konstellation/types"
)

const (
	ModuleName       = govtypes.ModuleName
	MinDepositTokens = 1000
)

// GenesisUpdater implements an genesis updater for the gov module.
type GenesisUpdater struct{}

// Name returns the gov module's name.
func (GenesisUpdater) Name() string {
	return ModuleName
}

// UpdateGenesis returns genesis state after changes as raw bytes for the gov module.
func (GenesisUpdater) UpdateGenesis(cdc codec.JSONMarshaler, appState map[string]json.RawMessage) {
	var genesisState govtypes.GenesisState
	err := cdc.UnmarshalJSON(appState[ModuleName], &genesisState)
	if err != nil {
		panic(err)
	}

	updateGenesisParams(&genesisState)

	appState[ModuleName] = cdc.MustMarshalJSON(&genesisState)
}

func updateGenesisParams(genesisState *govtypes.GenesisState) {
	genesisState.DepositParams = govtypes.DepositParams{
		MinDeposit:       sdk.Coins{sdk.NewCoin(types.DefaultBondDenom, sdk.NewInt(MinDepositTokens))},
		MaxDepositPeriod: govtypes.DefaultPeriod,
	}
}
