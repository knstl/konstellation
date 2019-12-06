package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/supply/exported"
)

type CoinKeeper interface {
	GetCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	AddCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins) (sdk.Coins, sdk.Error)
	SubtractCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins) (sdk.Coins, sdk.Error)
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) sdk.Error
}

type SupplyKeeper interface {
	GetSupply(sdk.Context) exported.SupplyI
	SetSupply(sdk.Context, exported.SupplyI)
	MintCoins(sdk.Context, string, sdk.Coins) sdk.Error
	SendCoinsFromModuleToAccount(sdk.Context, string, sdk.AccAddress, sdk.Coins) sdk.Error
	//BurnCoins(sdk.Context, string, sdk.Coins) sdk.Error
}
