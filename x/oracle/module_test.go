package oracle_test

import (
	"github.com/konstellation/konstellation/types"
	types2 "github.com/konstellation/konstellation/x/oracle/types"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/konstellation/konstellation/app"
)

func TestMain(m *testing.M) {
	// Write code here to run before tests
	types.RegisterBech32Prefix()

	// Run tests
	exitVal := m.Run()

	// Write code here to run after tests

	// Exit with exit value from tests
	os.Exit(exitVal)
}

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	simapp := app.Setup(false)
	simapp.Commit()
	ctx := simapp.NewContext(true, tmproto.Header{})
	simapp.GetOracleKeeper().SetTestAllowedAddresses(ctx, []types2.AdminAddr{*types2.NewAdminAddr("darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx")})
	acc := simapp.GetOracleKeeper().GetAllowedAddresses(ctx)
	require.NotNil(t, acc)
}
