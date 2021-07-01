package query

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
)

func Allowances(ctx sdk.Context, k keeper.Keeper, denom string, owner string) ([]byte, sdk.Error) {
	ownerAddress, _ := sdk.AccAddressFromBech32(owner)
	allowances := k.Allowances(ctx, ownerAddress, denom)

	bz, err := codec.MarshalJSONIndent(k.GetCodec(), allowances)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
