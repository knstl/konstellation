package query

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/konstellation/x/issue/keeper"
)

func Issue(ctx sdk.Context, k keeper.Keeper, denom string) ([]byte, *sdkerrors.Error) {
	coin, er := k.GetIssue(ctx, denom)
	if er != nil {
		return nil, er
	}

	bz, err := codec.MarshalJSONIndent(k.GetCodec(), coin)
	if err != nil {
		return nil, sdkerrors.ErrJSONMarshal
	}
	return bz, nil
}
