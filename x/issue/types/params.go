package types

import (
	"bytes"
	"strings"

	"github.com/cosmos/cosmos-sdk/x/params/subspace"
)

// DefaultParamspace defines the default issue module parameter subspace
const DefaultParamspace = ModuleName

// Default parameter values
const ()

// Parameter keys
var ()

var _ subspace.ParamSet = &Params{}

// Params defines the parameters for the auth module.
type Params struct {
}

// NewParams creates a new Params object
func NewParams() Params {
	return Params{}
}

// ParamKeyTable for auth module
func ParamKeyTable() subspace.KeyTable {
	return subspace.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
// pairs of auth module's parameters.
// nolint
func (p *Params) ParamSetPairs() subspace.ParamSetPairs {
	return subspace.ParamSetPairs{}
}

// Equal returns a boolean determining if two Params types are identical.
func (p Params) Equal(p2 Params) bool {
	bz1 := ModuleCdc.MustMarshalBinaryLengthPrefixed(&p)
	bz2 := ModuleCdc.MustMarshalBinaryLengthPrefixed(&p2)
	return bytes.Equal(bz1, bz2)
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{}
}

// String implements the stringer interface.
func (p Params) String() string {
	var sb strings.Builder
	sb.WriteString("Params: \n")
	return sb.String()
}

// Validate checks that the parameters have valid values.
func (p Params) Validate() error {
	return nil
}
