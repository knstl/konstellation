package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

// RandomAddress generates n random address
func RandomAddress(r *rand.Rand) string {
	privkeySeed := make([]byte, 15)
	r.Read(privkeySeed)

	privKey := secp256k1.GenPrivKeyFromSecret(privkeySeed)
	return privKey.PubKey().Address().String()
}
