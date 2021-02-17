package cmd

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/module"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/spf13/cobra"
)

// InitCmd returns a command that initializes all files needed for Tendermint
// and the respective application.
func GenTxCmd(mbm module.BasicManager, txEncCfg client.TxEncodingConfig, genBalIterator genutiltypes.GenesisBalancesIterator, defaultNodeHome string) *cobra.Command { // nolint: golint
	//_, _, _, flagAmount, _ := smbh.CreateValidatorMsgHelpers(ipDefault)
	// TODO flagAMount
	//viper.Set(flagAmount, sdk.TokensFromConsensusPower(types.DefaultConsensusPower).String()+types.DefaultBondDenom)

	return genutilcli.GenTxCmd(mbm, txEncCfg, genBalIterator, defaultNodeHome)
}
