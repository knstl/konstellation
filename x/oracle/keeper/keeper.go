package keeper

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/konstellation/konstellation/x/oracle/types"
)

// Keeper of the oracle store
type Keeper struct {
	storeKey   sdk.StoreKey
	paramSpace paramtypes.Subspace
	cdc        codec.BinaryMarshaler
}

// NewKeeper creates an oracle keeper
func NewKeeper(cdc codec.BinaryMarshaler, key sdk.StoreKey, paramSpace paramtypes.Subspace) Keeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	keeper := Keeper{
		storeKey:   key,
		paramSpace: paramSpace,
		cdc:        cdc,
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
	err := dec.Decode(&allowedAddresses)
	if err != nil {
		panic(err)
	}
	return
}

func (k Keeper) SetAllowedAddresses(ctx sdk.Context, sender string, newAllowedAddresses []string) error {
	allowedAddresses := k.GetAllowedAddresses(ctx)
	if !isValidSender(allowedAddresses, sender) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect sender") // If not, throw an error
	}
	preAddressListNum := len(allowedAddresses)
	for _, address := range newAllowedAddresses { // skip duplicated address
		if !isValidSender(allowedAddresses, address) {
			allowedAddresses = append(allowedAddresses, address)
		}
	}
	postAddressListNum := len(allowedAddresses)
	if preAddressListNum == postAddressListNum {
		return sdkerrors.Wrap(sdkerrors.ErrLogic, "no address to add") // If not, throw an error
	}
	var allowedAddressesBytes bytes.Buffer
	enc := gob.NewEncoder(&allowedAddressesBytes)
	err := enc.Encode(allowedAddresses)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(types.AllowedAddressKey, allowedAddressesBytes.Bytes())
	return nil
}

func (k Keeper) DeleteAllowedAddresses(ctx sdk.Context, sender string, addressesToDelete []string) error {
	allowedAddresses := k.GetAllowedAddresses(ctx)
	if !isValidSender(allowedAddresses, sender) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect sender") // If not, throw an error
	}
	for _, address := range addressesToDelete {
		allowedAddresses = removeAddress(allowedAddresses, address)
	}
	var allowedAddressesBytes bytes.Buffer
	enc := gob.NewEncoder(&allowedAddressesBytes)
	err := enc.Encode(allowedAddresses)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrLogic, err.Error()) // If not, throw an error
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(types.AllowedAddressKey, allowedAddressesBytes.Bytes())
	return nil
}

func removeAddress(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func (k Keeper) GetExchangeRate(ctx sdk.Context, pair string) (exchangeRate types.ExchangeRate, found bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.GetExchangeRateKey(pair))
	if b == nil {
		return exchangeRate, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &exchangeRate)
	return exchangeRate, true
}

func (k Keeper) GetAllExchangeRates(ctx sdk.Context) (rates []types.ExchangeRate) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ExchangeRateKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		rate := MustUnmarshalExchangeRate(k.cdc, iterator.Value())
		rates = append(rates, rate)
	}

	return rates
}

func (k Keeper) SetExchangeRate(ctx sdk.Context, sender string, rate *types.ExchangeRate) error {
	allowedAddresses := k.GetAllowedAddresses(ctx)
	if !isValidSender(allowedAddresses, sender) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect sender") // If not, throw an error
	}

	// todo check rate validity

	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryBare(rate)
	store.Set(types.GetExchangeRateKey(rate.Pair), b)
	return nil
}

func (k Keeper) DeleteExchangeRate(ctx sdk.Context, sender string, pair string) error {
	allowedAddresses := k.GetAllowedAddresses(ctx)
	if !isValidSender(allowedAddresses, sender) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect sender") // If not, throw an error
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetExchangeRateKey(pair))
	return nil
}

func (k Keeper) SetAdminAddr(ctx sdk.Context, sender string, add []string, del []string) error {
	if len(add) > 0 {
		err := k.SetAllowedAddresses(ctx, sender, add)
		if err != nil {
			return err
		}
	}
	if len(del) > 0 {
		err := k.DeleteAllowedAddresses(ctx, sender, del)
		if err != nil {
			return err
		}
	}
	return nil
}

func isValidSender(allowedAddresses []string, sender string) bool {
	for _, address := range allowedAddresses {
		if address == sender {
			return true
		}
	}
	return false
}

func MustMarshalExchangeRate(cdc codec.BinaryMarshaler, e *types.ExchangeRate) []byte {
	return cdc.MustMarshalBinaryBare(e)
}

func MustUnmarshalExchangeRate(cdc codec.BinaryMarshaler, value []byte) types.ExchangeRate {
	validator, err := UnmarshalExchangeRate(cdc, value)
	if err != nil {
		panic(err)
	}

	return validator
}

func UnmarshalExchangeRate(cdc codec.BinaryMarshaler, value []byte) (e types.ExchangeRate, err error) {
	err = cdc.UnmarshalBinaryBare(value, &e)
	return e, err
}
