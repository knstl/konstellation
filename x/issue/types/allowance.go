package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewAllowance(amount sdk.Coin, spender sdk.AccAddress) *Allowance {
	return &Allowance{Amount: amount, Spender: spender.String()}
}

type Allowances []Allowance

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
