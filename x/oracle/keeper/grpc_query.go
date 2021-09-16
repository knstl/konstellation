package keeper

import (
	"github.com/konstellation/konstellation/x/oracle/types"
)

var _ types.QueryServer = Keeper{}
