// +build norace

package cli_test

import (
	"github.com/stretchr/testify/suite"

	testnet "github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/client/cli"
	oracletypes "github.com/konstellation/konstellation/x/oracle/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     testnet.Config
	network *testnet.Network
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := testnet.DefaultConfig()
	genesisState := cfg.GenesisState
	cfg.NumValidators = 1

	var oracleData oracletypes.GenesisState
	s.Require().NoError(cfg.Codec.UnmarshalJSON(genesisState[oracletypes.ModuleName], &oracleData))

	oracleData.AllowedAddress = "abc"
	oracleDataBz, err := cfg.Codec.MarshalJSON(&oracleData)
	s.Require().NoError(err)
	genesisState[oracletypes.ModuleName] = oracleDataBz
	cfg.GenesisState = genesisState

	s.cfg = cfg
	s.network = testnet.New(s.T(), cfg)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

//TODO: check how to make test case with protobuf
