package query

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
)

func Freezes(ctx sdk.Context, k keeper.Keeper, denom string) ([]byte, sdk.Error) {
	freezes := k.GetFreezes(ctx, denom)

	bz, err := codec.MarshalJSONIndent(k.GetCodec(), freezes)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
