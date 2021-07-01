package types

import (
	"bytes"
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
)

// GenesisState - all auth state that must be provided at genesis
type GenesisState struct {
	StartingIssueId uint64     `json:"starting_issue_id" yaml:"starting_issue_id"`
	Issues          CoinIssues `json:"issues" yaml:"issues"`
	Params          Params     `json:"params" yaml:"params"`
}

// NewGenesisState - Create a new genesis state
func NewGenesisState(startingIssueId uint64, params Params) GenesisState {
	return GenesisState{
		StartingIssueId: startingIssueId,
		Params:          params,
	}
}

// DefaultGenesisState - Return a default genesis state
func DefaultGenesisState() GenesisState {
	return GenesisState{
		StartingIssueId: InitLastId,
		Params:          DefaultParams(),
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
func GetGenesisStateFromAppState(cdc *codec.Codec, appState map[string]json.RawMessage) GenesisState {
	var genesisState GenesisState
	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return genesisState
}
