package query

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/konstellation/x/issue/keeper"
)

func Freeze(ctx sdk.Context, k keeper.Keeper, denom string, holder string) ([]byte, error) {
	holderAddress, _ := sdk.AccAddressFromBech32(holder)
	freeze := k.GetFreeze(ctx, denom, holderAddress)

	bz, err := k.GetCodec().MarshalBinaryBare(freeze)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, sdkerrors.ErrJSONMarshal.Error())
	}
	return bz, nil
}
