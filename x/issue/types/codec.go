package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc auth module wide codec
var ModuleCdc = codec.New()

// RegisterCodec registers concrete types on the codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgIssue{}, "issue/MsgIssue", nil)
	//cdc.RegisterConcrete(bank.MsgSend{}, "issue/MsgMint", nil)
	//cdc.RegisterInterface((*exported.GenesisAccount)(nil), nil)
	//cdc.RegisterInterface((*exported.Account)(nil), nil)
	//cdc.RegisterConcrete(&BaseAccount{}, "cosmos-sdk/Account", nil)
	//cdc.RegisterConcrete(StdTx{}, "cosmos-sdk/StdTx", nil)

	cdc.RegisterInterface((*IIssue)(nil), nil)
	cdc.RegisterConcrete(&CoinIssue{}, "issue/CoinIssue", nil)
}

// RegisterAccountTypeCodec registers an external account type defined in
// another module for the internal ModuleCdc.
//func RegisterAccountTypeCodec(o interface{}, name string) {
//	ModuleCdc.RegisterConcrete(o, name, nil)
//}

func init() {
	RegisterCodec(ModuleCdc)
	//codec.RegisterCrypto(ModuleCdc)
}
