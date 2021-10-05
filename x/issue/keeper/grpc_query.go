package keeper

import (
	"github.com/konstellation/konstellation/x/issue/types"
)

type queryServer struct {
	keeper Keeper
}

// NewQueryServerImpl returns an implementation of the QueryServer interface
// for the provided Keeper.
func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return &queryServer{keeper: keeper}
}

var _ types.QueryServer = queryServer{}
