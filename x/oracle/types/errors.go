package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrAddressIsNotAllowed = sdkerrors.Register(ModuleName, 1, "address is not allowed")
	ErrAddressIsNotAdmin   = sdkerrors.Register(ModuleName, 2, "address is not admin")
	ErrInvalidPair         = sdkerrors.Register(ModuleName, 3, "invalid pair")
	ErrInvalidRate         = sdkerrors.Register(ModuleName, 4, "invalid rate")
	ErrInvalidDenoms       = sdkerrors.Register(ModuleName, 5, "invalid denoms")
)
