package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/x/genaccounts"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tendermint/tendermint/libs/common"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	srvconfig "github.com/cosmos/cosmos-sdk/server/config"
	"github.com/cosmos/cosmos-sdk/types/module"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"

	"github.com/konstellation/kn-sdk/types"
)

// get cmd to initialize all files for tendermint testnet and application
func TestnetCmd(
	ctx *server.Context,
	cdc *codec.Codec,
	mbm module.BasicManager,
	gus types.GenesisUpdaters,
	_ genutilcli.StakingMsgBuildingHelpers,
	genAccIterator genutiltypes.GenesisAccountsIterator,
) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "testnet",
		Short: "Initialize files for a Konstellation testnet",
		Long: `testnet will create "v" number of directories and populate each with
necessary files (private validator, genesis, config, etc.).

Note, strict routability for addresses is turned off in the config file.

Example:
	konstellation testnet --output-dir ./output --starting-ip-address 192.168.10.2
	`,
		RunE: func(_ *cobra.Command, _ []string) error {
			config := ctx.Config
			configFile := srvconfig.DefaultConfig()
			configFile.MinGasPrices = viper.GetString(server.FlagMinGasPrices)
			nodesInfoFile := viper.GetString(flagNodesInfoFile)

			nodeDaemonHomeName = viper.GetString(flagNodeDaemonHome)
			nodeCliHomeName = viper.GetString(flagNodeCliHome)

			outDir = viper.GetString(flagOutputDir)
			gentxsDir = filepath.Join(outDir, "gentxs")
			configDir = filepath.Join(outDir)

			chainID = viper.GetString(client.FlagChainID)
			if chainID == "" {
				chainID = fmt.Sprintf("test-chain-%v", common.RandStr(6))
			}

			if err := configClientNodes(config, configFile); err != nil {
				return err
			}

			nodes, err := configNodes(config, configFile, nodesInfoFile)
			if err != nil {
				return err
			}

			accs, err := genAccounts(nodes)
			if err != nil {
				return err
			}

			if err := initGenFiles(cdc, mbm, gus, nodes, accs, config); err != nil {
				return err
			}

			if err := genTxs(cdc, mbm, genAccIterator, nodes); err != nil {
				return err
			}

			if err := collectGenFiles(cdc, config, genaccounts.AppModuleBasic{}, nodes); err != nil {
				return err
			}

			fmt.Printf("Successfully initialized %d node directories\n", len(nodes))
			return nil
		},
	}

	cmd.Flags().StringP(flagOutputDir, "o", "./testnet",
		"Directory to store initialization data for the testnet",
	)
	cmd.Flags().String(flagNodeDirPrefix, "node",
		"Prefix the directory name for each node with (node results in node0, node1, ...)",
	)
	cmd.Flags().String(flagNodeDaemonHome, "konstellation",
		"Home directory of the node's daemon configuration",
	)
	cmd.Flags().String(flagNodeCliHome, "konstellationcli",
		"Home directory of the node's cli configuration",
	)
	cmd.Flags().String(flagNodesInfoFile, "./config/testnet.json",
		"Nodes configuration file",
	)
	cmd.Flags().String(flagStartingIPAddress, "testnode",
		"Starting IP address (testnode results in persistent peers list ID0@testnode-0:26656, ID1@testnode-1:26656, ...)")

	cmd.Flags().String(client.FlagChainID, "", "genesis file chain-id, if left blank will be randomly created")

	cmd.Flags().String(
		server.FlagMinGasPrices, fmt.Sprintf("0.000006%s", types.StakeDenom),
		"Minimum gas prices to accept for transactions; All fees in a tx must meet this minimum (e.g. 0.01apple,0.001darc)",
	)

	return cmd
}
