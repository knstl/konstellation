package oracle

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/keeper"
	"github.com/konstellation/konstellation/x/oracle/types"
)

func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data *types.GenesisState) {
	fmt.Println("init gen")
	//keeper.SetAllowedAddresses()
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	allowedAddresses := keeper.GetAllowedAddresses(ctx)
	return types.NewGenesisState(allowedAddresses)
}
