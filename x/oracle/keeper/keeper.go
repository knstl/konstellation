package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/types"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc      codec.Codec
		storeKey sdk.Store
		memKey   sdk.Store
		// this line is used by starport scaffolding # ibc/keeper/attribute
		paramSpace paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.Codec,
	storeKey,
	memKey sdk.Store,
	// this line is used by starport scaffolding # ibc/keeper/parameter
	paramSpace paramtypes.Subspace,

) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		// this line is used by starport scaffolding # ibc/keeper/return
		paramSpace: paramSpace,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
