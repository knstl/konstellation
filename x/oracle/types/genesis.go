package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	// TODO: change default actual address
	DefaultAllowedAddress = []byte("")
	DefaultExchangeRate   = []byte("")
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(allowedAddresses []string) *GenesisState {
	return &GenesisState{
		AllowedAddresses: allowedAddresses,
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		AllowedAddresses: []string{},
	}
}

func (s *GenesisState) ValidateBasic() error {
	//if len(s.AllowedAddresses) == 0 {
	//	return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "no addresses")
	//}
	for _, address := range s.AllowedAddresses {
		if len(address) == 0 {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "address is empty string")
		}
	}
	return nil
}

// ValidateGenesis performs basic validation of supply genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data *GenesisState) error {
	return data.ValidateBasic()
}
