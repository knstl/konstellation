package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/konstellation/konstellation/x/oracle/types"
)

// Keeper of the nameservice store
type Keeper struct {
	AuthKeeper           authkeeper.AccountKeeper
	BankKeeper           bankkeeper.Keeper
	exchangeRateStoreKey sdk.StoreKey
	cdc                  codec.Marshaler
	paramspace           paramtypes.Subspace
}

// NewKeeper creates a nameservice keeper
func NewKeeper(cdc codec.Marshaler, key sdk.StoreKey, paramSpace paramtypes.Subspace, bankKeeper bankkeeper.Keeper, authKeeper authkeeper.AccountKeeper) Keeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	keeper := Keeper{
		AuthKeeper:           authKeeper,
		BankKeeper:           bankKeeper,
		exchangeRateStoreKey: key,
		cdc:                  cdc,
		paramspace:           paramSpace,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
