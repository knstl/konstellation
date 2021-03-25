package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrAddressIsNotAllowed = sdkerrors.Register(ModuleName, 1, "address is not allowed")
)
