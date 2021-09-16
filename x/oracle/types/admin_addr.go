package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewAdminAddr(addr string) *AdminAddr {
	return &AdminAddr{
		Address: addr,
	}
}

func (m AdminAddr) GetAdminAddress() sdk.AccAddress {
	if m.Address == "" {
		return nil
	}

	addr, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		panic(err)
	}

	return addr
}
