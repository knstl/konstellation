package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTransferOwnership{}

func NewMsgTransferOwnership(owner, toAddr sdk.AccAddress, denom string) *MsgTransferOwnership {
	return &MsgTransferOwnership{
		Owner:     owner.String(),
		ToAddress: toAddr.String(),
		Denom:     denom,
	}
}

func (m *MsgTransferOwnership) Route() string {
	return RouterKey
}

func (m *MsgTransferOwnership) Type() string {
	return TypeMsgTransferOwnership
}

func (m *MsgTransferOwnership) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(m.Owner)}
}

func (m *MsgTransferOwnership) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgTransferOwnership) ValidateBasic() error {
	if sdk.AccAddress(m.Owner).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing owner address")
	}
	if sdk.AccAddress(m.ToAddress).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing recipient address")
	}
	if m.Denom == "" {
		return ErrInvalidDenom(m.Denom)
	}
	return nil
}
