package issue

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/keeper"
	"github.com/konstellation/konstellation/x/issue/types"
)

/*
// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the coinIssueDenoms
	for _, elem := range genState.CoinIssueDenomsList {
		k.SetCoinIssueDenoms(ctx, *elem)
	}

	// Set coinIssueDenoms count
	k.SetCoinIssueDenomsCount(ctx, genState.CoinIssueDenomsCount)

	// Set all the coinIssueList
	for _, elem := range genState.CoinIssueListList {
		k.SetCoinIssueList(ctx, *elem)
	}

	// Set coinIssueList count
	k.SetCoinIssueListCount(ctx, genState.CoinIssueListCount)

	// Set all the issuesParams
	for _, elem := range genState.IssuesParamsList {
		k.SetIssuesParams(ctx, *elem)
	}

	// Set issuesParams count
	k.SetIssuesParamsCount(ctx, genState.IssuesParamsCount)

	// Set all the issueParams
	for _, elem := range genState.IssueParamsList {
		k.SetIssueParams(ctx, *elem)
	}

	// Set issueParams count
	k.SetIssueParamsCount(ctx, genState.IssueParamsCount)

	// Set all the issueFeatures
	for _, elem := range genState.IssueFeaturesList {
		k.SetIssueFeatures(ctx, *elem)
	}

	// Set issueFeatures count
	k.SetIssueFeaturesCount(ctx, genState.IssueFeaturesCount)

	// Set all the params
	for _, elem := range genState.ParamsList {
		k.SetParams(ctx, *elem)
	}

	// Set params count
	k.SetParamsCount(ctx, genState.ParamsCount)

	// Set all the issues
	for _, elem := range genState.IssuesList {
		k.SetIssues(ctx, *elem)
	}

	// Set issues count
	k.SetIssuesCount(ctx, genState.IssuesCount)

	// Set all the coinIssueDenom
	for _, elem := range genState.CoinIssueDenomList {
		k.SetCoinIssueDenom(ctx, *elem)
	}

	// Set coinIssueDenom count
	k.SetCoinIssueDenomCount(ctx, genState.CoinIssueDenomCount)

	// Set all the coinIssue
	for _, elem := range genState.CoinIssueList {
		k.SetCoinIssue(ctx, *elem)
	}

	// Set coinIssue count
	k.SetCoinIssueCount(ctx, genState.CoinIssueCount)

	// Set all the addressFreezeList
	for _, elem := range genState.AddressFreezeListList {
		k.SetAddressFreezeList(ctx, *elem)
	}

	// Set addressFreezeList count
	k.SetAddressFreezeListCount(ctx, genState.AddressFreezeListCount)

	// Set all the addressFreeze
	for _, elem := range genState.AddressFreezeList {
		k.SetAddressFreeze(ctx, *elem)
	}

	// Set addressFreeze count
	k.SetAddressFreezeCount(ctx, genState.AddressFreezeCount)

	// Set all the freeze
	for _, elem := range genState.FreezeList {
		k.SetFreeze(ctx, *elem)
	}

	// Set freeze count
	k.SetFreezeCount(ctx, genState.FreezeCount)

	// Set all the coins
	for _, elem := range genState.CoinsList {
		k.SetCoins(ctx, *elem)
	}

	// Set coins count
	k.SetCoinsCount(ctx, genState.CoinsCount)

	// Set all the allowanceList
	for _, elem := range genState.AllowanceListList {
		k.SetAllowanceList(ctx, *elem)
	}

	// Set allowanceList count
	k.SetAllowanceListCount(ctx, genState.AllowanceListCount)

	// Set all the allowance
	for _, elem := range genState.AllowanceList {
		k.SetAllowance(ctx, *elem)
	}

	// Set allowance count
	k.SetAllowanceCount(ctx, genState.AllowanceCount)

	// Set all the address
	for _, elem := range genState.AddressList {
		k.SetAddress(ctx, *elem)
	}

	// Set address count
	k.SetAddressCount(ctx, genState.AddressCount)

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all coinIssueDenoms
	coinIssueDenomsList := k.GetAllCoinIssueDenoms(ctx)
	for _, elem := range coinIssueDenomsList {
		elem := elem
		genesis.CoinIssueDenomsList = append(genesis.CoinIssueDenomsList, &elem)
	}

	// Set the current count
	genesis.CoinIssueDenomsCount = k.GetCoinIssueDenomsCount(ctx)

	// Get all coinIssueList
	coinIssueListList := k.GetAllCoinIssueList(ctx)
	for _, elem := range coinIssueListList {
		elem := elem
		genesis.CoinIssueListList = append(genesis.CoinIssueListList, &elem)
	}

	// Set the current count
	genesis.CoinIssueListCount = k.GetCoinIssueListCount(ctx)

	// Get all issuesParams
	issuesParamsList := k.GetAllIssuesParams(ctx)
	for _, elem := range issuesParamsList {
		elem := elem
		genesis.IssuesParamsList = append(genesis.IssuesParamsList, &elem)
	}

	// Set the current count
	genesis.IssuesParamsCount = k.GetIssuesParamsCount(ctx)

	// Get all issueParams
	issueParamsList := k.GetAllIssueParams(ctx)
	for _, elem := range issueParamsList {
		elem := elem
		genesis.IssueParamsList = append(genesis.IssueParamsList, &elem)
	}

	// Set the current count
	genesis.IssueParamsCount = k.GetIssueParamsCount(ctx)

	// Get all issueFeatures
	issueFeaturesList := k.GetAllIssueFeatures(ctx)
	for _, elem := range issueFeaturesList {
		elem := elem
		genesis.IssueFeaturesList = append(genesis.IssueFeaturesList, &elem)
	}

	// Set the current count
	genesis.IssueFeaturesCount = k.GetIssueFeaturesCount(ctx)

	// Get all params
	paramsList := k.GetAllParams(ctx)
	for _, elem := range paramsList {
		elem := elem
		genesis.ParamsList = append(genesis.ParamsList, &elem)
	}

	// Set the current count
	genesis.ParamsCount = k.GetParamsCount(ctx)

	// Get all issues
	issuesList := k.GetAllIssues(ctx)
	for _, elem := range issuesList {
		elem := elem
		genesis.IssuesList = append(genesis.IssuesList, &elem)
	}

	// Set the current count
	genesis.IssuesCount = k.GetIssuesCount(ctx)

	// Get all coinIssueDenom
	coinIssueDenomList := k.GetAllCoinIssueDenom(ctx)
	for _, elem := range coinIssueDenomList {
		elem := elem
		genesis.CoinIssueDenomList = append(genesis.CoinIssueDenomList, &elem)
	}

	// Set the current count
	genesis.CoinIssueDenomCount = k.GetCoinIssueDenomCount(ctx)

	// Get all coinIssue
	coinIssueList := k.GetAllCoinIssue(ctx)
	for _, elem := range coinIssueList {
		elem := elem
		genesis.CoinIssueList = append(genesis.CoinIssueList, &elem)
	}

	// Set the current count
	genesis.CoinIssueCount = k.GetCoinIssueCount(ctx)

	// Get all addressFreezeList
	addressFreezeListList := k.GetAllAddressFreezeList(ctx)
	for _, elem := range addressFreezeListList {
		elem := elem
		genesis.AddressFreezeListList = append(genesis.AddressFreezeListList, &elem)
	}

	// Set the current count
	genesis.AddressFreezeListCount = k.GetAddressFreezeListCount(ctx)

	// Get all addressFreeze
	addressFreezeList := k.GetAllAddressFreeze(ctx)
	for _, elem := range addressFreezeList {
		elem := elem
		genesis.AddressFreezeList = append(genesis.AddressFreezeList, &elem)
	}

	// Set the current count
	genesis.AddressFreezeCount = k.GetAddressFreezeCount(ctx)

	// Get all freeze
	freezeList := k.GetAllFreeze(ctx)
	for _, elem := range freezeList {
		elem := elem
		genesis.FreezeList = append(genesis.FreezeList, &elem)
	}

	// Set the current count
	genesis.FreezeCount = k.GetFreezeCount(ctx)

	// Get all coins
	coinsList := k.GetAllCoins(ctx)
	for _, elem := range coinsList {
		elem := elem
		genesis.CoinsList = append(genesis.CoinsList, &elem)
	}

	// Set the current count
	genesis.CoinsCount = k.GetCoinsCount(ctx)

	// Get all allowanceList
	allowanceListList := k.GetAllAllowanceList(ctx)
	for _, elem := range allowanceListList {
		elem := elem
		genesis.AllowanceListList = append(genesis.AllowanceListList, &elem)
	}

	// Set the current count
	genesis.AllowanceListCount = k.GetAllowanceListCount(ctx)

	// Get all allowance
	allowanceList := k.GetAllAllowance(ctx)
	for _, elem := range allowanceList {
		elem := elem
		genesis.AllowanceList = append(genesis.AllowanceList, &elem)
	}

	// Set the current count
	genesis.AllowanceCount = k.GetAllowanceCount(ctx)

	// Get all address
	addressList := k.GetAllAddress(ctx)
	for _, elem := range addressList {
		elem := elem
		genesis.AddressList = append(genesis.AddressList, &elem)
	}

	// Set the current count
	genesis.AddressCount = k.GetAddressCount(ctx)

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
*/

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
