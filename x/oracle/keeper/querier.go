package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/konstellation/konstellation/x/oracle/types"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {

		case types.QueryExchangeRate:
			return queryExchangeRate(ctx, k, path[1], legacyQuerierCdc)

		case types.QueryAllExchangeRates:
			return queryAllExchangeRates(ctx, k, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown oracle query endpoint")
		}
	}
}

func queryExchangeRate(ctx sdk.Context, keeper Keeper, pair string, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	exchangeRate, _ := keeper.GetExchangeRate(ctx, pair)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, exchangeRate)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryAllExchangeRates(ctx sdk.Context, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	exchangeRates := keeper.GetAllExchangeRates(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, exchangeRates)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
