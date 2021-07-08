package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// ModuleCdc auth module wide codec
var ModuleCdc = codec.NewLegacyAmino()

// RegisterCodec registers concrete types on the codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
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
	RegisterLegacyAminoCodec(ModuleCdc)
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgIssueCreate{},
		&MsgFeatures{},
		&MsgDescription{},
		&MsgTransfer{},
		&MsgTransferFrom{},
		&MsgApprove{},
		&MsgIncreaseAllowance{},
		&MsgDecreaseAllowance{},
		&MsgMint{},
		&MsgBurn{},
		&MsgBurnFrom{},
		&MsgTransferOwnership{},
		&MsgFreeze{},
		&MsgUnfreeze{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
