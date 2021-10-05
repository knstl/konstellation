package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/konstellation/x/issue/types"
)

// NewQuerier creates a querier for auth REST endpoints
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryIssue:
			return issue(ctx, k, path[1])
		case types.QueryIssues:
			return issues(ctx, k, req.Data)
		case types.QueryIssuesAll:
			return issuesAll(ctx, k)
		case types.QueryAllowance:
			return allowance(ctx, k, path[1], path[2], path[3])
		case types.QueryAllowances:
			return allowances(ctx, k, path[1], path[2])
		case types.QueryFreeze:
			return freeze(ctx, k, path[1], path[2])
		case types.QueryFreezes:
			return freezes(ctx, k, path[1])
		case types.QueryParams:
			return params(ctx, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown issue query endpoint")
		}
	}
}

func allowance(ctx sdk.Context, k Keeper, denom string, owner string, spender string) ([]byte, error) {
	//ownerAddress, _ := sdk.AccAddressFromBech32(owner)
	//spenderAddress, _ := sdk.AccAddressFromBech32(spender)
	//amount := k.Allowance(ctx, ownerAddress, spenderAddress, denom)
	req := types.QueryGetAllowanceRequest{
		Denom:   denom,
		Owner:   owner,
		Spender: spender,
	}
	amount, err := k.Allowance(ctx.Context(), &req)
	if err != nil {
		return nil, err
	}

	//if amount.GT(sdk.ZeroInt()) {
	//	issue := k.GetIssue(ctx, issueID)
	//	amount = issue.QuoDecimals(amount)
	//}

	bz, err := k.GetCodec().MarshalBinaryBare(amount)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, sdkerrors.ErrJSONMarshal.Error())
	}
	return bz, nil
}

func allowances(ctx sdk.Context, k Keeper, denom string, owner string) ([]byte, error) {
	ownerAddress, _ := sdk.AccAddressFromBech32(owner)
	allowances := k.Allowances(ctx, ownerAddress, denom)
	allowanceSlice := []*types.Allowance{}
	for _, allowance := range allowances {
		allowanceSlice = append(allowanceSlice, allowance)
	}
	allowanceList := types.AllowanceList{Allowances: allowanceSlice}
	bz, err := k.GetCodec().MarshalBinaryBare(&allowanceList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, sdkerrors.ErrJSONMarshal.Error())
	}
	return bz, nil
}

func freeze(ctx sdk.Context, k Keeper, denom string, holder string) ([]byte, error) {
	holderAddress, err := sdk.AccAddressFromBech32(holder)
	if err != nil {
		return nil, err
	}
	freeze := k.GetFreeze(ctx, denom, holderAddress)

	bz, err := k.GetCodec().MarshalBinaryBare(freeze)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, sdkerrors.ErrJSONMarshal.Error())
	}
	return bz, nil
}

func freezes(ctx sdk.Context, k Keeper, denom string) ([]byte, error) {
	//holderAddress, err := sdk.AccAddressFromBech32(holder)
	//if err != nil {
	//	return nil, err
	//}
	//freezes := k.GetFreezesOfDenom(ctx, denom)
	//freezeList := types.AddressFreezeList{AddressFreezes: freezes}
	//
	//bz, err := k.GetCodec().MarshalBinaryBare(&freezeList)
	//if err != nil {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, sdkerrors.ErrJSONMarshal.Error())
	//}
	//return bz, nil
}

func freezesAll(ctx sdk.Context, k Keeper, denom string) ([]byte, error) {
	freezes := k.GetFreezesOfDenom(ctx, denom)
	freezeList := types.AddressFreezeList{AddressFreezes: freezes}

	bz, err := k.GetCodec().MarshalBinaryBare(freezes)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, sdkerrors.ErrJSONMarshal.Error())
	}
	return bz, nil
}

func issue(ctx sdk.Context, k Keeper, denom string) ([]byte, error) {
	coin, er := k.GetIssue(ctx, denom)
	if er != nil {
		return nil, sdkerrors.Wrap(er, er.Error())
	}

	bz, err := k.GetCodec().MarshalBinaryBare(coin)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, sdkerrors.ErrJSONMarshal.Error())
	}
	return bz, nil
}

func issuesAll(ctx sdk.Context, k Keeper) ([]byte, error) {
	issueList := types.CoinIssueList{CoinIssues: []*types.CoinIssue{}}
	issues := k.ListAll(ctx)
	for _, anIssue := range issues {
		issueList.CoinIssues = append(issueList.CoinIssues, anIssue)
	}
	bz, err := k.GetCodec().MarshalBinaryBare(&issueList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, sdkerrors.ErrJSONMarshal.Error())
	}
	return bz, nil
}

func issues(ctx sdk.Context, k Keeper, data []byte) ([]byte, error) {
	var params types.IssuesParams
	if err := k.GetCodec().UnmarshalBinaryBare(data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, sdkerrors.ErrJSONUnmarshal.Error())
	}

	issueList := types.CoinIssueList{CoinIssues: []*types.CoinIssue{}}
	issues := k.List(ctx, params)
	for _, anIssue := range issues {
		issueList.CoinIssues = append(issueList.CoinIssues, anIssue)
	}
	bz, err := k.GetCodec().MarshalBinaryBare(&issueList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, sdkerrors.ErrJSONMarshal.Error())
	}

	return bz, nil
}

func params(ctx sdk.Context, k Keeper) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := k.GetCodec().MarshalBinaryBare(&params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, sdkerrors.ErrJSONMarshal.Error())
	}

	return res, nil
}
