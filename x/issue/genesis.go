package issue

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis sets distribution information for genesis.
func InitGenesis(ctx sdk.Context, data GenesisState) {
	//err := keeper.SetInitialIssueStartingIssueId(ctx, data.StartingIssueId)
	//if err != nil {
	//	panic(err)
	//}
	//
	//keeper.SetParams(ctx, data.Params)
	//
	//for _, issue := range data.Issues {
	//	keeper.AddIssue(ctx, &issue)
	//}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context) GenesisState {
	genesisState := GenesisState{}
	return genesisState
}
