package keeper_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/konstellation/konstellation/app"
	"github.com/konstellation/konstellation/x/oracle/keeper"
)

func TestNewQuerier(t *testing.T) {
	app := app.Setup(false)
	ctx := app.NewContext(true, tmproto.Header{})
	coin := sdk.NewCoin("Darc", sdk.NewInt(10))
	oracleKeeper := app.GetOracleKeeper()
	oracleKeeper.SetExchangeRate(ctx, coin)
	fmt.Printf("***** %+v\n", oracleKeeper)

	query := abci.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	legacyQuerierCdc := codec.NewAminoCodec(app.LegacyAmino())
	querier := keeper.NewQuerier(oracleKeeper, legacyQuerierCdc.LegacyAmino)
	bz, err := querier(ctx, []string{"exchange-rate"}, query)
	require.Error(t, err)
	require.Nil(t, bz)
}
