package simulation_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/konstellation/konstellation/x/oracle/simulation"
)

func TestNewQuerier(t *testing.T) {
	rand := rand.New(rand.NewSource(int64(1)))
	address := simulation.RandomAddress(rand)
	require.NotZero(t, len(address))
}
