package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc auth module wide codec
var ModuleCdc = codec.New()

// RegisterCodec registers concrete types on the codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgIssueCreate{}, "issue/MsgIssueCreate", nil)
	cdc.RegisterConcrete(MsgDisableFeature{}, "issue/MsgDisableFeature", nil)
	cdc.RegisterConcrete(MsgEnableFeature{}, "issue/MsgEnableFeature", nil)
	cdc.RegisterConcrete(MsgFeatures{}, "issue/MsgFeatures", nil)
	cdc.RegisterConcrete(MsgDescription{}, "issue/MsgDescription", nil)
	cdc.RegisterConcrete(MsgTransfer{}, "issue/MsgTransfer", nil)
	cdc.RegisterConcrete(MsgTransferFrom{}, "issue/MsgTransferFrom", nil)
	cdc.RegisterConcrete(MsgApprove{}, "issue/MsgApprove", nil)
	cdc.RegisterConcrete(MsgIncreaseAllowance{}, "issue/MsgIncreaseAllowance", nil)
	cdc.RegisterConcrete(MsgDecreaseAllowance{}, "issue/MsgDecreaseAllowance", nil)
	cdc.RegisterConcrete(MsgMint{}, "issue/MsgMint", nil)
	cdc.RegisterConcrete(MsgBurn{}, "issue/MsgBurn", nil)
	cdc.RegisterConcrete(MsgBurnFrom{}, "issue/MsgBurnFrom", nil)
	cdc.RegisterConcrete(MsgTransferOwnership{}, "issue/MsgTransferOwnership", nil)
	cdc.RegisterConcrete(MsgFreeze{}, "issue/MsgFreeze", nil)
	cdc.RegisterConcrete(MsgUnfreeze{}, "issue/MsgUnfreeze", nil)

	cdc.RegisterInterface((*IIssue)(nil), nil)
	cdc.RegisterConcrete(&CoinIssue{}, "issue/CoinIssue", nil)
	cdc.RegisterConcrete(Params{}, "issue/Params", nil)
}

func init() {
	RegisterCodec(ModuleCdc)
}
