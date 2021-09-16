package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/types"
)

func (k Keeper) SetTestAllowedAddresses(ctx sdk.Context, addrs []types.AdminAddr) error {
	for _, addr := range addrs {
		if err := k.setAllowedAddress(ctx, addr); err != nil {
			return err
		}
	}
	return nil
}
