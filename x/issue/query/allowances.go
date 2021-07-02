package query

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/konstellation/x/issue/keeper"
)

func Allowances(ctx sdk.Context, k keeper.Keeper, denom string, owner string) ([]byte, *sdkerrors.Error) {
	ownerAddress, _ := sdk.AccAddressFromBech32(owner)
	allowances := k.Allowances(ctx, ownerAddress, denom)

	bz, err := codec.MarshalJSONIndent(k.GetCodec(), allowances)
	if err != nil {
		return nil, sdkerrors.ErrJSONMarshal
	}
	return bz, nil
}
