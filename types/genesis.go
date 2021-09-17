package types

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
)

// GenesisUpdater is the standard form for an application module
type GenesisUpdater interface {
	Name() string
	UpdateGenesis(cdc codec.JSONMarshaler, appState map[string]json.RawMessage)
}

// collections of GenesisUpdater
type GenesisUpdaters map[string]GenesisUpdater

func NewGenesisUpdaters(updaters ...GenesisUpdater) GenesisUpdaters {
	updaterMap := make(map[string]GenesisUpdater)
	for _, updater := range updaters {
		updaterMap[updater.Name()] = updater
	}
	return updaterMap
}
