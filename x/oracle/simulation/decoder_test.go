package simulation_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/konstellation/konstellation/x/oracle/simulation"
	"github.com/konstellation/konstellation/x/oracle/types"
)

func TestDecodeStore(t *testing.T) {
	cdc := simapp.MakeTestEncodingConfig().Marshaler
	dec := simulation.NewDecodeStore(cdc)

	rate := types.ExchangeRate{
		Denom: "udarc",
		Rate:  1.2,
	}
	rand := rand.New(rand.NewSource(int64(1)))
	address := simulation.RandomAddress(rand)

	exchangeRate := types.NewMsgSetExchangeRate(&rate, address)

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{Key: types.ExchangeRateKey, Value: cdc.MustMarshalBinaryBare(&exchangeRate)},
			{Key: []byte{0x99}, Value: []byte{0x99}},
		},
	}
	tests := []struct {
		name        string
		expectedLog string
	}{
		{"ExchangeRate", fmt.Sprintf("%v\n%v", exchangeRate, exchangeRate)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
