package types

import (
	"github.com/tendermint/tendermint/crypto"

	"github.com/cosmos/cosmos-sdk/x/genaccounts"
)

type NodeConfig struct {
	DirName   string
	DaemonDir string
	CliDir    string
}

type NodeInfo struct {
	Name  string `json:"name"`
	IP    string `json:"ip"`
	Index int    `json:"index"`
	Cors  string `json:"cors"`
}

type Node struct {
	Config     NodeConfig
	Index      int
	ChainID    string
	Moniker    string
	ID         string
	GenFile    string
	GenAccount *genaccounts.GenesisAccount
	Pass       string
	Memo       string
	Cors       string
	ValPubKey  crypto.PubKey
	IP         string
}
