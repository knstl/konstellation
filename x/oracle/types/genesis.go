package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ParamStoreKeyAllowedAddress = []byte("AllowedAddress")
	ParamStoreKeyExchangeRate   = []byte("ExchangeRate")
	// TODO: change default actual address
	DefaultAllowedAddress = []byte("abc")
	DefaultExchangeRate   = []byte("")
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(allowedAddress string) *GenesisState {
	return &GenesisState{
		AllowedAddress: allowedAddress,
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		AllowedAddress: string(DefaultAllowedAddress),
	}
}

func (s *GenesisState) ValidateBasic() error {
	if len(s.AllowedAddress) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, s.AllowedAddress)
	}
	return nil
}

// ValidateGenesis performs basic validation of supply genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data *GenesisState) error {
	return data.ValidateBasic()
}
