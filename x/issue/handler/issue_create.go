package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/konstellation/konstellation/x/issue/keeper"
	"github.com/konstellation/konstellation/x/issue/types"
)

func HandleMsgIssueCreate(ctx sdk.Context, k keeper.Keeper, msg *types.MsgIssueCreate) *sdk.Result {
	// Sub fee from issuer
	fee := k.GetParams(ctx).IssueFee
	if err := k.ChargeFee(ctx, sdk.AccAddress(msg.Issuer), fee); err != nil {
		return &sdk.Result{Log: err.Error()}
	}

	params, err := types.NewIssueParams(msg.IssueParams)
	if err != nil {
		return &sdk.Result{Log: types.ErrInvalidIssueParams.Error()}
	}

	ci := k.CreateIssue(ctx, sdk.AccAddress(msg.Owner), sdk.AccAddress(msg.Issuer), params)
	if err := k.Issue(ctx, ci); err != nil {
		return &sdk.Result{Log: err.Error()}
	}

	events := []abcitypes.Event{}
	for _, event := range ctx.EventManager().Events() {
		evt := abcitypes.Event(event)
		events = append(events, evt)
	}
	return &sdk.Result{Events: events}
}
