package simapp

import (
	"time"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/tendermint/spm/cosmoscmd"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
	tmdb "github.com/tendermint/tm-db"

	"github.com/konstellation/konstellation/app"
)

// New creates application instance with in-memory database and disabled logging.
func New(dir string) cosmoscmd.App {
	db := tmdb.NewMemDB()
	logger := log.NewNopLogger()

	//encoding := cosmoscmd.MakeEncodingConfig(app.ModuleBasics) // removed in APP doesn't enconding, moved that inside the function

	app := app.New(logger, db, nil, false, map[int64]bool{}, dir, uint(1), simapp.EmptyAppOptions{}, app.GetWasmEnabledProposals(), app.EmptyWasmOpts)
	// InitChain updates deliverState which is required when app.NewContext is called
	app.InitChain(abci.RequestInitChain{
		ConsensusParams: defaultConsensusParams,
		AppStateBytes:   []byte("{}"),
	})
	return app
}

var defaultConsensusParams = &abci.ConsensusParams{
	Block: &abci.BlockParams{
		MaxBytes: 200000,
		MaxGas:   2000000,
	},
	Evidence: &tmproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &tmproto.ValidatorParams{
		PubKeyTypes: []string{
			tmtypes.ABCIPubKeyTypeEd25519,
		},
	},
}
