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
	for _, issue := range gs.Issues.Issues {
		k.AddIssue(ctx, issue)
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	genesisState := types.GenesisState{}
	genesisState.StartingIssueId = k.GetLastId(ctx)
	genesisState.Params = k.GetParams(ctx)
	issueList := types.Issues{
		Issues: []*types.CoinIssue{},
	}
	for _, issue := range k.ListAll(ctx) {
		issueList.Issues = append(issueList.Issues, issue)
	}
	genesisState.Issues = &issueList
	return genesisState
}

// ValidateGenesis performs basic validation of auth genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(gs types.GenesisState) error {
	return gs.Params.Validate()
}
