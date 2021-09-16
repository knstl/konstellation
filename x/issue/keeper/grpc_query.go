package keeper

import (
	"github.com/konstellation/konstellation/x/issue/types"
)

var _ types.QueryServer = Keeper{}
