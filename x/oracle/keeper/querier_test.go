package keeper_test

import (
	"testing"
)

func TestNewQuerier(t *testing.T) {
	//simapp := app.Setup(false)
	//ctx := simapp.NewContext(true, tmproto.Header{})
	//
	//abc := types.NewAdminAddr("darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx")
	//
	//allowedAddresses := []types.AdminAddr{abc}
	//simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)
	//
	//rate := types.ExchangeRate{
	//	Pair: "udarc",
	//	Rate:  uint64(1.2 * float64(1000000000000000000)),
	//}
	//oracleKeeper := simapp.GetOracleKeeper()
	//oracleKeeper.SetExchangeRate(ctx, abc.GetAdminAddress(), &rate)

	//	query := abcitypes.RequestQuery{
	//		Path: "",
	//		Data: []byte{},
	//	}
	//
	//	legacyQuerierCdc := codec.NewAminoCodec(simapp.LegacyAmino())
	//	querier := keeper.NewQuerier(oracleKeeper, legacyQuerierCdc.LegacyAmino)
	//	bz, err := querier(ctx, []string{"exchange_rate", "kbtckusd"}, query)
	//	require.Nil(t, err)
	//	expected :=
	//		`{
	//  "denom": "udarc",
	//  "rate": "1200000000000000000"
	//}`
	//	require.Equal(t, []byte(expected), bz)
}
