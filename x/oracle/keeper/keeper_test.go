package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/konstellation/konstellation/app"
	"github.com/konstellation/konstellation/x/oracle/types"
)

func TestLogger(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	require.Equal(t, ctx.Logger(), simapp.GetOracleKeeper().Logger(ctx))
}

func TestSetAllowedAddresses(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	allowedAddresses := []string{"abc"}
	newAllowedAddresses := []string{"def"}
	updatedAllowedAddresses := []string{"abc", "def"}
	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)
	simapp.GetOracleKeeper().SetAllowedAddresses(ctx, allowedAddresses[0], newAllowedAddresses)

	require.Equal(t, updatedAllowedAddresses, simapp.GetOracleKeeper().GetAllowedAddresses(ctx))
}

func TestSetAllowedAddressesIncludeDuplicateList(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	allowedAddresses := []string{"abc"}
	newAllowedAddresses := []string{"abc", "def"}
	updatedAllowedAddresses := []string{"abc", "def"}
	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)
	simapp.GetOracleKeeper().SetAllowedAddresses(ctx, allowedAddresses[0], newAllowedAddresses)

	require.Equal(t, updatedAllowedAddresses, simapp.GetOracleKeeper().GetAllowedAddresses(ctx))
}

func TestSetAllowedAddressesFailure(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	allowedAddresses := []string{"abc"}
	newAllowedAddresses := []string{"def"}
	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)
	require.Error(t, simapp.GetOracleKeeper().SetAllowedAddresses(ctx, newAllowedAddresses[0], newAllowedAddresses))
}

func TestSetAllowedAddressesFailureIncludeDuplicateList(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	allowedAddresses := []string{"abc"}
	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)
	require.Error(t, simapp.GetOracleKeeper().SetAllowedAddresses(ctx, allowedAddresses[0], allowedAddresses))
}

func TestDeleteAllowedAddresses(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	allowedAddresses := []string{"abc"}
	newAllowedAddresses := []string{"def"}
	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)
	simapp.GetOracleKeeper().SetAllowedAddresses(ctx, allowedAddresses[0], newAllowedAddresses)
	simapp.GetOracleKeeper().DeleteAllowedAddresses(ctx, allowedAddresses[0], allowedAddresses)

	require.Equal(t, newAllowedAddresses, simapp.GetOracleKeeper().GetAllowedAddresses(ctx))
}

func TestDeleteAllowedAddressesFailure(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	allowedAddresses := []string{"abc"}
	newAllowedAddresses := []string{"def"}
	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)
	require.Error(t, simapp.GetOracleKeeper().DeleteAllowedAddresses(ctx, newAllowedAddresses[0], allowedAddresses))
}

func TestSetAdminAddrSetAllowedAddresses(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	allowedAddresses := []string{"abc"}
	newAllowedAddresses := []string{"def"}
	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)
	simapp.GetOracleKeeper().SetAdminAddr(ctx, allowedAddresses[0], newAllowedAddresses, allowedAddresses)

	require.Equal(t, newAllowedAddresses, simapp.GetOracleKeeper().GetAllowedAddresses(ctx))
}

func TestSetAdminAddrSetAllowedAddressesFailure(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	allowedAddresses := []string{"abc"}
	newAllowedAddresses := []string{"def"}
	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)
	require.Error(t, simapp.GetOracleKeeper().SetAdminAddr(ctx, newAllowedAddresses[0], newAllowedAddresses, allowedAddresses))
}

func TestSetExchangeRate(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	allowedAddresses := []string{"abc"}
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
	oracleKeeper.SetExchangeRate(ctx, "abc", &rate)
	exrate, found := oracleKeeper.GetExchangeRate(ctx, rate.Pair)
	oracleKeeper.SetExchangeRate(ctx, "abc", &rate2)
	exrate2, found2 := oracleKeeper.GetExchangeRate(ctx, rate2.Pair)

	require.Equal(t, rate, exrate)
	require.Equal(t, true, found)
	require.Equal(t, rate2, exrate2)
	require.Equal(t, true, found2)

	aer := oracleKeeper.GetAllExchangeRates(ctx)
	require.Equal(t, []types.ExchangeRate{rate, rate2}, aer)
}

func TestSetExchangeRateFailure(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	allowedAddresses := []string{"abc"}
	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)

	rate := types.ExchangeRate{
		Pair: "udarc",
		Rate: uint64(1.2 * float64(1000000000000000000)),
	}
	oracleKeeper := simapp.GetOracleKeeper()
	require.Error(t, oracleKeeper.SetExchangeRate(ctx, "def", &rate))
}

func TestDeleteExchangeRate(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	allowedAddresses := []string{"abc"}
	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)

	rate1 := types.ExchangeRate{
		Pair: "kbtckusd",
		Rate: uint64(1.2 * float64(1000000000000000000)),
		Denoms:    []string{"kbtc", "kusd"},
		Timestamp: time.Now().Unix(),
	}
	rate2 := types.ExchangeRate{
		Pair: "kethkusd",
		Rate: uint64(1.3 * float64(1000000000000000000)),
		Denoms:    []string{"keth", "kusd"},
		Timestamp: time.Now().Unix(),
	}

	oracleKeeper := simapp.GetOracleKeeper()
	oracleKeeper.SetExchangeRate(ctx, "abc", &rate1)
	oracleKeeper.SetExchangeRate(ctx, "abc", &rate2)

	aer := oracleKeeper.GetAllExchangeRates(ctx)
	require.Equal(t, []types.ExchangeRate{rate1, rate2}, aer)

	require.Nil(t, oracleKeeper.DeleteExchangeRate(ctx, "abc", rate1.Pair))
	aer2 := oracleKeeper.GetAllExchangeRates(ctx)
	require.Equal(t, []types.ExchangeRate{rate2}, aer2)
}