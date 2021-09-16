package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/konstellation/konstellation/x/oracle/types"
)

func NewDecodeStore(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key, []byte(types.ExchangeRateKey)):
			exchangeRateA := types.MsgSetExchangeRate{}
			exchangeRateB := types.MsgSetExchangeRate{}
			cdc.MustUnmarshalBinaryBare(kvA.Value, &exchangeRateA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &exchangeRateB)
			return fmt.Sprintf("%v\n%v", exchangeRateA, exchangeRateB)
		default:
			panic(fmt.Sprintf("invalid exchange rate key %X", kvA.Key))
		}
	}
}
