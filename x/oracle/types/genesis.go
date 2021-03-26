package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ParamStoreKeyAllowedAddress = []byte("AllowedAddress")
	ParamStoreKeyExchangeRate   = []byte("ExchangeRate")
	DefaultAllowedAddress       = []byte("")
	DefaultExchangeRate         = []byte("")
)

// GenesisState - all oracle state that must be provided at genesis
type GenesisState struct {
	AllowedAddress sdk.AccAddress `json:"allowed_address"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState(allowedAddress sdk.AccAddress) *GenesisState {
	return &GenesisState{
		AllowedAddress: allowedAddress,
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	defaultAllowedAddress, _ := sdk.AccAddressFromHex(string(DefaultAllowedAddress))
	return GenesisState{
		AllowedAddress: defaultAllowedAddress,
	}
}

func (s GenesisState) ValidateBasic() error {
	if len(s.AllowedAddress.String()) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, s.AllowedAddress.String())
	}
	return nil
}

// ValidateGenesis performs basic validation of supply genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error {
	return data.ValidateBasic()
}
