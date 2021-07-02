package query

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/konstellation/x/issue/keeper"
)

func Freeze(ctx sdk.Context, k keeper.Keeper, denom string, holder string) ([]byte, *sdkerrors.Error) {
	holderAddress, _ := sdk.AccAddressFromBech32(holder)
	freeze := k.GetFreeze(ctx, denom, holderAddress)

	bz, err := codec.MarshalJSONIndent(k.GetCodec(), freeze)
	if err != nil {
		return nil, sdkerrors.ErrJSONMarshal
	}
	return bz, nil
}
