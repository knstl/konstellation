package types

import (
	"bytes"
	"fmt"
	"github.com/konstellation/konstellation/const"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
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

var _ paramtypes.ParamSet = (*Params)(nil)

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
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
// pairs of auth module's parameters.
// nolint
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyIssueFee, &p.IssueFee, validateFee),
		paramtypes.NewParamSetPair(KeyMintFee, &p.MintFee, validateFee),
		paramtypes.NewParamSetPair(KeyFreezeFee, &p.FreezeFee, validateFee),
		paramtypes.NewParamSetPair(KeyUnFreezeFee, &p.UnfreezeFee, validateFee),
		paramtypes.NewParamSetPair(KeyBurnFee, &p.BurnFee, validateFee),
		paramtypes.NewParamSetPair(KeyBurnFromFee, &p.BurnFromFee, validateFee),
		paramtypes.NewParamSetPair(KeyTransferOwnerFee, &p.TransferOwnerFee, validateFee),
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
		IssueFee:         sdk.NewCoin(_const.DefaultBondDenom, sdk.NewInt(200000000)),
		MintFee:          sdk.NewCoin(_const.DefaultBondDenom, sdk.NewInt(100000000)),
		FreezeFee:        sdk.NewCoin(_const.DefaultBondDenom, sdk.NewInt(200000000)),
		UnfreezeFee:      sdk.NewCoin(_const.DefaultBondDenom, sdk.NewInt(200000000)),
		BurnFee:          sdk.NewCoin(_const.DefaultBondDenom, sdk.NewInt(100000000)),
		BurnFromFee:      sdk.NewCoin(_const.DefaultBondDenom, sdk.NewInt(100000000)),
		TransferOwnerFee: sdk.NewCoin(_const.DefaultBondDenom, sdk.NewInt(200000000)),
	}
}

// unmarshal the current staking params value from store key or panic
func MustUnmarshalParams(cdc *codec.LegacyAmino, value []byte) Params {
	ps, err := UnmarshalParams(cdc, value)
	if err != nil {
		panic(err)
	}
	return ps
}

// unmarshal the current staking params value from store key
func UnmarshalParams(cdc *codec.LegacyAmino, value []byte) (params Params, err error) {
	err = cdc.UnmarshalBinaryLengthPrefixed(value, &params)
	if err != nil {
		return
	}
	return
}

// Validate checks that the parameters have valid values.
func (p Params) Validate() error {
	if err := validateFee(p.IssueFee); err != nil {
		return ErrInvalidIssueFee(p.IssueFee.String())
	}
	if err := validateFee(p.MintFee); err != nil {
		return ErrInvalidMintFee(p.MintFee.String())
	}
	if err := validateFee(p.BurnFee); err != nil {
		return ErrInvalidBurnFee(p.BurnFee.String())
	}
	if err := validateFee(p.BurnFromFee); err != nil {
		return ErrInvalidBurnFromFee(p.BurnFromFee.String())
	}
	if err := validateFee(p.FreezeFee); err != nil {
		return ErrInvalidFreezeFee(p.FreezeFee.String())
	}
	if err := validateFee(p.UnfreezeFee); err != nil {
		return ErrInvalidUnfreezeFee(p.UnfreezeFee.String())
	}
	if err := validateFee(p.TransferOwnerFee); err != nil {
		return ErrInvalidTransferOwnerFee(p.TransferOwnerFee.String())
	}

	return nil
}

func validateFee(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("fee must be not negative: %v", v)
	}

	return nil
}
