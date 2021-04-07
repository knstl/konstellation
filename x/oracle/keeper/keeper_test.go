package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/konstellation/konstellation/app"
)

func TestLogger(t *testing.T) {
	app := app.Setup(false)

	ctx := app.NewContext(true, tmproto.Header{})
	require.Equal(t, ctx.Logger(), app.GetOracleKeeper().Logger(ctx))
}
