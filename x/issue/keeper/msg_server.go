package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	keeper Keeper
}

func NewMsgServerImpl(k Keeper) types.MsgServer {
	return &msgServer{keeper: k}
}

func (m msgServer) HandleMsgIssueCreate(goCtx context.Context, msgIssueCreate *types.MsgIssueCreate) (*types.MsgIssueCreateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	issue := m.keeper.CreateIssue(ctx, msgIssueCreate.Owner, msgIssueCreate.Issuer, msgIssueCreate.IssueParams)
	return &types.MsgIssueCreateResponse{Amount: issue}, nil
}

func (m msgServer) HandleMsgFeatures(goCtx context.Context, msgFeatures *types.MsgFeatures) (*types.MsgFeaturesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := m.keeper.ChangeFeatures(ctx, msgFeatures.Owner, msgFeatures.Denom, msgFeatures.IssueFeatures)
	if err != nil {
		return nil, err
	}
	return &types.MsgFeaturesResponse{}, nil
}

func (m msgServer) HandleMsgDescription(goCtx context.Context, msgDescription *types.MsgDescription) (*types.MsgDescriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := m.keeper.ChangeDescription(ctx, msgDescription.Owner, msgDescription.Denom, msgDescription.Description)
	if err != nil {
		return nil, err
	}
	return &types.MsgDescriptionResponse{}, nil
}

func (m msgServer) HandleMsgTransfer(goCtx context.Context, msgTransfer *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := m.keeper.Transfer(ctx, msgTransfer.FromAddress, msgTransfer.ToAddress, msgTransfer.Amount)
	if err != nil {
		return nil, err
	}
	return &types.MsgTransferResponse{}, nil
}

func (m msgServer) HandleMsgTransferFrom(goCtx context.Context, msgTransferFrom *types.MsgTransferFrom) (*types.MsgTransferFromResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := m.keeper.TransferFrom(ctx, msgTransferFrom.Sender, msgTransferFrom.FromAddress, msgTransferFrom.ToAddress, msgTransferFrom.Amount)
	if err != nil {
		return nil, err
	}
	return &types.MsgTransferFromResponse{}, nil
}

func (m msgServer) HandleMsgApprove(goCtx context.Context, msgApprove *types.MsgApprove) (*types.MsgApproveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := m.keeper.Approve(ctx, msgApprove.Owner, msgApprove.Spender, msgApprove.Amount)
	if err != nil {
		return nil, err
	}
	return &types.MsgApproveResponse{}, nil
}

func (m msgServer) HandleMsgIncreaseAllowance(goCtx context.Context, msgIncreaseAllowance *types.MsgIncreaseAllowance) (*types.MsgIncreaseAllowanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := m.keeper.IncreaseAllowance(ctx, msgIncreaseAllowance.Owner, msgIncreaseAllowance.Spender, msgIncreaseAllowance.Amount)
	if err != nil {
		return nil, err
	}
	return &types.MsgIncreaseAllowanceResponse{}, nil
}

func (m msgServer) HandleMsgDecreaseAllowance(goCtx context.Context, msgDecreaseAllowance *types.MsgDecreaseAllowance) (*types.MsgDecreaseAllowanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := m.keeper.DecreaseAllowance(ctx, msgDecreaseAllowance.Owner, msgDecreaseAllowance.Spender, msgDecreaseAllowance.Amount)
	if err != nil {
		return nil, err
	}
	return &types.MsgDecreaseAllowanceResponse{}, nil
}

func (m msgServer) HandleMsgMint(goCtx context.Context, msgMint *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := m.keeper.Mint(ctx, msgMint.Minter, msgMint.ToAddress, msgMint.Amount)
	if err != nil {
		return nil, err
	}
	return &types.MsgMintResponse{}, nil
}

func (m msgServer) HandleMsgBurn(goCtx context.Context, msgBurn *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := m.keeper.Burn(ctx, msgBurn.Burner, msgBurn.Amount)
	if err != nil {
		return nil, err
	}
	return &types.MsgBurnResponse{}, nil
}

func (m msgServer) HandleMsgBurnFrom(goCtx context.Context, msgBurnFrom *types.MsgBurnFrom) (*types.MsgBurnFromResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := m.keeper.BurnFrom(ctx, msgBurnFrom.Burner, msgBurnFrom.FromAddress, msgBurnFrom.Amount)
	if err != nil {
		return nil, err
	}
	return &types.MsgBurnFromResponse{}, nil
}

func (m msgServer) HandleMsgTransferOwnership(goCtx context.Context, msgTransferOwnership *types.MsgTransferOwnership) (*types.MsgTransferOwnershipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := m.keeper.TransferOwnership(ctx, msgTransferOwnership.Owner, msgTransferOwnership.ToAddress, msgTransferOwnership.Denom)

	if err != nil {
		return nil, err
	}
	return &types.MsgTransferOwnershipResponse{}, nil
}

func (m msgServer) HandleMsgFreeze(goCtx context.Context, msgFreeze *types.MsgFreeze) (*types.MsgFreezeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := m.keeper.Freeze(ctx, msgFreeze.Freezer, msgFreeze.Holder, msgFreeze.Denom, msgFreeze.Op)
	if err != nil {
		return nil, err
	}
	return &types.MsgFreezeResponse{}, nil
}

func (m msgServer) HandleMsgUnfreeze(goCtx context.Context, msgUnfreeze *types.MsgUnfreeze) (*types.MsgUnfreezeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := m.keeper.Unfreeze(ctx, msgUnfreeze.Freezer, msgUnfreeze.Holder, msgUnfreeze.Denom, msgUnfreeze.Op)
	if err != nil {
		return nil, err
	}
	return &types.MsgUnfreezeResponse{}, nil
}
