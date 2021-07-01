package query

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/kn-sdk/x/issue/keeper"
)

func Freeze(ctx sdk.Context, k keeper.Keeper, denom string, holder string) ([]byte, sdk.Error) {
	holderAddress, _ := sdk.AccAddressFromBech32(holder)
	freeze := k.GetFreeze(ctx, denom, holderAddress)

	bz, err := codec.MarshalJSONIndent(k.GetCodec(), freeze)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
