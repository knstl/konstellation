package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) QueryAllowance(c context.Context, r *types.QueryAllowanceRequest) (*types.QueryAllowanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	allowance := k.Allowance(ctx, r.Owner, r.Spender, r.Denom)

	return &types.QueryAllowanceResponse{Allowance: allowance}, nil
}

func (k Keeper) QueryAllowances(c context.Context, r *types.QueryAllowancesRequest) (*types.QueryAllowancesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	allowances := k.Allowances(ctx, r.Owner, r.Denom)

	allowanceList := []*types.Allowance{}
	for _, allowance := range allowances {
		allowanceList = append(allowanceList, allowance)
	}

	return &types.QueryAllowancesResponse{Allowances: allowanceList}, nil
}

func (k Keeper) QueryFreeze(c context.Context, r *types.QueryFreezeRequest) (*types.QueryFreezeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	freeze := k.GetFreeze(ctx, r.Denom, r.Holder)

	return &types.QueryFreezeResponse{Freeze: freeze}, nil
}

func (k Keeper) QueryFreezes(c context.Context, r *types.QueryFreezesRequest) (*types.QueryFreezesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	freezes := k.GetFreezes(ctx, r.Denom)

	freezeList := []*types.AddressFreeze{}
	for _, freeze := range freezes {
		freezeList = append(freezeList, freeze)
	}

	return &types.QueryFreezesResponse{Freezes: freezeList}, nil
}

func (k Keeper) QueryIssue(c context.Context, r *types.QueryIssueRequest) (*types.QueryIssueResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	issue, err := k.GetIssue(ctx, r.Denom)
	if err != nil {
		return nil, err
	}

	return &types.QueryIssueResponse{Issue: issue}, nil
}

func (k Keeper) QueryAllIssues(c context.Context, r *types.QueryAllIssuesRequest) (*types.QueryAllIssuesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	issues := k.ListAll(ctx)

	issueList := []*types.CoinIssue{}
	for _, issue := range issues {
		issueList = append(issueList, issue)
	}

	return &types.QueryAllIssuesResponse{Issues: issueList}, nil
}

func (k Keeper) QueryIssues(c context.Context, r *types.QueryIssuesRequest) (*types.QueryIssuesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	issues := k.List(ctx, *r.Params)

	issueList := []*types.CoinIssue{}
	for _, issue := range issues {
		issueList = append(issueList, issue)
	}

	return &types.QueryIssuesResponse{Issues: issueList}, nil
}

func (k Keeper) QueryParams(c context.Context, r *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: &params}, nil
}
