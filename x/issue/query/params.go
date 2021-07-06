package query

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/konstellation/x/issue/keeper"
)

func Params(ctx sdk.Context, k keeper.Keeper) ([]byte, *sdkerrors.Error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(k.GetCodec(), params)
	if err != nil {
		return nil, sdkerrors.ErrJSONMarshal
	}

	return res, nil
}
