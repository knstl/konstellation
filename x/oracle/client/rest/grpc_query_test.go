package rest_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	testnet "github.com/cosmos/cosmos-sdk/testutil/network"
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

func (s *IntegrationTestSuite) TestQueryGRPC() {
	//val := s.network.Validators[0]
	//baseURL := val.APIAddress
	//rate := oracletypes.ExchangeRate{
	//	Denom: "udarc",
	//	Rate:  1.2,
	//}
	//testCases := []struct {
	//	name     string
	//	url      string
	//	headers  map[string]string
	//	respType proto.Message
	//	expected proto.Message
	//}{
	//	{
	//		"gRPC request exchange rate",
	//		fmt.Sprintf("%s/konstellation/oracle/exchange_rate", baseURL),
	//		map[string]string{},
	//		&oracletypes.QueryExchangeRateResponse{},
	//		&oracletypes.QueryExchangeRateResponse{
	//			ExchangeRate: &rate,
	//		},
	//	},
	//}
	//for _, tc := range testCases {
	//	resp, err := testutil.GetRequestWithHeaders(tc.url, tc.headers)
	//	fmt.Println(string(resp))
	//	s.Run(tc.name, func() {
	//		s.Require().NoError(err)
	//		s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(resp, tc.respType))
	//		s.Require().Equal(tc.expected.String(), tc.respType.String())
	//	})
	//}
}

func TestIntegrationTestSuite(t *testing.T) {
	//suite.Run(t, new(IntegrationTestSuite))
}
