package query

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/konstellation/x/issue/keeper"
)

func Allowance(ctx sdk.Context, k keeper.Keeper, denom string, owner string, spender string) ([]byte, error) {
	ownerAddress, _ := sdk.AccAddressFromBech32(owner)
	spenderAddress, _ := sdk.AccAddressFromBech32(spender)
	amount := k.Allowance(ctx, ownerAddress, spenderAddress, denom)

	//if amount.GT(sdk.ZeroInt()) {
	//	issue := k.GetIssue(ctx, issueID)
	//	amount = issue.QuoDecimals(amount)
	//}

	bz, err := k.GetCodec().MarshalBinaryBare(&amount)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, sdkerrors.ErrJSONMarshal.Error())
	}
	return bz, nil
}
