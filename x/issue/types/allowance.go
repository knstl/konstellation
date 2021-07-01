package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Allowance struct {
	Amount  sdk.Int `json:"amount"`
	Spender string  `json:"spender"`
}

func NewAllowance(amount sdk.Coin, spender sdk.AccAddress) *Allowance {
	return &Allowance{amount.Amount, spender.String()}
}

func (a Allowance) String() string {
	return fmt.Sprintf(`%s:%s`, a.Spender, a.Amount)
}

type Allowances []*Allowance

func (as *Allowances) String() (str string) {
	for _, allowance := range *as {
		str += fmt.Sprintf(`%s:%s`, allowance.Spender, allowance.Amount)
	}
	return
}

func (as *Allowances) ContainsI(al *Allowance) int {
	for i, a := range *as {
		if a.Spender == al.Spender {
			return i
		}
	}

	return -1
}

func (as *Allowances) Contains(al *Allowance) bool {
	for _, a := range *as {
		if a.Spender == al.Spender {
			return true
		}
	}

	return false
}
