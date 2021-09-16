package types

import (
	"fmt"
	"math/big"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mitchellh/mapstructure"
)

var (
	CoinMaxTotalSupply, _ = sdk.NewIntFromString("1000000000000000000000000000000000000")
)

const (
	InitLastId uint64 = 1
	//CoinIssueMaxId        uint64 = 999999999999
	Custom = "custom"
)

type Issue interface {
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

type CoinIssues []*CoinIssue

//nolint
func (coinIssues CoinIssues) String() string {
	out := fmt.Sprintf("%-17s|%-44s|%-10s|%-6s|%-18s|%-8s|%s\n",
		"IssueID", "Owner", "Name", "Symbol", "TotalSupply", "Decimals", "IssueTime")
	for _, issue := range coinIssues {
		out += fmt.Sprintf("%-44s|%-10s|%-6s|%-18s|%-8d|%d\n", issue.GetOwner(), issue.Denom, issue.Symbol, issue.TotalSupply.String(), issue.Decimals, issue.IssueTime)
	}
	return strings.TrimSpace(out)
}

func NewCoinIssue(owner, issuer sdk.AccAddress, params *IssueParams) *CoinIssue {
	var ci CoinIssue
	_ = mapstructure.Decode(params, &ci)
	ci.SetFeatures(params.IssueFeatures)
	ci.Owner = owner.String()
	ci.Issuer = issuer.String()
	ci.Symbol = strings.ToUpper(ci.Symbol)
	ci.TotalSupply = params.TotalSupply

	return &ci
}

func (ci *CoinIssue) SetId(id uint64) {
	ci.Id = id
}

func (ci *CoinIssue) SetDenom(denom string) {
	ci.Denom = denom
}

func (ci *CoinIssue) SetIssuer(issuer sdk.AccAddress) {
	ci.Issuer = issuer.String()
}

func (ci *CoinIssue) SetOwner(owner sdk.AccAddress) {
	ci.Owner = owner.String()
}

func (ci *CoinIssue) SetSymbol(symbol string) {
	ci.Symbol = symbol
}

func (ci *CoinIssue) SetDecimals(decimals uint) {
	ci.Decimals = int32(decimals)
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

func (ci *CoinIssue) SetIssueTime(time int64) {
	ci.IssueTime = int32(time)
}

func (ci *CoinIssue) ToCoin() sdk.Coin {
	return sdk.NewCoin(ci.Denom, ci.TotalSupply)
}

func (ci *CoinIssue) ToCoins() sdk.Coins {
	return sdk.NewCoins(ci.ToCoin())
}

func (ci *CoinIssue) SetFeatures(features *IssueFeatures) {
	ci.BurnOwnerDisabled = features.BurnOwnerDisabled
	ci.BurnHolderDisabled = features.BurnHolderDisabled
	ci.BurnFromDisabled = features.BurnFromDisabled
	ci.MintDisabled = features.MintDisabled
	ci.FreezeDisabled = features.FreezeDisabled
}

func getDecimalsInt(decimals uint) sdk.Int {
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	return sdk.NewIntFromBigInt(exp)
}

func (ci *CoinIssue) QuoDecimals(amount sdk.Int) sdk.Int {
	return amount.Quo(getDecimalsInt(uint(ci.GetDecimals())))
}

/*
func (ci *CoinIssue) String() string {
	out := fmt.Sprintf("%-17s|%-44s|%-10s|%-6s|%-18s|%-8s|%s\n",
		"IssueID", "Owner", "Name", "Symbol", "TotalSupply", "Decimals", "IssueTime")
	out += fmt.Sprintf("%-44s|%-10s|%-6s|%-18s|%-8d|%d\n", ci.GetOwner().String(), ci.Denom, ci.Symbol, ci.TotalSupply.String(), ci.Decimals, ci.IssueTime)
	return strings.TrimSpace(out)
}
*/
