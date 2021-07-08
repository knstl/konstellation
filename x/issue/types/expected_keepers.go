package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank/exported"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authexported.AccountI
}

type CoinKeeper interface {
	GetCoins(sdk.Context, sdk.AccAddress) sdk.Coins
	AddCoins(sdk.Context, sdk.AccAddress, sdk.Coins) (sdk.Coins, *sdkerrors.Error)
	SubtractCoins(sdk.Context, sdk.AccAddress, sdk.Coins) (sdk.Coins, *sdkerrors.Error)
	SendCoins(sdk.Context, sdk.AccAddress, sdk.AccAddress, sdk.Coins) *sdkerrors.Error
	GetSendEnabled(sdk.Context) bool
	BlacklistedAddr(sdk.AccAddress) bool
}

type SupplyKeeper interface {
	GetSupply(sdk.Context) exported.SupplyI
	SetSupply(sdk.Context, exported.SupplyI)
	MintCoins(sdk.Context, string, sdk.Coins) *sdkerrors.Error
	BurnCoins(sdk.Context, string, sdk.Coins) *sdkerrors.Error
	SendCoinsFromModuleToAccount(sdk.Context, string, sdk.AccAddress, sdk.Coins) *sdkerrors.Error
	SendCoinsFromAccountToModule(sdk.Context, sdk.AccAddress, string, sdk.Coins) *sdkerrors.Error
}
