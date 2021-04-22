package rest_test

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"

	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	testnet "github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/konstellation/app"
	oracletypes "github.com/konstellation/konstellation/x/oracle/types"
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
	s.T().Log("setting up integration test suite")

	cfg := testnet.DefaultConfig()

	encCfg := app.MakeEncodingConfig()
	cfg.Codec = encCfg.Marshaler
	cfg.TxConfig = encCfg.TxConfig
	cfg.LegacyAmino = encCfg.Amino
	cfg.InterfaceRegistry = encCfg.InterfaceRegistry

	simapp := app.Setup(false)
	cfg.GenesisState = app.ModuleBasics.DefaultGenesis(encCfg.Marshaler)
	genesisState := cfg.GenesisState
	cfg.NumValidators = 1

	//cfg.GenesisState[oracletypes.ModuleName] = []byte(`{"allowed_address": "abc"}`)
	var oracleData oracletypes.GenesisState
	s.Require().NoError(cfg.Codec.UnmarshalJSON(genesisState[oracletypes.ModuleName], &oracleData))
	oracleData.AllowedAddresses = []string{"abc"}
	oracleDataBz, err := cfg.Codec.MarshalJSON(&oracleData)
	s.Require().NoError(err)
	genesisState[oracletypes.ModuleName] = oracleDataBz
	cfg.GenesisState = genesisState

	cfg.AppConstructor = NewAppConstructor(simapp)

	s.cfg = cfg
	s.network = testnet.New(s.T(), cfg)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestQueryGRPC() {
	val := s.network.Validators[0]
	baseURL := val.APIAddress
	coin := sdk.NewCoin("Darc", sdk.NewInt(10))
	testCases := []struct {
		name     string
		url      string
		headers  map[string]string
		respType proto.Message
		expected proto.Message
	}{
		{
			"gRPC request exchange rate",
			fmt.Sprintf("%s/konstellation/oracle/exchange_rate", baseURL),
			map[string]string{},
			&oracletypes.QueryExchangeRateResponse{},
			&oracletypes.QueryExchangeRateResponse{
				ExchangeRate: coin,
			},
		},
	}
	for _, tc := range testCases {
		resp, err := testutil.GetRequestWithHeaders(tc.url, tc.headers)
		fmt.Println(string(resp))
		s.Run(tc.name, func() {
			s.Require().NoError(err)
			s.Require().NoError(val.ClientCtx.JSONMarshaler.UnmarshalJSON(resp, tc.respType))
			s.Require().Equal(tc.expected.String(), tc.respType.String())
		})
	}
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
