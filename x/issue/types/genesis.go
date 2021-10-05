package types

import (
	//"fmt"
	// this line is used by starport scaffolding # ibc/genesistype/import
	"bytes"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
)

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		StartingIssueId: InitLastId,
		Issues:          []*CoinIssue{},
		Params:          DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {

	return nil
}

// NewGenesisState - Create a new genesis state
func NewGenesisState(startingIssueId uint64, params Params) GenesisState {
	return GenesisState{
		StartingIssueId: startingIssueId,
		Params:          params,
	}
}

// Returns if a GenesisState is empty or has data in it
func (gs GenesisState) IsEmpty() bool {
	emptyGenState := GenesisState{}
	return gs.Equal(emptyGenState)
}

func (gs GenesisState) Equal(gs2 GenesisState) bool {
	bz1 := ModuleCdc.MustMarshalBinaryLengthPrefixed(&gs)
	bz2 := ModuleCdc.MustMarshalBinaryLengthPrefixed(&gs2)
	return bytes.Equal(bz1, bz2)
}

// GetGenesisStateFromAppState returns x/auth GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc *codec.LegacyAmino, appState map[string]json.RawMessage) GenesisState {
	var genesisState GenesisState
	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return genesisState
}
