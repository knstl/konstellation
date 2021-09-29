package conf

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type ValNodeConfig struct {
	DirName   string
	DaemonDir string
	CliDir    string
}

type Description struct {
	Moniker         string `json:"moniker"`
	Identity        string `json:"identity"`
	Website         string `json:"website"`
	Details         string `json:"details"`
	SecurityContact string `json:"security_contact"`
}

type ValidatorKey struct {
	Address      string           `json:"address"`
	AccAddr      types.AccAddress `json:"-"`
	PublicKey    []byte           `json:"-"`
	CoinDelegate int64            `json:"coin_delegate"`
}

type ValidatorInfo struct {
	Name        string        `json:"name"`
	IP          string        `json:"ip"`
	Index       int           `json:"index"`
	Cors        string        `json:"cors"`
	Faucet      bool          `json:"faucet"`
	Key         *ValidatorKey `json:"key"`
	Description *Description  `json:"description"`
	Config      *NodeConfig   `json:"config"`
}

type Validator struct {
	ValNodeConfig ValNodeConfig
	Index         int
	Name          string
	ChainID       string
	Moniker       string
	ID            string
	GenFile       string
	GenAccount    *authTypes.GenesisAccount
	Memo          string
	Cors          string
	ValPubKey     cryptotypes.PubKey
	IP            string
	Key           *ValidatorKey
	Description   *Description
}
