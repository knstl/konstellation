package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/konstellation/konstellation/x/oracle/types"
)

func (k Keeper) GetAllExchangeRates(ctx sdk.Context) (rates []types.ExchangeRate) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ExchangeRateKeyValue)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		rate := MustUnmarshalExchangeRate(k.cdc, iterator.Value())
		rates = append(rates, rate)
	}

	return rates
}

func (k Keeper) setExchangeRate(ctx sdk.Context, rate *types.ExchangeRate) error {
	store := ctx.KVStore(k.storeKey)

	rate.Height = int32(ctx.BlockHeight())
	rate.Timestamp = ctx.BlockHeader().Time

	b := k.cdc.MustMarshal(rate)
	store.Set(types.GetExchangeRateKey(rate.Pair), b)
	return nil
}

func (k Keeper) deleteExchangeRate(ctx sdk.Context, pair string) error {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetExchangeRateKey(pair))
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

func (k Keeper) GetExchangeRate(ctx sdk.Context, pair string) (exchangeRate types.ExchangeRate, found bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.GetExchangeRateKey(pair))
	if b == nil {
		return exchangeRate, false
	}

	k.cdc.MustUnmarshal(b, &exchangeRate)
	return exchangeRate, true
}

func MustMarshalExchangeRate(cdc codec.BinaryCodec, r *types.ExchangeRate) []byte {
	return cdc.MustMarshal(r)
}

func MustUnmarshalExchangeRate(cdc codec.BinaryCodec, value []byte) types.ExchangeRate {
	r, err := UnmarshalExchangeRate(cdc, value)
	if err != nil {
		panic(err)
	}

	return r
}

func UnmarshalExchangeRate(cdc codec.BinaryCodec, value []byte) (r types.ExchangeRate, err error) {
	err = cdc.Unmarshal(value, &r)
	return r, err
}
