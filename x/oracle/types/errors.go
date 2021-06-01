package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrAddressIsNotAllowed = sdkerrors.Register(ModuleName, 1, "address is not allowed")
	ErrAddressIsNotAdmin   = sdkerrors.Register(ModuleName, 2, "address is not admin")
)
