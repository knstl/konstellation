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

func (msg *MsgTransferOwnership) Route() string {
	return RouterKey
}

func (msg *MsgTransferOwnership) Type() string {
	return TypeMsgTransferOwnership
}

func (msg *MsgTransferOwnership) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func (msg *MsgTransferOwnership) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransferOwnership) ValidateBasic() error {
	if sdk.AccAddress(msg.Owner).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing owner address")
	}
	if sdk.AccAddress(msg.ToAddress).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing recipient address")
	}
	if msg.Denom == "" {
		return ErrInvalidDenom(msg.Denom)
	}
	return nil
}
