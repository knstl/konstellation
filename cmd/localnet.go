package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/common"
	tmtypes "github.com/tendermint/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	srvconfig "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/konstellation/konstellation/common/utils"
	"github.com/konstellation/konstellation/types"
)

var (
	flagNodeDirPrefix     = "node-dir-prefix"
	flagNumValidators     = "v"
	flagOutputDir         = "output-dir"
	flagNodeDaemonHome    = "node-daemon-home"
	flagNodeCliHome       = "node-cli-home"
	flagStartingIPAddress = "starting-ip-address"

	outDir             = ""
	gentxsDir          = ""
	chainID            = ""
	nodeDaemonHomeName = ""
	nodeCliHomeName    = ""
	numValidators      = 0
)

const nodeDirPerm = 0755

// get cmd to initialize all files for tendermint localnet and application
func LocalnetCmd(
	ctx *server.Context,
	cdc *codec.Codec,
	mbm module.BasicManager,
	gus types.GenesisUpdaters,
	_ genutilcli.StakingMsgBuildingHelpers,
	genAccIterator genutiltypes.GenesisAccountsIterator,
) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "localnet",
		Short: "Initialize files for a Konstellation localnet",
		Long: `localnet will create "v" number of directories and populate each with
necessary files (private validator, genesis, config, etc.).

Note, strict routability for addresses is turned off in the config file.

Example:
	konstellation localnet --v 4 --output-dir ./output --starting-ip-address 192.168.10.2
	`,
		RunE: func(_ *cobra.Command, _ []string) error {
			config := ctx.Config
			configFile := srvconfig.DefaultConfig()
			configFile.MinGasPrices = viper.GetString(server.FlagMinGasPrices)

			nodeDaemonHomeName = viper.GetString(flagNodeDaemonHome)
			nodeCliHomeName = viper.GetString(flagNodeCliHome)
			numValidators = viper.GetInt(flagNumValidators)

			outDir = viper.GetString(flagOutputDir)
			gentxsDir = filepath.Join(outDir, "gentxs")

			chainID = viper.GetString(client.FlagChainID)
			if chainID == "" {
				chainID = fmt.Sprintf("test-chain-%v", common.RandStr(6))
			}

			nodes, err := configLocalNodes(config, configFile)
			if err != nil {
				return err
			}

			accs, err := genAccounts(nodes)
			if err != nil {
				return err
			}

			if err := initGenFiles(cdc, mbm, gus, nodes, accs); err != nil {
				return err
			}

			if err := genTxs(cdc, mbm, genAccIterator, nodes); err != nil {
				return err
			}

			if err := collectGenFiles(cdc, config, genaccounts.AppModuleBasic{}, nodes); err != nil {
				return err
			}

			fmt.Printf("Successfully initialized %d node directories\n", numValidators)
			return nil
		},
	}

	cmd.Flags().Int(flagNumValidators, 4,
		"Number of validators to initialize the localnet with",
	)
	cmd.Flags().StringP(flagOutputDir, "o", "./localnet",
		"Directory to store initialization data for the localnet",
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
	cmd.Flags().String(flagStartingIPAddress, "testnode",
		"Starting IP address (testnode results in persistent peers list ID0@testnode-0:26656, ID1@testnode-1:26656, ...)")

	cmd.Flags().String(client.FlagChainID, "", "genesis file chain-id, if left blank will be randomly created")

	cmd.Flags().String(
		server.FlagMinGasPrices, fmt.Sprintf("0.000006%s", types.StakeDenom),
		"Minimum gas prices to accept for transactions; All fees in a tx must meet this minimum (e.g. 0.01apple,0.001darc)",
	)

	return cmd
}

func configNode(config *cfg.Config, configFile *srvconfig.Config, nodeDirName string, ip string, index int) (node *types.Node, err error) {
	nodeDir := filepath.Join(outDir, nodeDirName, nodeDaemonHomeName)
	clientDir := filepath.Join(outDir, nodeDirName, nodeCliHomeName)
	nodeConfig := types.NodeConfig{
		DirName:   nodeDirName,
		DaemonDir: nodeDir,
		CliDir:    clientDir,
	}

	config.SetRoot(nodeDir)
	config.Moniker = nodeDirName

	err = os.MkdirAll(clientDir, nodeDirPerm)
	if err != nil {
		_ = os.RemoveAll(outDir)
		return nil, err
	}

	err = os.MkdirAll(filepath.Join(nodeDir, "config"), nodeDirPerm)
	if err != nil {
		_ = os.RemoveAll(outDir)
		return nil, err
	}

	configFilePath := filepath.Join(nodeDir, "config/konstellation.toml")
	srvconfig.WriteConfigFile(configFilePath, configFile)

	nodeID, valPubKey, err := genutil.InitializeNodeValidatorFiles(config)
	if err != nil {
		_ = os.RemoveAll(outDir)
		return nil, err
	}

	memo := fmt.Sprintf("%s@%s:26656", nodeID, ip)

	return &types.Node{
		Index:     index,
		Moniker:   nodeDirName,
		Config:    nodeConfig,
		GenFile:   config.GenesisFile(),
		Memo:      memo,
		ID:        nodeID,
		ChainID:   chainID,
		ValPubKey: valPubKey,
	}, nil
}

func configLocalNodes(config *cfg.Config, configFile *srvconfig.Config) (nodes []*types.Node, err error) {
	for i := 0; i < numValidators; i++ {
		nodeDirName := fmt.Sprintf("%s%d", viper.GetString(flagNodeDirPrefix), i)

		ip := fmt.Sprintf("%s-%d", viper.GetString(flagStartingIPAddress), i)

		node, err := configNode(config, configFile, nodeDirName, ip, i)
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, node)
	}
	return
}

func genAccounts(nodes []*types.Node) (accs []*genaccounts.GenesisAccount, err error) {
	for _, node := range nodes {
		// 	"Password for account '%s' (default %s):", nodeDirName, "12345678",
		// )
		// keyPass, err := client.GetPassword(prompt, buf)
		node.Pass = "12345678"

		addr, secret, err := server.GenerateSaveCoinKey(node.Config.CliDir, node.Config.DirName, node.Pass, true)
		if err != nil {
			_ = os.RemoveAll(outDir)
			return nil, err
		}

		info := map[string]string{"secret": secret}
		cliPrint, err := json.Marshal(info)
		if err != nil {
			return nil, err
		}

		err = utils.WriteFile(fmt.Sprintf("%v.json", "key_seed"), node.Config.CliDir, cliPrint)
		if err != nil {
			return nil, err
		}

		genacc := &genaccounts.GenesisAccount{
			Address: addr,
			Coins: sdk.NewCoins(
				sdk.NewCoin(types.StakeDenom, sdk.TokensFromConsensusPower(types.DefaultConsensusPower)),
			),
		}

		accs = append(accs, genacc)
		node.GenAccount = genacc
	}

	return
}

func initGenFiles(cdc *codec.Codec, mbm module.BasicManager, gus types.GenesisUpdaters, nodes []*types.Node, accs []*genaccounts.GenesisAccount) error {
	appGenState := mbm.DefaultGenesis()

	appGenState[genaccounts.ModuleName] = cdc.MustMarshalJSON(accs)

	// Update default genesis
	for _, gu := range gus {
		gu.UpdateGenesis(cdc, appGenState)
	}

	if err := mbm.ValidateGenesis(appGenState); err != nil {
		return fmt.Errorf("error validating genesis: %s", err.Error())
	}

	appState := cdc.MustMarshalJSON(appGenState)

	genDoc := &tmtypes.GenesisDoc{}
	genDoc.ChainID = chainID
	genDoc.Validators = nil
	genDoc.AppState = appState

	// generate empty genesis files for each validator and save
	for _, node := range nodes {
		if err := genutil.ExportGenesisFile(genDoc, node.GenFile); err != nil {
			return err
		}

		toPrint := utils.NewPrintInfo(node.Moniker, chainID, node.ID, "", appState)
		if err := utils.DisplayInfo(cdc, toPrint); err != nil {
			return err
		}
	}

	return nil
}

func genTxs(
	cdc *codec.Codec,
	mbm module.BasicManager,
	genAccIterator genutiltypes.GenesisAccountsIterator,
	nodes []*types.Node,
) error {
	for _, node := range nodes {
		genDoc, err := tmtypes.GenesisDocFromFile(node.GenFile)
		if err != nil {
			return err
		}

		var genesisState map[string]json.RawMessage
		if err = cdc.UnmarshalJSON(genDoc.AppState, &genesisState); err != nil {
			return err
		}

		if err = mbm.ValidateGenesis(genesisState); err != nil {
			return err
		}

		kb, err := client.NewKeyBaseFromDir(node.Config.CliDir)
		if err != nil {
			return err
		}

		key, err := kb.Get(node.Moniker)
		if err != nil {
			return err
		}

		c := sdk.NewCoin(types.StakeDenom, sdk.TokensFromConsensusPower(100))
		coins := sdk.NewCoins(c)
		err = genutil.ValidateAccountInGenesis(genesisState, genAccIterator, key.GetAddress(), coins, cdc)
		if err != nil {
			return err
		}

		msg := staking.NewMsgCreateValidator(
			sdk.ValAddress(node.GenAccount.Address),
			node.ValPubKey,
			c,
			staking.NewDescription(node.Moniker, "", "", ""),
			staking.NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec()),
			sdk.OneInt(),
		)

		tx := auth.NewStdTx([]sdk.Msg{msg}, auth.StdFee{}, []auth.StdSignature{}, node.Memo)
		txBldr := auth.NewTxBuilderFromCLI().WithChainID(chainID).WithMemo(node.Memo).WithKeybase(kb)
		signedTx, err := txBldr.SignStdTx(node.Moniker, node.Pass, tx, false)
		if err != nil {
			_ = os.RemoveAll(outDir)
			return err
		}

		txBytes, err := cdc.MarshalJSON(signedTx)
		if err != nil {
			_ = os.RemoveAll(outDir)
			return err
		}

		err = utils.WriteFile(fmt.Sprintf("%v.json", node.Moniker), gentxsDir, txBytes)
		if err != nil {
			_ = os.RemoveAll(outDir)
			return err
		}
	}

	return nil
}

func collectGenFiles(
	cdc *codec.Codec,
	config *cfg.Config,
	genAccIterator genutiltypes.GenesisAccountsIterator,
	nodes []*types.Node,
) error {

	var appState json.RawMessage
	genTime := tmtime.Now()

	for _, node := range nodes {
		config.SetRoot(node.Config.DaemonDir)
		config.Moniker = node.Moniker
		initCfg := genutil.NewInitConfig(chainID, gentxsDir, node.Moniker, node.ID, node.ValPubKey)

		genDoc, err := tmtypes.GenesisDocFromFile(node.GenFile)
		if err != nil {
			return err
		}

		nodeAppState, err := genutil.GenAppStateFromConfig(cdc, config, initCfg, *genDoc, genAccIterator)
		if err != nil {
			return err
		}

		if appState == nil {
			// set the canonical application state (they should not differ)
			appState = nodeAppState
		}

		genFile := config.GenesisFile()

		// overwrite each validator's genesis file to have a canonical genesis time
		err = genutil.ExportGenesisFileWithTime(genFile, chainID, nil, appState, genTime)
		if err != nil {
			return err
		}
	}

	return nil
}
