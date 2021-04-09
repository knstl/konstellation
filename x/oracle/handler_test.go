package oracle_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/crisis/types"
	"github.com/konstellation/konstellation/app"
	"github.com/konstellation/konstellation/x/oracle"
	"github.com/konstellation/konstellation/x/oracle/simulation"
	oracletypes "github.com/konstellation/konstellation/x/oracle/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	testModuleName        = "dummy"
	dummyRouteWhichPasses = types.NewInvarRoute(testModuleName, "which-passes", func(_ sdk.Context) (string, bool) { return "", false })
	dummyRouteWhichFails  = types.NewInvarRoute(testModuleName, "which-fails", func(_ sdk.Context) (string, bool) { return "whoops", true })
)

func TestHandleMsgSetExchangeRate(t *testing.T) {
	simapp := app.Setup(false)
	simapp.Commit()
	ctx := simapp.NewContext(true, tmproto.Header{})
	coin := sdk.NewCoin("Darc", sdk.NewInt(10))
	rand := rand.New(rand.NewSource(int64(1)))
	address := simulation.RandomAddress(rand)
	incorrectMsg := oracletypes.NewMsgSetExchangeRate(&coin, address)
	correctMsg := oracletypes.NewMsgSetExchangeRate(&coin, simapp.GetOracleKeeper().GetAllowedAddress(ctx))

	cases := []struct {
		name           string
		msg            sdk.Msg
		expectedResult string
	}{
		{"not allowed address", &incorrectMsg, "fail"},
		{"set_exchange_rate", &correctMsg, "pass"},
		{"invalid msg", testdata.NewTestMsg(), "fail"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			h := oracle.NewHandler(simapp.GetOracleKeeper())

			switch tc.expectedResult {
			case "fail":
				res, err := h(ctx, tc.msg)
				require.Error(t, err)
				require.Nil(t, res)

			case "pass":
				res, err := h(ctx, tc.msg)
				require.NoError(t, err)
				require.NotNil(t, res)

			case "panic":
				require.Panics(t, func() {
					h(ctx, tc.msg) // nolint:errcheck
				})
			}
		})
	}
}

func TestHandleMsgDeleteExchangRate(t *testing.T) {
	simapp := app.Setup(false)
	simapp.Commit()
	ctx := simapp.NewContext(true, tmproto.Header{})
	rand := rand.New(rand.NewSource(int64(1)))
	address := simulation.RandomAddress(rand)
	incorrectMsg := oracletypes.NewMsgDeleteExchangeRate("Darc", address)
	correctMsg := oracletypes.NewMsgDeleteExchangeRate("Darc", simapp.GetOracleKeeper().GetAllowedAddress(ctx))

	cases := []struct {
		name           string
		msg            sdk.Msg
		expectedResult string
	}{
		{"not allowed address", &incorrectMsg, "fail"},
		{"delete_exchange_rate", &correctMsg, "pass"},
		{"invalid msg", testdata.NewTestMsg(), "fail"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			h := oracle.NewHandler(simapp.GetOracleKeeper())

			switch tc.expectedResult {
			case "fail":
				res, err := h(ctx, tc.msg)
				require.Error(t, err)
				require.Nil(t, res)

			case "pass":
				res, err := h(ctx, tc.msg)
				require.NoError(t, err)
				require.NotNil(t, res)

			case "panic":
				require.Panics(t, func() {
					h(ctx, tc.msg) // nolint:errcheck
				})
			}
		})
	}
}
