package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/konstellation/konstellation/app"
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

func TestSetExchangeRate(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	allowedAddresses := []string{"abc"}
	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)

	coin := sdk.NewCoin("Darc", sdk.NewInt(10))
	oracleKeeper := simapp.GetOracleKeeper()
	oracleKeeper.SetExchangeRate(ctx, "abc", coin)

	require.Equal(t, coin, simapp.GetOracleKeeper().GetExchangeRate(ctx))
}

func TestSetExchangeRateFailure(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	allowedAddresses := []string{"abc"}
	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)

	coin := sdk.NewCoin("Darc", sdk.NewInt(10))
	oracleKeeper := simapp.GetOracleKeeper()
	require.Error(t, oracleKeeper.SetExchangeRate(ctx, "def", coin))
}

func TestDeleteExchangeRate(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	allowedAddresses := []string{"abc"}
	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)

	coin := sdk.NewCoin("Darc", sdk.NewInt(10))
	oracleKeeper := simapp.GetOracleKeeper()
	oracleKeeper.SetExchangeRate(ctx, "abc", coin)

	require.Nil(t, oracleKeeper.DeleteExchangeRate(ctx, "abc"))
}

func TestDeleteExchangeRateFailure(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	allowedAddresses := []string{"abc"}
	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)

	coin := sdk.NewCoin("Darc", sdk.NewInt(10))
	oracleKeeper := simapp.GetOracleKeeper()
	oracleKeeper.SetExchangeRate(ctx, "abc", coin)

	require.Error(t, oracleKeeper.DeleteExchangeRate(ctx, "def"))
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
