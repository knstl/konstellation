package types

import (
	//"fmt"
	// this line is used by starport scaffolding # ibc/genesistype/import
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

var (
	// TODO: change default actual address
	DefaultAllowedAddress = []byte("")
	DefaultExchangeRate   = []byte("")
)

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # ibc/genesistype/default
		// this line is used by starport scaffolding # genesis/types/default
		//ParamsList:       []*Params{},
		//AdminAddrList:    []*AdminAddr{},
		//ExchangeRateList: []*ExchangeRate{},
		AllowedAddresses: nil,
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated ID in params
	/*
		paramsIdMap := make(map[uint64]bool)

		for _, elem := range gs.ParamsList {
			if _, ok := paramsIdMap[elem.Id]; ok {
				return fmt.Errorf("duplicated id for params")
			}
			paramsIdMap[elem.Id] = true
		}
		// Check for duplicated ID in adminAddr
		adminAddrIdMap := make(map[uint64]bool)

		for _, elem := range gs.AdminAddrList {
			if _, ok := adminAddrIdMap[elem.Id]; ok {
				return fmt.Errorf("duplicated id for adminAddr")
			}
			adminAddrIdMap[elem.Id] = true
		}
		// Check for duplicated ID in exchangeRate
		exchangeRateIdMap := make(map[uint64]bool)

		for _, elem := range gs.ExchangeRateList {
			if _, ok := exchangeRateIdMap[elem.Id]; ok {
				return fmt.Errorf("duplicated id for exchangeRate")
			}
			exchangeRateIdMap[elem.Id] = true
		}
	*/
	for _, address := range gs.AllowedAddresses {
		if address == nil {
			// todo change
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "address is empty string")
		}
	}

	return nil
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState(allowedAddresses []*AdminAddr) *GenesisState {
	return &GenesisState{
		AllowedAddresses: allowedAddresses,
	}
}

// ValidateGenesis performs basic validation of supply genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data *GenesisState) error {
	return data.Validate()
}
