package query

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/konstellation/x/issue/keeper"
)

func Params(ctx sdk.Context, k keeper.Keeper) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := k.GetCodec().MarshalBinaryBare(&params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, sdkerrors.ErrJSONMarshal.Error())
	}

	return res, nil
}
