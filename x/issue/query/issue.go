package query

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/konstellation/x/issue/keeper"
)

func Issue(ctx sdk.Context, k keeper.Keeper, denom string) ([]byte, error) {
	coin, er := k.GetIssue(ctx, denom)
	if er != nil {
		return nil, sdkerrors.Wrap(er, er.Error())
	}

	bz, err := k.GetCodec().MarshalBinaryBare(coin)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, sdkerrors.ErrJSONMarshal.Error())
	}
	return bz, nil
}
