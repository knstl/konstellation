package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDescription{}

func NewMsgDescription(owner sdk.AccAddress, denom string, description string) *MsgDescription {
	return &MsgDescription{
		Owner:       owner.String(),
		Denom:       denom,
		Description: description,
	}
}

func (msg *MsgDescription) Route() string {
	return RouterKey
}

func (msg *MsgDescription) Type() string {
	//return "Description"
	return TypeMsgDescription
}

func (msg *MsgDescription) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Owner)}
}

func (msg *MsgDescription) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDescription) ValidateBasic() error {
	if sdk.AccAddress(msg.Owner).Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing owner address")
	}
	if msg.Denom == "" {
		return ErrInvalidDenom(msg.Denom)
	}

	if len(msg.Description) > CoinDescriptionMaxLength {
		return ErrCoinDescriptionMaxLengthNotValid
	}
	return nil
}
