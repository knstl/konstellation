package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/pkg/errors"
)

// Default parameter namespace
const (
	DefaultParamspace = ModuleName
	// TODO: Define your default parameters
)

// Parameter store keys
var (
	DefaultAllowedAddress = []byte("")
)

// ParamKeyTable for nameservice module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// Params - used for initializing default parameter for nameservice at genesis
type Params struct {
	AllowedAddress string `json:"allowed_address"`
}

// NewParams creates a new Params object
func NewParams(address string) Params {
	return Params{
		AllowedAddress: address,
	}
}

// String implements the stringer interface for Params
func (p Params) String() string {
	return fmt.Sprintf("%s", p.AllowedAddress)
}

// ParamSetPairs - Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		// TODO: Pair your key with the param
		paramtypes.NewParamSetPair(ParamStoreKeyAllowedAddress, &p.AllowedAddress, validateAllowedAddress),
	}
}

func validateAllowedAddress(i interface{}) error {
	_, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}

// ValidateBasic performs basic validation on wasm parameters
func (p Params) ValidateBasic() error {
	if err := validateAllowedAddress(p.AllowedAddress); err != nil {
		return errors.Wrap(err, "allowed address")
	}
	return nil
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params {
	return NewParams(string(DefaultAllowedAddress))
}
