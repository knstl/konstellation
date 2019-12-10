package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Allowance struct {
	Amount sdk.Int `json:"amount"`
}

func NewAllowance(amount sdk.Coin) Allowance {
	return Allowance{amount.Amount}
}

func (a Allowance) String() string {
	return fmt.Sprintf(`Amount:%s`, a.Amount)
}
