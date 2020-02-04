package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"

	"github.com/konstellation/kn-sdk/types"
)

// InitCmd returns a command that initializes all files needed for Tendermint
// and the respective application.
func GenTxCmd(ctx *server.Context, cdc *codec.Codec, mbm module.BasicManager, smbh genutilcli.StakingMsgBuildingHelpers,
	genAccIterator genutiltypes.GenesisAccountsIterator, defaultNodeHome, defaultCLIHome string) *cobra.Command { // nolint: golint
	ipDefault, _ := server.ExternalIP()
	_, _, _, flagAmount, _ := smbh.CreateValidatorMsgHelpers(ipDefault)

	viper.Set(flagAmount, sdk.TokensFromConsensusPower(types.DefaultConsensusPower).String()+types.DefaultBondDenom)

	return genutilcli.GenTxCmd(ctx, cdc, mbm, smbh, genAccIterator, defaultNodeHome, defaultCLIHome)
}
