package query

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/konstellation/x/issue/keeper"
)

func IssuesAll(ctx sdk.Context, k keeper.Keeper) ([]byte, *sdkerrors.Error) {
	issues := k.ListAll(ctx)
	bz, err := codec.MarshalJSONIndent(k.GetCodec(), issues)
	if err != nil {
		return nil, sdkerrors.ErrJSONMarshal
	}
	return bz, nil
}
