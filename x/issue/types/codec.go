package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgUnfreeze{}, "issue/Unfreeze", nil)

	cdc.RegisterConcrete(&MsgBurnFrom{}, "issue/BurnFrom", nil)

	cdc.RegisterConcrete(&MsgBurn{}, "issue/Burn", nil)

	cdc.RegisterConcrete(&MsgMint{}, "issue/Mint", nil)

	cdc.RegisterConcrete(&MsgDecreaseAllowance{}, "issue/DecreaseAllowance", nil)

	cdc.RegisterConcrete(&MsgIncreaseAllowance{}, "issue/IncreaseAllowance", nil)

	cdc.RegisterConcrete(&MsgApprove{}, "issue/Approve", nil)

	cdc.RegisterConcrete(&MsgTransferOwnership{}, "issue/TransferOwnership", nil)

	cdc.RegisterConcrete(&MsgTransferFrom{}, "issue/TransferFrom", nil)

	cdc.RegisterConcrete(&MsgTransfer{}, "issue/Transfer", nil)

	cdc.RegisterConcrete(&MsgFeatures{}, "issue/Features", nil)

	cdc.RegisterConcrete(&MsgEnableFeature{}, "issue/EnableFeature", nil)

	cdc.RegisterConcrete(&MsgDisableFeature{}, "issue/DisableFeature", nil)

	cdc.RegisterConcrete(&MsgDescription{}, "issue/Description", nil)

	cdc.RegisterConcrete(&MsgIssueCreate{}, "issue/IssueCreate", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnfreeze{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBurnFrom{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBurn{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMint{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDecreaseAllowance{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgIncreaseAllowance{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApprove{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferOwnership{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferFrom{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransfer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFeatures{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEnableFeature{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDisableFeature{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDescription{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgIssueCreate{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
