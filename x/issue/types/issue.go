package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mitchellh/mapstructure"
	"math/big"
	"strings"
)

var (
	CoinMaxTotalSupply, _ = sdk.NewIntFromString("1000000000000000000000000000000000000")
)

const (
	InitLastId uint64 = 1
	//CoinIssueMaxId        uint64 = 999999999999
	Custom = "custom"
)

type IIssue interface {
	GetId() uint64
	SetId(uint64)

	GetDenom() string
	SetDenom(string)

	GetIssuer() sdk.AccAddress
	SetIssuer(sdk.AccAddress)

	GetOwner() sdk.AccAddress
	SetOwner(sdk.AccAddress)

	GetSymbol() string
	SetSymbol(string)

	GetDecimals() uint
	SetDecimals(uint)

	GetIssueTime() int
	SetIssueTime(int)

	GetTotalSupply() sdk.Int
	SetTotalSupply(sdk.Int)
	AddTotalSupply(sdk.Int)
	SubTotalSupply(sdk.Int)

	ToCoin() sdk.Coin
}

type CoinIssues []CoinIssue

//nolint
func (coinIssues CoinIssues) String() string {
	out := fmt.Sprintf("%-17s|%-44s|%-10s|%-6s|%-18s|%-8s|%s\n",
		"IssueID", "Owner", "Name", "Symbol", "TotalSupply", "Decimals", "IssueTime")
	for _, issue := range coinIssues {
		out += fmt.Sprintf("%-44s|%-10s|%-6s|%-18s|%-8d|%d\n", issue.GetOwner().String(), issue.Denom, issue.Symbol, issue.TotalSupply.String(), issue.Decimals, issue.IssueTime)
	}
	return strings.TrimSpace(out)
}

type CoinIssue struct {
	Id                 uint64         `json:"id"`
	Issuer             sdk.AccAddress `json:"issuer"`
	Owner              sdk.AccAddress `json:"owner"`
	Denom              string         `json:"denom"`
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

func (ci *CoinIssue) GetId() uint64 {
	return ci.Id
}

func (ci *CoinIssue) SetId(id uint64) {
	ci.Id = id
}

func (ci *CoinIssue) GetDenom() string {
	return ci.Denom
}

func (ci *CoinIssue) SetDenom(denom string) {
	ci.Denom = denom
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

func (ci *CoinIssue) GetDecimals() uint {
	return ci.Decimals
}

func (ci *CoinIssue) SetDecimals(decimals uint) {
	ci.Decimals = decimals
}

func (ci *CoinIssue) GetTotalSupply() sdk.Int {
	return ci.TotalSupply
}

func (ci *CoinIssue) SetTotalSupply(totalSupply sdk.Int) {
	ci.TotalSupply = totalSupply
}

func (ci *CoinIssue) AddTotalSupply(amount sdk.Int) {
	ci.TotalSupply = ci.TotalSupply.Add(amount)
}

func (ci *CoinIssue) SubTotalSupply(amount sdk.Int) {
	ci.TotalSupply = ci.TotalSupply.Sub(amount)
	if ci.TotalSupply.IsNegative() {
		ci.TotalSupply = sdk.ZeroInt()
	}
}

func (ci *CoinIssue) GetIssueTime() int64 {
	return ci.IssueTime
}

func (ci *CoinIssue) SetIssueTime(time int64) {
	ci.IssueTime = time
}

func (ci *CoinIssue) ToCoin() sdk.Coin {
	return sdk.NewCoin(ci.Denom, ci.TotalSupply)
}

func (ci *CoinIssue) ToCoins() sdk.Coins {
	return sdk.NewCoins(ci.ToCoin())
}

func getDecimalsInt(decimals uint) sdk.Int {
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	return sdk.NewIntFromBigInt(exp)
}

func (ci *CoinIssue) QuoDecimals(amount sdk.Int) sdk.Int {
	return amount.Quo(getDecimalsInt(ci.GetDecimals()))
}
