package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mitchellh/mapstructure"
	"strings"
)

var (
	CoinMaxTotalSupply, _        = sdk.NewIntFromString("1000000000000000000000000000000000000")
	CoinIssueMaxId        uint64 = 999999999999
	CoinIssueMinId        uint64 = 100000000000
)

const (
	IDPreStr = "coin"
	Custom   = "custom"
)

type IssueParams struct {
	Name               string  `json:"name"`
	Symbol             string  `json:"symbol"`
	TotalSupply        sdk.Int `json:"total_supply"`
	Decimals           uint    `json:"decimals"`
	Description        string  `json:"description"`
	BurnOwnerDisabled  bool    `json:"burn_owner_disabled"`
	BurnHolderDisabled bool    `json:"burn_holder_disabled"`
	BurnFromDisabled   bool    `json:"burn_from_disabled"`
	MintingFinished    bool    `json:"minting_finished"`
	FreezeDisabled     bool    `json:"freeze_disabled"`
}

func NewIssueParams(data interface{}) (*IssueParams, error) {
	var issue IssueParams
	err := mapstructure.Decode(data, &issue)
	return &issue, err
}

type IIssue interface {
	GetIssueId() string
	SetIssueId(string)

	GetIssuer() sdk.AccAddress
	SetIssuer(sdk.AccAddress)

	GetOwner() sdk.AccAddress
	SetOwner(sdk.AccAddress)

	GetSymbol() string
	SetSymbol(string)

	ToCoin() sdk.Coin
}

type CoinIssue struct {
	IssueId            string         `json:"issue_id"`
	Issuer             sdk.AccAddress `json:"issuer"`
	Owner              sdk.AccAddress `json:"owner"`
	Name               string         `json:"name"`
	Symbol             string         `json:"symbol"`
	TotalSupply        sdk.Int        `json:"total_supply"`
	Decimals           uint           `json:"decimals"`
	Description        string         `json:"description"`
	IssueTime          int64          `json:"issue_time"`
	BurnOwnerDisabled  bool           `json:"burn_owner_disabled"`
	BurnHolderDisabled bool           `json:"burn_holder_disabled"`
	BurnFromDisabled   bool           `json:"burn_from_disabled"`
	FreezeDisabled     bool           `json:"freeze_disabled"`
	MintingFinished    bool           `json:"minting_finished"`
}

func NewCoinIssue(owner, issuer sdk.AccAddress, params *IssueParams) *CoinIssue {
	var ci CoinIssue
	_ = mapstructure.Decode(params, &ci)
	ci.Owner = owner
	ci.Issuer = issuer
	ci.Symbol = strings.ToUpper(ci.Symbol)
	ci.TotalSupply = params.TotalSupply

	return &ci
}

func (ci *CoinIssue) GetIssueId() string {
	return ci.IssueId
}

func (ci *CoinIssue) SetIssueId(issueId string) {
	ci.IssueId = issueId
}

func (ci *CoinIssue) GetIssuer() sdk.AccAddress {
	return ci.Issuer
}

func (ci *CoinIssue) SetIssuer(issuer sdk.AccAddress) {
	ci.Issuer = issuer
}

func (ci *CoinIssue) GetOwner() sdk.AccAddress {
	return ci.Owner
}

func (ci *CoinIssue) SetOwner(owner sdk.AccAddress) {
	ci.Owner = owner
}

func (ci *CoinIssue) GetSymbol() string {
	return ci.Symbol
}

func (ci *CoinIssue) SetSymbol(symbol string) {
	ci.Symbol = symbol
}

func (ci *CoinIssue) ToCoin() sdk.Coin {
	return sdk.NewCoin(ci.IssueId, ci.TotalSupply)
}

func (ci *CoinIssue) ToCoins() sdk.Coins {
	return sdk.NewCoins(ci.ToCoin())
}
