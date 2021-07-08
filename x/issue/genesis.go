package issue

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/keeper"
	"github.com/konstellation/konstellation/x/issue/types"
)

// InitGenesis sets distribution information for genesis.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, gs types.GenesisState) {
	k.SetLastId(ctx, gs.StartingIssueId)
	k.SetParams(ctx, gs.Params)
	for _, issue := range gs.Issues {
		k.AddIssue(ctx, issue)
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	genesisState := types.GenesisState{}
	genesisState.StartingIssueId = k.GetLastId(ctx)
	genesisState.Params = k.GetParams(ctx)
	genesisState.Issues = k.ListAll(ctx)
	return genesisState
}

// ValidateGenesis performs basic validation of auth genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(gs types.GenesisState) error {
	return gs.Params.Validate()
}
