package query

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/keeper"
	"github.com/konstellation/konstellation/x/issue/types"
)

func Allowance(ctx sdk.Context, k keeper.Keeper, denom string, owner string, spender string) ([]byte, sdk.Error) {
	ownerAddress, _ := sdk.AccAddressFromBech32(owner)
	spenderAddress, _ := sdk.AccAddressFromBech32(spender)
	amount := k.Allowance(ctx, ownerAddress, spenderAddress, denom)

	//if amount.GT(sdk.ZeroInt()) {
	//	issue := k.GetIssue(ctx, issueID)
	//	amount = issue.QuoDecimals(amount)
	//}

	bz, err := codec.MarshalJSONIndent(k.GetCodec(), types.NewAllowance(amount))
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
