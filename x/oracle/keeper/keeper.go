package keeper

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/types"
)

// Keeper of the oracle store
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.BinaryMarshaler
}

// NewKeeper creates an oracle keeper
func NewKeeper(cdc codec.BinaryMarshaler, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		storeKey: key,
		cdc:      cdc,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetAllowedAddresses(ctx sdk.Context) (allowedAddresses []string) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.AllowedAddressKey)
	if b == nil {
		panic("stored allowed address should not have been nil")
	}

	allowedAddressesBytes := bytes.NewBuffer(b)
	dec := gob.NewDecoder(allowedAddressesBytes)
	dec.Decode(&allowedAddresses)
	return
}

func (k Keeper) SetAllowedAddresses(ctx sdk.Context, allowedAddresses []string) {
	var allowedAddressesBytes bytes.Buffer
	enc := gob.NewEncoder(&allowedAddressesBytes)
	enc.Encode(allowedAddresses)

	store := ctx.KVStore(k.storeKey)
	store.Set(types.AllowedAddressKey, allowedAddressesBytes.Bytes())
}

func (k Keeper) DeleteAllowedAddresses(ctx sdk.Context, addressesToDelete []string) {
	allowedAddresses := k.GetAllowedAddresses(ctx)
	for _, address := range addressesToDelete {
		allowedAddresses = removeAddress(allowedAddresses, address)
	}
	var allowedAddressesBytes bytes.Buffer
	enc := gob.NewEncoder(&allowedAddressesBytes)
	enc.Encode(allowedAddresses)

	store := ctx.KVStore(k.storeKey)
	store.Set(types.AllowedAddressKey, allowedAddressesBytes.Bytes())
}

func removeAddress(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func (k Keeper) GetExchangeRate(ctx sdk.Context) (exchangeRate sdk.Coin) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.ExchangeRateKey)
	if b == nil {
		panic("stored exchange rate should not have been nil")
	}

	k.cdc.MustUnmarshalBinaryBare(b, &exchangeRate)
	return
}

func (k Keeper) SetExchangeRate(ctx sdk.Context, exchangeRate sdk.Coin) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryBare(&exchangeRate)
	store.Set(types.ExchangeRateKey, b)
}

func (k Keeper) DeleteExchangeRate(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.ExchangeRateKey)
}

func (k Keeper) SetAdminAddr(ctx sdk.Context, sender string, add []string, del []string) {
	if len(add) > 0 {
		k.SetAllowedAddresses(ctx, add)
	}
	if len(del) > 0 {
		k.DeleteAllowedAddresses(ctx, add)
	}
}
