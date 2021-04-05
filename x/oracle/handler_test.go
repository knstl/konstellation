package oracle_test

import (
	"encoding/json"
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
	abcitypes "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	testModuleName        = "dummy"
	dummyRouteWhichPasses = types.NewInvarRoute(testModuleName, "which-passes", func(_ sdk.Context) (string, bool) { return "", false })
	dummyRouteWhichFails  = types.NewInvarRoute(testModuleName, "which-fails", func(_ sdk.Context) (string, bool) { return "whoops", true })
)

func createTestApp() (*app.KonstellationApp, sdk.Context) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	genesisState := app.NewDefaultGenesisState()
	stateBytes, _ := json.MarshalIndent(genesisState, "", "  ")

	simapp.InitChain(
		abcitypes.RequestInitChain{
			ChainId:       "test-chain-id",
			AppStateBytes: stateBytes,
		},
	)
	simapp.Commit()

	return simapp, ctx
}

func TestHandleMsgSetExchangeRate(t *testing.T) {
	simapp, ctx := createTestApp()
	coin := sdk.NewCoin("Darc", sdk.NewInt(10))
	rand := rand.New(rand.NewSource(int64(1)))
	address := simulation.RandomAddress(rand)

	cases := []struct {
		name           string
		msg            sdk.Msg
		expectedResult string
	}{
		{"not allowed address", oracletypes.NewMsgSetExchangeRate(&coin, address), "fail"},
		{"set_exchange_rate", oracletypes.NewMsgSetExchangeRate(&coin, simapp.GetOracleKeeper().GetAllowedAddress(ctx)), "pass"},
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
	simapp, ctx := createTestApp()
	rand := rand.New(rand.NewSource(int64(1)))
	address := simulation.RandomAddress(rand)

	cases := []struct {
		name           string
		msg            sdk.Msg
		expectedResult string
	}{
		{"not allowed address", oracletypes.NewMsgDeleteExchangeRate("Darc", address), "fail"},
		{"delete_exchange_rate", oracletypes.NewMsgDeleteExchangeRate("Darc", simapp.GetOracleKeeper().GetAllowedAddress(ctx)), "pass"},
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
