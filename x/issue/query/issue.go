package query

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
)

func Issue(ctx sdk.Context, k keeper.Keeper, denom string) ([]byte, sdk.Error) {
	coin, er := k.GetIssue(ctx, denom)
	if er != nil {
		return nil, er
	}

	bz, err := codec.MarshalJSONIndent(k.GetCodec(), coin)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
