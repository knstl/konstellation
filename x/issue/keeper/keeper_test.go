package keeper_test

import (
	types2 "github.com/konstellation/konstellation/types"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Write code here to run before tests
	types2.RegisterBech32Prefix()

	// Run tests
	exitVal := m.Run()

	// Write code here to run after tests

	// Exit with exit value from tests
	os.Exit(exitVal)
}
//
//func TestSetAllowedAddresses(t *testing.T) {
//	simapp := app.Setup(false)
//	ctx := simapp.NewContext(true, tmproto.Header{})
//
//	abc := types.NewAdminAddr("darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx")
//	def := types.NewAdminAddr("darc18c7c95p5cqqxlsg4m56626djna6eyqq4xkvtu3")
//
//	allowedAddresses := []types.AdminAddr{*abc}
//	newAllowedAddresses := []types.AdminAddr{*def}
//	updatedAllowedAddresses := []types.AdminAddr{*abc, *def}
//
//	k := simapp.GetOracleKeeper()
//	k.SetTestAllowedAddresses(ctx, allowedAddresses)
//	k.SetAllowedAddresses(ctx, allowedAddresses[0].GetAdminAddress(), newAllowedAddresses)
//
//	require.Equal(t, updatedAllowedAddresses, simapp.GetOracleKeeper().GetAllowedAddresses(ctx))
//}
//
//func TestDeleteAllowedAddresses(t *testing.T) {
//	simapp := app.Setup(false)
//	ctx := simapp.NewContext(true, tmproto.Header{})
//
//	abc := types.NewAdminAddr("darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx")
//	def := types.NewAdminAddr("darc18c7c95p5cqqxlsg4m56626djna6eyqq4xkvtu3")
//
//	allowedAddresses := []types.AdminAddr{*abc}
//	newAllowedAddresses := []types.AdminAddr{*def}
//
//	k := simapp.GetOracleKeeper()
//	err := k.SetTestAllowedAddresses(ctx, allowedAddresses)
//	require.Equal(t, nil, err)
//	err = k.SetAllowedAddresses(ctx, abc.GetAdminAddress(), newAllowedAddresses)
//	require.Equal(t, nil, err)
//	err = k.DeleteAllowedAddresses(ctx, abc.GetAdminAddress(), allowedAddresses)
//	require.Equal(t, nil, err)
//	err = k.DeleteAllowedAddresses(ctx, abc.GetAdminAddress(), allowedAddresses)
//	require.NotEqual(t, nil, err)
//
//	require.Equal(t, newAllowedAddresses, k.GetAllowedAddresses(ctx))
//}
//
//func TestGetAllowedAddress(t *testing.T) {
//	simapp := app.Setup(false)
//	ctx := simapp.NewContext(true, tmproto.Header{})
//
//	abc := types.NewAdminAddr("darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx")
//
//	allowedAddresses := []types.AdminAddr{*abc}
//
//	k := simapp.GetOracleKeeper()
//	k.SetTestAllowedAddresses(ctx, allowedAddresses)
//	adm, _ := k.GetAllowedAddress(ctx, allowedAddresses[0].GetAdminAddress())
//
//	require.Equal(t, *abc, adm)
//}
//
//func TestIsAllowedAddress(t *testing.T) {
//	simapp := app.Setup(false)
//	ctx := simapp.NewContext(true, tmproto.Header{})
//
//	abc := types.NewAdminAddr("darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx")
//	def := types.NewAdminAddr("darc18c7c95p5cqqxlsg4m56626djna6eyqq4xkvtu3")
//
//	allowedAddresses := []types.AdminAddr{*abc}
//
//	k := simapp.GetOracleKeeper()
//	k.SetTestAllowedAddresses(ctx, allowedAddresses)
//
//	require.Equal(t, true, k.IsAllowedAddress(ctx, allowedAddresses[0].GetAdminAddress()))
//	require.Equal(t, false, k.IsAllowedAddress(ctx, def.GetAdminAddress()))
//}
//
//func TestSetExchangeRate(t *testing.T) {
//	simapp := app.Setup(false)
//	ctx := simapp.NewContext(true, tmproto.Header{})
//
//	abc := types.NewAdminAddr("darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx")
//
//	allowedAddresses := []types.AdminAddr{*abc}
//	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)
//
//	d1, _ := types3.NewDecFromStr("1200000000000")
//	rate := types.ExchangeRate{
//		Pair:   "kbtckusd",
//		Rate:   d1,
//		Denoms: []string{"kbtc", "kusd"},
//	}
//
//	d2, _ := types3.NewDecFromStr("1300000000000")
//	rate2 := types.ExchangeRate{
//		Pair:   "kethkusd",
//		Rate:   d2,
//		Denoms: []string{"keth", "kusd"},
//	}
//	oracleKeeper := simapp.GetOracleKeeper()
//	oracleKeeper.SetExchangeRate(ctx, abc.GetAdminAddress(), &rate)
//	exrate, found := oracleKeeper.GetExchangeRate(ctx, rate.Pair)
//	oracleKeeper.SetExchangeRate(ctx, abc.GetAdminAddress(), &rate2)
//	exrate2, found2 := oracleKeeper.GetExchangeRate(ctx, rate2.Pair)
//
//	require.Equal(t, rate, exrate)
//	require.Equal(t, true, found)
//	require.Equal(t, rate2, exrate2)
//	require.Equal(t, true, found2)
//
//	aer := oracleKeeper.GetAllExchangeRates(ctx)
//	require.Equal(t, []types.ExchangeRate{rate, rate2}, aer)
//}
//
//func TestSetExchangeRates(t *testing.T) {
//	simapp := app.Setup(false)
//	ctx := simapp.NewContext(true, tmproto.Header{})
//
//	abc := types.NewAdminAddr("darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx")
//
//	allowedAddresses := []types.AdminAddr{*abc}
//	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)
//
//	d1, _ := types3.NewDecFromStr("1200000000000")
//	rate1 := &types.ExchangeRate{
//		Pair:   "kbtckusd",
//		Rate:   d1,
//		Denoms: []string{"kbtc", "kusd"},
//	}
//
//	d2, _ := types3.NewDecFromStr("1300000000000")
//	rate2 := &types.ExchangeRate{
//		Pair:   "kethkusd",
//		Rate:   d2,
//		Denoms: []string{"keth", "kusd"},
//	}
//
//	rates := []*types.ExchangeRate{rate1, rate2}
//	oracleKeeper := simapp.GetOracleKeeper()
//	err := oracleKeeper.SetExchangeRates(ctx, abc.GetAdminAddress(), rates)
//	require.Equal(t, nil, err)
//	exrate, found := oracleKeeper.GetExchangeRate(ctx, rate1.Pair)
//	exrate2, found2 := oracleKeeper.GetExchangeRate(ctx, rate2.Pair)
//
//	require.Equal(t, *rate1, exrate)
//	require.Equal(t, true, found)
//	require.Equal(t, *rate2, exrate2)
//	require.Equal(t, true, found2)
//
//	aer := oracleKeeper.GetAllExchangeRates(ctx)
//	require.Equal(t, []types.ExchangeRate{*rate1, *rate2}, aer)
//}
//
//func TestDeleteExchangeRate(t *testing.T) {
//	simapp := app.Setup(false)
//	ctx := simapp.NewContext(true, tmproto.Header{})
//
//	abc := types.NewAdminAddr("darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx")
//
//	allowedAddresses := []types.AdminAddr{*abc}
//	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)
//
//	d1, _ := types3.NewDecFromStr("1200000000000")
//	rate1 := types.ExchangeRate{
//		Pair:   "kbtckusd",
//		Rate:   d1,
//		Denoms: []string{"kbtc", "kusd"},
//	}
//
//	d2, _ := types3.NewDecFromStr("1300000000000")
//	rate2 := types.ExchangeRate{
//		Pair:   "kethkusd",
//		Rate:   d2,
//		Denoms: []string{"keth", "kusd"},
//	}
//
//	oracleKeeper := simapp.GetOracleKeeper()
//	oracleKeeper.SetExchangeRate(ctx, abc.GetAdminAddress(), &rate1)
//	oracleKeeper.SetExchangeRate(ctx, abc.GetAdminAddress(), &rate2)
//
//	aer := oracleKeeper.GetAllExchangeRates(ctx)
//	require.Equal(t, []types.ExchangeRate{rate1, rate2}, aer)
//
//	require.Nil(t, oracleKeeper.DeleteExchangeRate(ctx, abc.GetAdminAddress(), rate1.Pair))
//	aer2 := oracleKeeper.GetAllExchangeRates(ctx)
//	require.Equal(t, []types.ExchangeRate{rate2}, aer2)
//}
//
//func TestDeleteExchangeRates(t *testing.T) {
//	simapp := app.Setup(false)
//	ctx := simapp.NewContext(true, tmproto.Header{})
//
//	abc := types.NewAdminAddr("darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx")
//
//	allowedAddresses := []types.AdminAddr{*abc}
//	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)
//
//	d1, _ := types3.NewDecFromStr("1200000000000")
//	rate1 := &types.ExchangeRate{
//		Pair:   "kbtckusd",
//		Rate:   d1,
//		Denoms: []string{"kbtc", "kusd"},
//	}
//
//	d2, _ := types3.NewDecFromStr("1300000000000")
//	rate2 := &types.ExchangeRate{
//		Pair:   "kethkusd",
//		Rate:   d2,
//		Denoms: []string{"keth", "kusd"},
//	}
//
//	rates := []*types.ExchangeRate{rate1, rate2}
//
//	oracleKeeper := simapp.GetOracleKeeper()
//	err := oracleKeeper.SetExchangeRates(ctx, abc.GetAdminAddress(), rates)
//	require.Equal(t, nil, err)
//
//	aer := oracleKeeper.GetAllExchangeRates(ctx)
//	require.Equal(t, []types.ExchangeRate{*rate1, *rate2}, aer)
//
//	require.Nil(t, oracleKeeper.DeleteExchangeRates(ctx, abc.GetAdminAddress(), []string{rate1.Pair}))
//	aer2 := oracleKeeper.GetAllExchangeRates(ctx)
//	require.Equal(t, []types.ExchangeRate{*rate2}, aer2)
//}
