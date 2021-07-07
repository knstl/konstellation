package query

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/konstellation/x/issue/keeper"
	"github.com/konstellation/konstellation/x/issue/types"
)

func Freezes(ctx sdk.Context, k keeper.Keeper, denom string) ([]byte, error) {
	freezes := k.GetFreezes(ctx, denom)
	freezeSlice := []*types.AddressFreeze{}
	for _, freeze := range freezes {
		freezeSlice = append(freezeSlice, freeze)
	}
	freezeList := types.AddressFreezeList{AddressFreezes: freezeSlice}

	bz, err := k.GetCodec().MarshalBinaryBare(&freezeList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, sdkerrors.ErrJSONMarshal.Error())
	}
	return bz, nil
}
