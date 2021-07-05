package oracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/keeper"
	"github.com/konstellation/konstellation/x/oracle/types"
)

func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data *types.GenesisState) {
	//abc := types.AdminAddr{Address: "darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx"}
	//allowedAddresses := []*types.AdminAddr{&abc}
	// todo get data from genesis
	//keeper.SetAllowedAddresses(ctx, "darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx", allowedAddresses)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	//allowedAddresses := keeper.GetAllowedAddresses(ctx)
	return types.NewGenesisState(nil)
}
