package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/konstellation/konstellation/app"
	"github.com/konstellation/konstellation/x/oracle/keeper"
)

func TestNewQuerier(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	coin := sdk.NewCoin("Darc", sdk.NewInt(10))
	oracleKeeper := simapp.GetOracleKeeper()
	oracleKeeper.SetExchangeRate(ctx, coin)

	query := abcitypes.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	legacyQuerierCdc := codec.NewAminoCodec(simapp.LegacyAmino())
	querier := keeper.NewQuerier(oracleKeeper, legacyQuerierCdc.LegacyAmino)
	bz, err := querier(ctx, []string{"exchange-rate"}, query)
	require.Nil(t, err)
	expected :=
		`{
  "denom": "Darc",
  "amount": "10"
}`
	require.Equal(t, []byte(expected), bz)
}
