package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/konstellation/konstellation/x/issue/types"
)

func (k *Keeper) getIssue(ctx sdk.Context, denom string) *types.CoinIssue {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(GetDenomKey(denom))
	if len(bz) == 0 {
		return nil
	}

	var coinIssue types.CoinIssue
	k.GetCodec().MustUnmarshalBinaryLengthPrefixed(bz, &coinIssue)

	return &coinIssue
}

func (k *Keeper) GetIssue(ctx sdk.Context, denom string) (*types.CoinIssue, *sdkerrors.Error) {
	issue := k.getIssue(ctx, denom)
	if issue == nil {
		return nil, types.ErrUnknownIssue(denom)
	}

	return issue, nil
}

func (k *Keeper) Issue(ctx sdk.Context, issue *types.CoinIssue) *sdkerrors.Error {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeIssue,
			sdk.NewAttribute(sdk.AttributeKeyAmount, issue.ToCoin().String()),
			sdk.NewAttribute(types.AttributeKeyIssuer, issue.GetIssuer()),
		),
	)

	i := k.getIssue(ctx, issue.Denom)
	if i != nil {
		return types.ErrIssueAlreadyExists
	}

	k.addIssue(ctx, issue)

	if err := k.csk.MintCoins(ctx, types.ModuleName, issue.ToCoins()); err != nil {
		return types.ErrCanNotMint(issue.Denom)
	}
	owner, _ := sdk.AccAddressFromBech32(issue.GetOwner())
	if err := k.csk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, issue.ToCoins()); err != nil {
		return types.ErrCanNotTransferIn(issue.Denom, owner.String())
	}

	return nil
}
