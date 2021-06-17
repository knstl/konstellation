package keeper_test

import (
	gocontext "context"
	"testing"

	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/app"
	"github.com/konstellation/konstellation/x/oracle/types"
)

type OracleTestSuite struct {
	suite.Suite

	app         *app.KonstellationApp
	ctx         sdk.Context
	queryClient types.QueryClient
}

func (suite *OracleTestSuite) SetupTest() {
	simapp := app.Setup(false)
	simapp.Commit()
	ctx := simapp.NewContext(true, tmproto.Header{})

	d, _ := sdk.NewDecFromStr("1200000000000")
	rate := types.ExchangeRate{
		Pair: "udarc",
		Rate: d,
	}

	abc := types.NewAdminAddr("darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx")
	allowedAddresses := []types.AdminAddr{*abc}
	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, allowedAddresses)
	simapp.GetOracleKeeper().SetExchangeRate(ctx, abc.GetAdminAddress(), &rate)

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, codectypes.NewInterfaceRegistry())
	types.RegisterQueryServer(queryHelper, simapp.GetOracleKeeper())
	queryClient := types.NewQueryClient(queryHelper)

	suite.app = simapp
	suite.ctx = ctx

	suite.queryClient = queryClient
}

func (suite *OracleTestSuite) TestGRPCExchangeRate() {
	app, ctx, queryClient := suite.app, suite.ctx, suite.queryClient

	rate, err := queryClient.ExchangeRate(gocontext.Background(), &types.QueryExchangeRateRequest{})
	suite.Require().NoError(err)
	r, _ := app.GetOracleKeeper().GetExchangeRate(ctx, rate.ExchangeRate.Pair)
	suite.Require().Equal(*rate.ExchangeRate, r)
}

func TestOracleTestSuite(t *testing.T) {
	suite.Run(t, new(OracleTestSuite))
}
