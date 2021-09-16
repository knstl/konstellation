package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/oracle module sentinel errors
var (
	//ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
	// this line is used by starport scaffolding # ibc/errors

	ErrAddressIsNotAllowed = sdkerrors.Register(ModuleName, 1, "address is not allowed")
	ErrAddressIsNotAdmin   = sdkerrors.Register(ModuleName, 2, "address is not admin")
	ErrInvalidPair         = sdkerrors.Register(ModuleName, 3, "invalid pair")
	ErrInvalidRate         = sdkerrors.Register(ModuleName, 4, "invalid rate")
	ErrInvalidDenoms       = sdkerrors.Register(ModuleName, 5, "invalid denoms")
	ErrPairDoesntExist     = sdkerrors.Register(ModuleName, 6, "pair doesn't exist")
	ErrInvalidPairs        = sdkerrors.Register(ModuleName, 7, "invalid pairs")
)
