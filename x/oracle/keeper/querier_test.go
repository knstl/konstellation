package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stretchr/testify/require"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/konstellation/konstellation/app"
	"github.com/konstellation/konstellation/x/oracle/keeper"
	"github.com/konstellation/konstellation/x/oracle/types"
)

func TestNewQuerier(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})
	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, []string{"abc"})

	rate := types.ExchangeRate{
		Pair: "udarc",
		Rate:  uint64(1.2 * float64(1000000000000000000)),
	}
	oracleKeeper := simapp.GetOracleKeeper()
	oracleKeeper.SetExchangeRate(ctx, "abc", &rate)

	query := abcitypes.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	legacyQuerierCdc := codec.NewAminoCodec(simapp.LegacyAmino())
	querier := keeper.NewQuerier(oracleKeeper, legacyQuerierCdc.LegacyAmino)
	bz, err := querier(ctx, []string{"exchange_rate", "kbtckusd"}, query)
	require.Nil(t, err)
	expected :=
		`{
  "denom": "udarc",
  "rate": "1200000000000000000"
}`
	require.Equal(t, []byte(expected), bz)
}
