package types

import (
	"bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/params/subspace"
	"github.com/konstellation/kn-sdk/types"
)

// DefaultParamspace defines the default issue module parameter subspace
const DefaultParamspace = ModuleName

// Default parameter values
const ()

// Parameter keys
var (
	KeyIssueFee         = []byte("IssueFee")
	KeyMintFee          = []byte("MintFee")
	KeyFreezeFee        = []byte("FreezeFee")
	KeyUnFreezeFee      = []byte("UnfreezeFee")
	KeyBurnFee          = []byte("BurnFee")
	KeyBurnFromFee      = []byte("BurnFromFee")
	KeyTransferOwnerFee = []byte("TransferOwnerFee")
)

var _ subspace.ParamSet = &Params{}

// Params defines the parameters for the auth module.
type Params struct {
	IssueFee         sdk.Coin `json:"issue_fee"`
	MintFee          sdk.Coin `json:"mint_fee"`
	FreezeFee        sdk.Coin `json:"freeze_fee"`
	UnfreezeFee      sdk.Coin `json:"unfreeze_fee"`
	BurnFee          sdk.Coin `json:"burn_fee"`
	BurnFromFee      sdk.Coin `json:"burn_from_fee"`
	TransferOwnerFee sdk.Coin `json:"transfer_owner_fee"`
}

// NewParams creates a new Params instance
func NewParams(issueFee, mintFee, freezeFee, unfreezeFee, burnFee, burnFromFee, transferOwnerFee sdk.Coin) Params {
	return Params{
		IssueFee:         issueFee,
		MintFee:          mintFee,
		FreezeFee:        freezeFee,
		UnfreezeFee:      unfreezeFee,
		BurnFee:          burnFee,
		BurnFromFee:      burnFromFee,
		TransferOwnerFee: transferOwnerFee,
	}
}

// ParamKeyTable for auth module
func ParamKeyTable() subspace.KeyTable {
	return subspace.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
// pairs of auth module's parameters.
// nolint
func (p *Params) ParamSetPairs() subspace.ParamSetPairs {
	return params.ParamSetPairs{
		{Key: KeyIssueFee, Value: &p.IssueFee},
		{Key: KeyMintFee, Value: &p.MintFee},
		{Key: KeyFreezeFee, Value: &p.FreezeFee},
		{Key: KeyUnFreezeFee, Value: &p.UnfreezeFee},
		{Key: KeyBurnFee, Value: &p.BurnFee},
		{Key: KeyBurnFromFee, Value: &p.BurnFromFee},
		{Key: KeyTransferOwnerFee, Value: &p.TransferOwnerFee},
	}
}

// Equal returns a boolean determining if two Params types are identical.
func (p Params) Equal(p2 Params) bool {
	bz1 := ModuleCdc.MustMarshalBinaryLengthPrefixed(&p)
	bz2 := ModuleCdc.MustMarshalBinaryLengthPrefixed(&p2)
	return bytes.Equal(bz1, bz2)
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{
		IssueFee:         sdk.NewCoin(types.DefaultBondDenom, sdk.NewInt(200000000)),
		MintFee:          sdk.NewCoin(types.DefaultBondDenom, sdk.NewInt(100000000)),
		FreezeFee:        sdk.NewCoin(types.DefaultBondDenom, sdk.NewInt(200000000)),
		UnfreezeFee:      sdk.NewCoin(types.DefaultBondDenom, sdk.NewInt(200000000)),
		BurnFee:          sdk.NewCoin(types.DefaultBondDenom, sdk.NewInt(100000000)),
		BurnFromFee:      sdk.NewCoin(types.DefaultBondDenom, sdk.NewInt(100000000)),
		TransferOwnerFee: sdk.NewCoin(types.DefaultBondDenom, sdk.NewInt(200000000)),
	}
}

// String implements the stringer interface.
func (p Params) String() string {
	return fmt.Sprintf(`Params:
  IssueFee:			%s
  MintFee:			%s
  FreezeFee:			%s
  UnfreezeFee:			%s
  BurnFee:			%s
  BurnFromFee:			%s
  TransferOwnerFee:		%s`,
		p.IssueFee.String(),
		p.MintFee.String(),
		p.FreezeFee.String(),
		p.UnfreezeFee.String(),
		p.BurnFee.String(),
		p.BurnFromFee.String(),
		p.TransferOwnerFee.String())
}

// unmarshal the current staking params value from store key or panic
func MustUnmarshalParams(cdc *codec.Codec, value []byte) Params {
	ps, err := UnmarshalParams(cdc, value)
	if err != nil {
		panic(err)
	}
	return ps
}

// unmarshal the current staking params value from store key
func UnmarshalParams(cdc *codec.Codec, value []byte) (params Params, err error) {
	err = cdc.UnmarshalBinaryLengthPrefixed(value, &params)
	if err != nil {
		return
	}
	return
}

// Validate checks that the parameters have valid values.
func (p Params) Validate() error {
	if p.IssueFee.IsNegative() {
		return ErrInvalidIssueFee(p.IssueFee.String())
	}
	if p.MintFee.IsNegative() {
		return ErrInvalidMintFee(p.MintFee.String())
	}
	if p.BurnFee.IsNegative() {
		return ErrInvalidBurnFee(p.BurnFee.String())
	}
	if p.BurnFromFee.IsNegative() {
		return ErrInvalidBurnFromFee(p.BurnFromFee.String())
	}
	if p.FreezeFee.IsNegative() {
		return ErrInvalidFreezeFee(p.FreezeFee.String())
	}
	if p.UnfreezeFee.IsNegative() {
		return ErrInvalidUnfreezeFee(p.UnfreezeFee.String())
	}
	if p.TransferOwnerFee.IsNegative() {
		return ErrInvalidTransferOwnerFee(p.TransferOwnerFee.String())
	}
	return nil
}
