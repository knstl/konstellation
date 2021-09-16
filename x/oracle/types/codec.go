package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgSetAdminAddr{}, "oracle/SetAdminAddr", nil)

	cdc.RegisterConcrete(&MsgDeleteExchangeRates{}, "oracle/DeleteExchangeRates", nil)

	cdc.RegisterConcrete(&MsgDeleteExchangeRate{}, "oracle/DeleteExchangeRate", nil)

	cdc.RegisterConcrete(&MsgSetExchangeRates{}, "oracle/SetExchangeRates", nil)

	cdc.RegisterConcrete(&MsgSetExchangeRate{}, "oracle/SetExchangeRate", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetAdminAddr{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteExchangeRates{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDeleteExchangeRate{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetExchangeRates{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetExchangeRate{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	//	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
