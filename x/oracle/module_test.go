package oracle_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	abcitypes "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/konstellation/konstellation/app"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	simapp := app.Setup(false)
	ctx := simapp.NewContext(true, tmproto.Header{})

	genesisState := app.NewDefaultGenesisState()
	stateBytes, err := json.MarshalIndent(genesisState, "", "  ")
	require.NoError(t, err)

	simapp.InitChain(
		abcitypes.RequestInitChain{
			ChainId:       "test-chain-id",
			AppStateBytes: stateBytes,
		},
	)
	simapp.Commit()

	acc := simapp.GetOracleKeeper().GetAllowedAddress(ctx)
	require.NotNil(t, acc)
}
