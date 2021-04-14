package keeper

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func (k Keeper) GetExchangeRate(ctx sdk.Context) (exchangeRate sdk.Coin) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.ExchangeRateKey)
	if b == nil {
		panic("stored exchange rate should not have been nil")
	}

	k.cdc.MustUnmarshalBinaryBare(b, &exchangeRate)
	return
}

func (k Keeper) SetExchangeRate(ctx sdk.Context, sender string, exchangeRate sdk.Coin) error {
	allowedAddresses := k.GetAllowedAddresses(ctx)
	if !isValidSender(allowedAddresses, sender) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect sender") // If not, throw an error
	}

	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryBare(&exchangeRate)
	store.Set(types.ExchangeRateKey, b)
	return nil
}

func (k Keeper) DeleteExchangeRate(ctx sdk.Context, sender string) error {
	allowedAddresses := k.GetAllowedAddresses(ctx)
	if !isValidSender(allowedAddresses, sender) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect sender") // If not, throw an error
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete(types.ExchangeRateKey)
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
