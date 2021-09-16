package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/konstellation/konstellation/x/oracle/types"
)

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

func (k Keeper) SetAllowedAddressesInternal(ctx sdk.Context, addrs []*types.AdminAddr) error {
	for _, addr := range addrs {
		if err := k.setAllowedAddress(ctx, *addr); err != nil {
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
