package query

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/konstellation/x/issue/keeper"
)

func Freezes(ctx sdk.Context, k keeper.Keeper, denom string) ([]byte, *sdkerrors.Error) {
	freezes := k.GetFreezes(ctx, denom)

	bz, err := codec.MarshalJSONIndent(k.GetCodec(), freezes)
	if err != nil {
		return nil, sdkerrors.ErrJSONMarshal
	}
	return bz, nil
}
