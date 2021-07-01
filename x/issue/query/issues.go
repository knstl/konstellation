package query

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
	"github.com/konstellation/kn-sdk/x/issue/types"
)

func Issues(ctx sdk.Context, k keeper.Keeper, data []byte) ([]byte, sdk.Error) {
	var params types.IssuesParams
	if err := k.GetCodec().UnmarshalJSON(data, &params); err != nil {
		return nil, sdk.ErrUnknownRequest(sdk.AppendMsgToErr("incorrectly formatted request data", err.Error()))
	}

	issues := k.List(ctx, params)
	bz, err := codec.MarshalJSONIndent(k.GetCodec(), issues)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}

	return bz, nil
}
