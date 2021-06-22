/*
// +build norace
*/

package cli_test

import (
	"testing"

	//	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"

	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	testnet "github.com/cosmos/cosmos-sdk/testutil/network"
	//	sdk "github.com/cosmos/cosmos-sdk/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     testnet.Config
	network *testnet.Network
}

func NewAppConstructor(simapp servertypes.Application) testnet.AppConstructor {
	return func(val testnet.Validator) servertypes.Application {
		return simapp
	}
}

func (s *IntegrationTestSuite) SetupSuite() {
	//s.T().Log("setting up integration test suite")
	//
	//cfg := testnet.DefaultConfig()
	//
	//encCfg := app.MakeEncodingConfig()
	//cfg.Codec = encCfg.Marshaler
	//cfg.TxConfig = encCfg.TxConfig
	//cfg.LegacyAmino = encCfg.Amino
	//cfg.InterfaceRegistry = encCfg.InterfaceRegistry
	//
	//simapp := app.Setup(false)
	//cfg.GenesisState = app.ModuleBasics.DefaultGenesis(encCfg.Marshaler)
	//genesisState := cfg.GenesisState
	//cfg.NumValidators = 1
	//
	////cfg.GenesisState[oracletypes.ModuleName] = []byte(`{"allowed_address": "darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx"}`)
	//var oracleData oracletypes.GenesisState
	//s.Require().NoError(cfg.Codec.UnmarshalJSON(genesisState[oracletypes.ModuleName], &oracleData))
	//oracleData.AllowedAddresses = []string{"darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx"}
	//oracleDataBz, err := cfg.Codec.MarshalJSON(&oracleData)
	//s.Require().NoError(err)
	//genesisState[oracletypes.ModuleName] = oracleDataBz
	//cfg.GenesisState = genesisState
	//
	//cfg.AppConstructor = NewAppConstructor(simapp)
	//
	//s.cfg = cfg
	//s.network = testnet.New(s.T(), cfg)
	//
	//_, err = s.network.WaitForHeight(1)
	//s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	//s.T().Log("tearing down integration test suite")
	//s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestGetCmdQueryExchangeRate() {
	//val := s.network.Validators[0]
	//
	//testCases := []struct {
	//	name           string
	//	args           []string
	//	expectedOutput string
	//}{
	//	{
	//		"json output",
	//		[]string{"exchange-rate"},
	//		`{"exchange_rate": {"denom": "Darc", "amount": "10"}, setter:"darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx"}`,
	//	},
	//}
	//
	//for _, tc := range testCases {
	//	tc := tc
	//
	//	s.Run(tc.name, func() {
	//		cmd := cli.GetQueryCmd()
	//		clientCtx := val.ClientCtx
	//
	//		out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
	//		s.Require().NoError(err)
	//		s.Require().Equal(tc.expectedOutput, strings.TrimSpace(out.String()))
	//	})
	//}
}

/*
func (s *IntegrationTestSuite) TestNewMsgSetExchangeReteCmd() {
	val := s.network.Validators[0]

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"valid transaction",
			[]string{
				"darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx",
				"Darc",
				"10",
			},
			false, &sdk.TxResponse{}, 0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.NewMsgSetExchangeRateCmd()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code)
			}
		})
	}
}

/*
func (s *IntegrationTestSuite) TestNewMsgDeleteExchangeReteCmd() {
	val := s.network.Validators[0]

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"valid transaction",
			[]string{
				"darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx",
			},
			false, &sdk.TxResponse{}, 0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.NewMsgDeleteExchangeRateCmd()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestNewMsgSetAdminAddrCmd() {
	val := s.network.Validators[0]

	testCases := []struct {
		name         string
		args         []string
		expectErr    bool
		respType     proto.Message
		expectedCode uint32
	}{
		{
			"valid transaction",
			[]string{
				"darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx",
			},
			false, &sdk.TxResponse{}, 0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.NewMsgSetAdminAddrCmd()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), tc.respType), out.String())

				txResp := tc.respType.(*sdk.TxResponse)
				s.Require().Equal(tc.expectedCode, txResp.Code)
			}
		})
	}
}

*/

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
