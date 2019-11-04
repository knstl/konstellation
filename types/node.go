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

type Key struct {
	Name         string `json:"name"`
	Password     string `json:"password"`
	Mnemonic     string `json:"mnemonic"`
	CoinGenesis  int64  `json:"coin_genesis"`
	CoinDelegate int64  `json:"coin_delegate"`
}

type Description struct {
	Moniker  string `json:"moniker"`
	Identity string `json:"identity"`
	Website  string `json:"website"`
	Details  string `json:"details"`
}

type NodeInfo struct {
	Name        string      `json:"name"`
	IP          string      `json:"ip"`
	Index       int         `json:"index"`
	Cors        string      `json:"cors"`
	Faucet      bool        `json:"faucet"`
	Key         Key         `json:"key"`
	Description Description `json:"description"`
}

type Node struct {
	Config      NodeConfig
	Index       int
	ChainID     string
	Moniker     string
	ID          string
	GenFile     string
	GenAccount  *genaccounts.GenesisAccount
	Memo        string
	Cors        string
	ValPubKey   crypto.PubKey
	IP          string
	Key         Key
	Description Description
}
