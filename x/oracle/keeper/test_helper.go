package keeper

import (
	"bytes"
	"encoding/gob"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/types"
)

func (k Keeper) SetTestAllowedAddresses(ctx sdk.Context, newAllowedAddresses []string) error {
	var allowedAddressesBytes bytes.Buffer
	enc := gob.NewEncoder(&allowedAddressesBytes)
	err := enc.Encode(newAllowedAddresses)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(types.AllowedAddressKey, allowedAddressesBytes.Bytes())
	return nil
}
