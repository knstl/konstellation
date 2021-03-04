package distribution

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

const (
	ModuleName = distrtypes.ModuleName
)

// GenesisUpdater implements an genesis updater for the distribution module.
type GenesisUpdater struct{}

// Name returns the distribution module's name.
func (GenesisUpdater) Name() string {
	return ModuleName
}

// UpdateGenesis returns genesis state after changes as raw bytes for the distribution module.
func (GenesisUpdater) UpdateGenesis(cdc codec.JSONMarshaler, appState map[string]json.RawMessage) {

}
