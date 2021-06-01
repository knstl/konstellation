package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/konstellation/konstellation/x/oracle/types"
)

// RandomAddress generates n random address
func GetAddress(r *rand.Rand) string {
	privkeySeed := make([]byte, 15)
	r.Read(privkeySeed)

	privKey := secp256k1.GenPrivKeyFromSecret(privkeySeed)
	return privKey.PubKey().Address().String()
}

func RandomizedGenState(simState *module.SimulationState) {
	addresses := []*types.AdminAddr{{Address: RandomAddress(simState.Rand)}}
	addressGenesis := types.NewGenesisState(addresses)

	bz, err := json.MarshalIndent(&addressGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated address parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(addressGenesis)
}
