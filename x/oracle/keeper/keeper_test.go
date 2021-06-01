package keeper_test

import (
	types2 "github.com/konstellation/konstellation/types"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/konstellation/konstellation/app"
	"github.com/konstellation/konstellation/x/oracle/types"
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

func TestSetAllowedAddresses(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	abc := types.NewAdminAddr("darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx")
	def := types.NewAdminAddr("darc18c7c95p5cqqxlsg4m56626djna6eyqq4xkvtu3")

	allowedAddresses := []types.AdminAddr{abc}
	newAllowedAddresses := []types.AdminAddr{def}
	updatedAllowedAddresses := []types.AdminAddr{abc, def}

	k := simapp.GetOracleKeeper()
	k.SetTestAllowedAddresses(ctx, allowedAddresses)
	k.SetAllowedAddresses(ctx, allowedAddresses[0].GetAdminAddress(), newAllowedAddresses)

	require.Equal(t, updatedAllowedAddresses, simapp.GetOracleKeeper().GetAllowedAddresses(ctx))
}

func TestDeleteAllowedAddresses(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	abc := types.NewAdminAddr("darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx")
	def := types.NewAdminAddr("darc18c7c95p5cqqxlsg4m56626djna6eyqq4xkvtu3")

	allowedAddresses := []types.AdminAddr{abc}
	newAllowedAddresses := []types.AdminAddr{def}

	k := simapp.GetOracleKeeper()
	err := k.SetTestAllowedAddresses(ctx, allowedAddresses)
	require.Equal(t, nil, err)
	err = k.SetAllowedAddresses(ctx, abc.GetAdminAddress(), newAllowedAddresses)
	require.Equal(t, nil, err)
	err = k.DeleteAllowedAddresses(ctx, abc.GetAdminAddress(), allowedAddresses)
	require.Equal(t, nil, err)
	err = k.DeleteAllowedAddresses(ctx, abc.GetAdminAddress(), allowedAddresses)
	require.NotEqual(t, nil, err)

	require.Equal(t, newAllowedAddresses, k.GetAllowedAddresses(ctx))
}

func TestGetAllowedAddress(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	abc := types.NewAdminAddr("darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx")

	allowedAddresses := []types.AdminAddr{abc}

	k := simapp.GetOracleKeeper()
	k.SetTestAllowedAddresses(ctx, allowedAddresses)
	adm, _ := k.GetAllowedAddress(ctx, allowedAddresses[0].GetAdminAddress())

	require.Equal(t, abc, adm)
}

func TestIsAllowedAddress(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	abc := types.NewAdminAddr("darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx")
	def := types.NewAdminAddr("darc18c7c95p5cqqxlsg4m56626djna6eyqq4xkvtu3")

	allowedAddresses := []types.AdminAddr{abc}

	k := simapp.GetOracleKeeper()
	k.SetTestAllowedAddresses(ctx, allowedAddresses)

	require.Equal(t, true, k.IsAllowedAddress(ctx, allowedAddresses[0].GetAdminAddress()))
	require.Equal(t, false, k.IsAllowedAddress(ctx, def.GetAdminAddress()))
}

func TestSetExchangeRate(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	abc := types.NewAdminAddr("darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx")

	allowedAddresses := []types.AdminAddr{abc}
	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)

	rate := types.ExchangeRate{
		Pair:      "kbtckusd",
		Rate:      uint64(1.2 * float64(1000000000000000000)),
		Denoms:    []string{"kbtc", "kusd"},
		Timestamp: time.Now().Unix(),
	}
	rate2 := types.ExchangeRate{
		Pair:      "kethkusd",
		Rate:      uint64(1.4 * float64(1000000000000000000)),
		Denoms:    []string{"keth", "kusd"},
		Timestamp: time.Now().Unix(),
	}
	oracleKeeper := simapp.GetOracleKeeper()
	oracleKeeper.SetExchangeRate(ctx, abc.GetAdminAddress(), &rate)
	exrate, found := oracleKeeper.GetExchangeRate(ctx, rate.Pair)
	oracleKeeper.SetExchangeRate(ctx, abc.GetAdminAddress(), &rate2)
	exrate2, found2 := oracleKeeper.GetExchangeRate(ctx, rate2.Pair)

	require.Equal(t, rate, exrate)
	require.Equal(t, true, found)
	require.Equal(t, rate2, exrate2)
	require.Equal(t, true, found2)

	aer := oracleKeeper.GetAllExchangeRates(ctx)
	require.Equal(t, []types.ExchangeRate{rate, rate2}, aer)
}

func TestDeleteExchangeRate(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	abc := types.NewAdminAddr("darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx")

	allowedAddresses := []types.AdminAddr{abc}
	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)

	rate1 := types.ExchangeRate{
		Pair:      "kbtckusd",
		Rate:      uint64(1.2 * float64(1000000000000000000)),
		Denoms:    []string{"kbtc", "kusd"},
		Timestamp: time.Now().Unix(),
	}
	rate2 := types.ExchangeRate{
		Pair:      "kethkusd",
		Rate:      uint64(1.3 * float64(1000000000000000000)),
		Denoms:    []string{"keth", "kusd"},
		Timestamp: time.Now().Unix(),
	}

	oracleKeeper := simapp.GetOracleKeeper()
	oracleKeeper.SetExchangeRate(ctx, abc.GetAdminAddress(), &rate1)
	oracleKeeper.SetExchangeRate(ctx, abc.GetAdminAddress(), &rate2)

	aer := oracleKeeper.GetAllExchangeRates(ctx)
	require.Equal(t, []types.ExchangeRate{rate1, rate2}, aer)

	require.Nil(t, oracleKeeper.DeleteExchangeRate(ctx, abc.GetAdminAddress(), rate1.Pair))
	aer2 := oracleKeeper.GetAllExchangeRates(ctx)
	require.Equal(t, []types.ExchangeRate{rate2}, aer2)
}
