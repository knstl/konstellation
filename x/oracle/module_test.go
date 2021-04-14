package oracle_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/konstellation/konstellation/app"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	simapp := app.Setup(false)
	simapp.Commit()
	ctx := simapp.NewContext(true, tmproto.Header{})
	acc := simapp.GetOracleKeeper().GetAllowedAddresses(ctx)
	require.NotNil(t, acc)
}
