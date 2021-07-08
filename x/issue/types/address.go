package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AccAddress sdk.AccAddress

// Size implements the gogo proto custom type interface.
func (aa AccAddress) Size() int {
	addr := sdk.AccAddress(aa)

	bz, _ := addr.Marshal()
	return len(bz)
}

// MarshalTo implements the gogo proto custom type interface.
func (aa AccAddress) MarshalTo(data []byte) (n int, err error) {
	if aa == nil {
		aa = AccAddress{}
	}
	addr := sdk.AccAddress(aa)
	if len(addr.Bytes()) == 0 {
		copy(data, []byte{0x30})
		return 1, nil
	}

	bz, err := addr.Marshal()
	if err != nil {
		return 0, err
	}

	copy(data, bz)
	return len(bz), nil
}

// Unmarshal implements the gogo proto custom type interface.
func (aa AccAddress) Unmarshal(data []byte) error {
	addr := sdk.AccAddress(aa)
	err := addr.Unmarshal(data)
	if err != nil {
		return err
	}
	return nil
}
