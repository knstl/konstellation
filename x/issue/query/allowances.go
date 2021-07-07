package query

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/konstellation/x/issue/keeper"
	"github.com/konstellation/konstellation/x/issue/types"
)

func Allowances(ctx sdk.Context, k keeper.Keeper, denom string, owner string) ([]byte, error) {
	ownerAddress, _ := sdk.AccAddressFromBech32(owner)
	allowances := k.Allowances(ctx, ownerAddress, denom)
	allowanceSlice := []*types.Allowance{}
	for _, allowance := range allowances {
		allowanceSlice = append(allowanceSlice, allowance)
	}
	allowanceList := types.AllowanceList{Allowances: allowanceSlice}
	bz, err := k.GetCodec().MarshalBinaryBare(&allowanceList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, sdkerrors.ErrJSONMarshal.Error())
	}
	return bz, nil
}
