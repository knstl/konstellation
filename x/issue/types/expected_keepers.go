package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authexported "github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/supply/exported"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authexported.Account
}

type CoinKeeper interface {
	GetCoins(sdk.Context, sdk.AccAddress) sdk.Coins
	AddCoins(sdk.Context, sdk.AccAddress, sdk.Coins) (sdk.Coins, sdk.Error)
	SubtractCoins(sdk.Context, sdk.AccAddress, sdk.Coins) (sdk.Coins, sdk.Error)
	SendCoins(sdk.Context, sdk.AccAddress, sdk.AccAddress, sdk.Coins) sdk.Error
	GetSendEnabled(sdk.Context) bool
	BlacklistedAddr(sdk.AccAddress) bool
}

type SupplyKeeper interface {
	GetSupply(sdk.Context) exported.SupplyI
	SetSupply(sdk.Context, exported.SupplyI)
	MintCoins(sdk.Context, string, sdk.Coins) sdk.Error
	BurnCoins(sdk.Context, string, sdk.Coins) sdk.Error
	SendCoinsFromModuleToAccount(sdk.Context, string, sdk.AccAddress, sdk.Coins) sdk.Error
	SendCoinsFromAccountToModule(sdk.Context, sdk.AccAddress, string, sdk.Coins) sdk.Error
}
