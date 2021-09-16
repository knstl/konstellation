package types

import (
	//"fmt"
	// this line is used by starport scaffolding # ibc/genesistype/import
	"bytes"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		/*
			// this line is used by starport scaffolding # ibc/genesistype/default
			// this line is used by starport scaffolding # genesis/types/default
			CoinIssueDenomsList:   []*CoinIssueDenoms{},
			CoinIssueListList:     []*CoinIssueList{},
			IssuesParamsList:      []*IssuesParams{},
			IssueParamsList:       []*IssueParams{},
			IssueFeaturesList:     []*IssueFeatures{},
			ParamsList:            []*Params{},
			IssuesList:            []*Issues{},
			CoinIssueDenomList:    []*CoinIssueDenom{},
			CoinIssueList:         []*CoinIssue{},
			AddressFreezeListList: []*AddressFreezeList{},
			AddressFreezeList:     []*AddressFreeze{},
			FreezeList:            []*Freeze{},
			CoinsList:             []*Coins{},
			AllowanceListList:     []*AllowanceList{},
			AllowanceList:         []*Allowance{},
			AddressList:           []*Address{},
		*/
		StartingIssueId: InitLastId,
		Issues: &Issues{
			Issues: []*CoinIssue{},
		},
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	/*
		// this line is used by starport scaffolding # ibc/genesistype/validate

		// this line is used by starport scaffolding # genesis/types/validate
		// Check for duplicated ID in coinIssueDenoms
		coinIssueDenomsIdMap := make(map[uint64]bool)

		for _, elem := range gs.CoinIssueDenomsList {
			if _, ok := coinIssueDenomsIdMap[elem.Id]; ok {
				return fmt.Errorf("duplicated id for coinIssueDenoms")
			}
			coinIssueDenomsIdMap[elem.Id] = true
		}
		// Check for duplicated ID in coinIssueList
		coinIssueListIdMap := make(map[uint64]bool)

		for _, elem := range gs.CoinIssueListList {
			if _, ok := coinIssueListIdMap[elem.Id]; ok {
				return fmt.Errorf("duplicated id for coinIssueList")
			}
			coinIssueListIdMap[elem.Id] = true
		}
		// Check for duplicated ID in issuesParams
		issuesParamsIdMap := make(map[uint64]bool)

		for _, elem := range gs.IssuesParamsList {
			if _, ok := issuesParamsIdMap[elem.Id]; ok {
				return fmt.Errorf("duplicated id for issuesParams")
			}
			issuesParamsIdMap[elem.Id] = true
		}
		// Check for duplicated ID in issueParams
		issueParamsIdMap := make(map[uint64]bool)

		for _, elem := range gs.IssueParamsList {
			if _, ok := issueParamsIdMap[elem.Id]; ok {
				return fmt.Errorf("duplicated id for issueParams")
			}
			issueParamsIdMap[elem.Id] = true
		}
		// Check for duplicated ID in issueFeatures
		issueFeaturesIdMap := make(map[uint64]bool)

		for _, elem := range gs.IssueFeaturesList {
			if _, ok := issueFeaturesIdMap[elem.Id]; ok {
				return fmt.Errorf("duplicated id for issueFeatures")
			}
			issueFeaturesIdMap[elem.Id] = true
		}
		// Check for duplicated ID in params
		paramsIdMap := make(map[uint64]bool)

		for _, elem := range gs.ParamsList {
			if _, ok := paramsIdMap[elem.Id]; ok {
				return fmt.Errorf("duplicated id for params")
			}
			paramsIdMap[elem.Id] = true
		}
		// Check for duplicated ID in issues
		issuesIdMap := make(map[uint64]bool)

		for _, elem := range gs.IssuesList {
			if _, ok := issuesIdMap[elem.Id]; ok {
				return fmt.Errorf("duplicated id for issues")
			}
			issuesIdMap[elem.Id] = true
		}
		// Check for duplicated ID in coinIssueDenom
		coinIssueDenomIdMap := make(map[uint64]bool)

		for _, elem := range gs.CoinIssueDenomList {
			if _, ok := coinIssueDenomIdMap[elem.Id]; ok {
				return fmt.Errorf("duplicated id for coinIssueDenom")
			}
			coinIssueDenomIdMap[elem.Id] = true
		}
		// Check for duplicated ID in coinIssue
		coinIssueIdMap := make(map[uint64]bool)

		for _, elem := range gs.CoinIssueList {
			if _, ok := coinIssueIdMap[elem.Id]; ok {
				return fmt.Errorf("duplicated id for coinIssue")
			}
			coinIssueIdMap[elem.Id] = true
		}
		// Check for duplicated ID in addressFreezeList
		addressFreezeListIdMap := make(map[uint64]bool)

		for _, elem := range gs.AddressFreezeListList {
			if _, ok := addressFreezeListIdMap[elem.Id]; ok {
				return fmt.Errorf("duplicated id for addressFreezeList")
			}
			addressFreezeListIdMap[elem.Id] = true
		}
		// Check for duplicated ID in addressFreeze
		addressFreezeIdMap := make(map[uint64]bool)

		for _, elem := range gs.AddressFreezeList {
			if _, ok := addressFreezeIdMap[elem.Id]; ok {
				return fmt.Errorf("duplicated id for addressFreeze")
			}
			addressFreezeIdMap[elem.Id] = true
		}
		// Check for duplicated ID in freeze
		freezeIdMap := make(map[uint64]bool)

		for _, elem := range gs.FreezeList {
			if _, ok := freezeIdMap[elem.Id]; ok {
				return fmt.Errorf("duplicated id for freeze")
			}
			freezeIdMap[elem.Id] = true
		}
		// Check for duplicated ID in coins
		coinsIdMap := make(map[uint64]bool)

		for _, elem := range gs.CoinsList {
			if _, ok := coinsIdMap[elem.Id]; ok {
				return fmt.Errorf("duplicated id for coins")
			}
			coinsIdMap[elem.Id] = true
		}
		// Check for duplicated ID in allowanceList
		allowanceListIdMap := make(map[uint64]bool)

		for _, elem := range gs.AllowanceListList {
			if _, ok := allowanceListIdMap[elem.Id]; ok {
				return fmt.Errorf("duplicated id for allowanceList")
			}
			allowanceListIdMap[elem.Id] = true
		}
		// Check for duplicated ID in allowance
		allowanceIdMap := make(map[uint64]bool)

		for _, elem := range gs.AllowanceList {
			if _, ok := allowanceIdMap[elem.Id]; ok {
				return fmt.Errorf("duplicated id for allowance")
			}
			allowanceIdMap[elem.Id] = true
		}
		// Check for duplicated ID in address
		addressIdMap := make(map[uint64]bool)

		for _, elem := range gs.AddressList {
			if _, ok := addressIdMap[elem.Id]; ok {
				return fmt.Errorf("duplicated id for address")
			}
			addressIdMap[elem.Id] = true
		}
	*/

	return nil
}

// NewGenesisState - Create a new genesis state
func NewGenesisState(startingIssueId uint64, params Params) GenesisState {
	return GenesisState{
		StartingIssueId: startingIssueId,
		Params:          params,
	}
}

// Returns if a GenesisState is empty or has data in it
func (gs GenesisState) IsEmpty() bool {
	emptyGenState := GenesisState{}
	return gs.Equal(emptyGenState)
}

func (gs GenesisState) Equal(gs2 GenesisState) bool {
	bz1 := ModuleCdc.MustMarshalBinaryLengthPrefixed(&gs)
	bz2 := ModuleCdc.MustMarshalBinaryLengthPrefixed(&gs2)
	return bytes.Equal(bz1, bz2)
}

// GetGenesisStateFromAppState returns x/auth GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc *codec.LegacyAmino, appState map[string]json.RawMessage) GenesisState {
	var genesisState GenesisState
	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return genesisState
}
