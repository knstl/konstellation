package keeper

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func (k Keeper) getAllowedAddress(ctx sdk.Context, addr sdk.AccAddress) (adm types.AdminAddr, found bool) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetAllowedAddressKey(addr))
	if value == nil {
		return adm, false
	}

	adm = MustUnmarshalAllowedAddress(k.cdc, value)
	return adm, true
}

func (k Keeper) setAllowedAddress(ctx sdk.Context, addr types.AdminAddr) error {
	store := ctx.KVStore(k.storeKey)

	b := k.cdc.MustMarshalBinaryBare(&addr)
	store.Set(types.GetAllowedAddressKey(addr.GetAdminAddress()), b)
	return nil
}

func (k Keeper) deleteAllowedAddress(ctx sdk.Context, addr types.AdminAddr) error {
	store := ctx.KVStore(k.storeKey)

	store.Delete(types.GetAllowedAddressKey(addr.GetAdminAddress()))
	return nil
}

func (k Keeper) GetAllowedAddresses(ctx sdk.Context) (allowedAddresses []types.AdminAddr) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.AllowedAddressKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		addr := MustUnmarshalAllowedAddress(k.cdc, iterator.Value())
		allowedAddresses = append(allowedAddresses, addr)
	}

	return allowedAddresses
}

func (k Keeper) IsAllowedAddress(ctx sdk.Context, addr sdk.AccAddress) bool {
	_, found := k.getAllowedAddress(ctx, addr)
	return found
}

func (k Keeper) GetAllowedAddress(ctx sdk.Context, addr sdk.AccAddress) (types.AdminAddr, bool) {
	return k.getAllowedAddress(ctx, addr)
}

func (k Keeper) SetAllowedAddresses(ctx sdk.Context, sender sdk.AccAddress, addrs []types.AdminAddr) error {
	if !k.IsAllowedAddress(ctx, sender) {
		return sdkerrors.Wrap(types.ErrAddressIsNotAdmin, "Sender address is not admin")
	}

	for _, addr := range addrs {
		if err := k.setAllowedAddress(ctx, addr); err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) DeleteAllowedAddresses(ctx sdk.Context, sender sdk.AccAddress, addrs []types.AdminAddr) error {
	if !k.IsAllowedAddress(ctx, sender) {
		return sdkerrors.Wrap(types.ErrAddressIsNotAdmin, "Sender address is not admin")
	}

	for _, addr := range addrs {
		if err := k.deleteAllowedAddress(ctx, addr); err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) SetAdminAddr(ctx sdk.Context, sender sdk.AccAddress, add []types.AdminAddr, del []types.AdminAddr) error {
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

func (k Keeper) setExchangeRate(ctx sdk.Context, rate *types.ExchangeRate) error {
	store := ctx.KVStore(k.storeKey)

	rate.Height = ctx.BlockHeight()
	rate.Timestamp = ctx.BlockHeader().Time

	b := k.cdc.MustMarshalBinaryBare(rate)
	store.Set(types.GetExchangeRateKey(rate.Pair), b)
	return nil
}

func (k Keeper) SetExchangeRate(ctx sdk.Context, sender sdk.AccAddress, rate *types.ExchangeRate) error {
	if !k.IsAllowedAddress(ctx, sender) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "sender address is not admin")
	}

	return k.setExchangeRate(ctx, rate)
}

func (k Keeper) SetExchangeRates(ctx sdk.Context, sender sdk.AccAddress, rates []*types.ExchangeRate) error {
	if !k.IsAllowedAddress(ctx, sender) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "sender address is not admin")
	}
	// todo check rate validity

	for _, r := range rates {
		if err := k.setExchangeRate(ctx, r); err != nil {
			return err
		}
	}
	return nil
}

func (k Keeper) deleteExchangeRate(ctx sdk.Context, pair string) error {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetExchangeRateKey(pair))
	return nil
}

func (k Keeper) DeleteExchangeRate(ctx sdk.Context, sender sdk.AccAddress, pair string) error {
	if !k.IsAllowedAddress(ctx, sender) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Sender address is not admin")
	}

	return k.deleteExchangeRate(ctx, pair)
}

func (k Keeper) DeleteExchangeRates(ctx sdk.Context, sender sdk.AccAddress, pairs []string) error {
	if !k.IsAllowedAddress(ctx, sender) {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Sender address is not admin")
	}

	for _, pair := range pairs {
		if err := k.deleteExchangeRate(ctx, pair); err != nil {
			return err
		}
	}

	return nil
}

func MustMarshalExchangeRate(cdc codec.BinaryMarshaler, r *types.ExchangeRate) []byte {
	return cdc.MustMarshalBinaryBare(r)
}

func MustUnmarshalExchangeRate(cdc codec.BinaryMarshaler, value []byte) types.ExchangeRate {
	r, err := UnmarshalExchangeRate(cdc, value)
	if err != nil {
		panic(err)
	}

	return r
}

func UnmarshalExchangeRate(cdc codec.BinaryMarshaler, value []byte) (r types.ExchangeRate, err error) {
	err = cdc.UnmarshalBinaryBare(value, &r)
	return r, err
}

func MustMarshalAllowedAddress(cdc codec.BinaryMarshaler, addr *types.AdminAddr) []byte {
	return cdc.MustMarshalBinaryBare(addr)
}

func MustUnmarshalAllowedAddress(cdc codec.BinaryMarshaler, value []byte) types.AdminAddr {
	addr, err := UnmarshalAllowedAddress(cdc, value)
	if err != nil {
		panic(err)
	}

	return addr
}

func UnmarshalAllowedAddress(cdc codec.BinaryMarshaler, value []byte) (addr types.AdminAddr, err error) {
	err = cdc.UnmarshalBinaryBare(value, &addr)
	return addr, err
}
