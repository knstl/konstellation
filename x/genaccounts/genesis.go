package genaccounts

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
)

const (
	ModuleName = genaccounts.ModuleName
)

// GenesisUpdater implements an genesis updater for the genaccounts module.
type GenesisUpdater struct{}

// Name returns the genaccounts module's name.
func (GenesisUpdater) Name() string {
	return ModuleName
}

// UpdateGenesis returns genesis state after changes as raw bytes for the genaccounts module.
func (GenesisUpdater) UpdateGenesis(cdc *codec.Codec, appState map[string]json.RawMessage) {

}
