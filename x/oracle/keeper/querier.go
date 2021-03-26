package keeper

import (
	"github.com/konstellation/konstellation/x/oracle/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier creates a new querier for nameservice clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		// this line is used by starport scaffolding # 2

		case types.QueryExchangeRate:
			return queryExchangeRate(ctx, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown oracle query endpoint")
		}
	}
}

// queryExchangeRate - returns the exchange rate
func queryExchangeRate(ctx sdk.Context, keeper Keeper) ([]byte, error) {
	exchangeRate := keeper.GetExchangeRate(ctx)

	res, err := codec.MarshalJSONIndent(keeper.cdc, exchangeRate)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
