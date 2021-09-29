package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/imdario/mergo"
	"github.com/konstellation/konstellation/app"
	conf "github.com/konstellation/konstellation/cmd/cosmodrome/pkg/chainconf"
	"github.com/konstellation/konstellation/cmd/cosmodrome/types"
	"github.com/konstellation/konstellation/const"
	"github.com/konstellation/konstellation/crypto/keybase"
	tmtypes "github.com/tendermint/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"
	"sort"
	"strings"
	"time"

	"io"
	"os"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/server"
	srvconfig "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/konstellation/konstellation/common/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	tmconfig "github.com/tendermint/tendermint/config"
	"path/filepath"
)

const nodeDirPerm = 0755

var (
	flagOutputDir         = "output-dir"
	flagNodeDaemonHome    = "node-daemon-home"
	flagNodeCliHome       = "node-cli-home"
	flagNumValidators     = "v"
	flagKeyStorageFile    = "key-storage"
	flagNetConfigFile     = "net-config"
	flagNodeDirPrefix     = "node-dir-prefix"
	flagStartingIPAddress = "starting-ip-address"

	outDir             = ""
	gentxsDir          = ""
	configDir          = ""
	nodeDaemonHomeName = ""
	nodeCliHomeName    = ""
)

func GenNetCmd(
	mbm module.BasicManager,
	genBalIterator banktypes.GenesisBalancesIterator,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "generate-network",
		Aliases: []string{"gen-net", "gn"},
		Short:   "Initialize files for a Konstellation network",
		Long: `Command will create "v" number of directories and populate each with
necessary files (private validator, genesis, config, etc.).

Note, strict routability for addresses is turned off in the config file.

Example:
	cosmodrome gn --chain-id darchub -n ./config/localnet.json  -o ./localnet
	cosmodrome generate-network --chain-id darchub --net-config ./config/testnet.json  --output-dir ./testnet
	`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			outputDir, _ := cmd.Flags().GetString(flagOutputDir)
			keyringBackend, _ := cmd.Flags().GetString(flags.FlagKeyringBackend)
			chainID, _ := cmd.Flags().GetString(flags.FlagChainID)
			minGasPrices, _ := cmd.Flags().GetString(server.FlagMinGasPrices)
			nodeDirPrefix, _ := cmd.Flags().GetString(flagNodeDirPrefix)
			nodeDaemonHome, _ := cmd.Flags().GetString(flagNodeDaemonHome)
			startingIPAddress, _ := cmd.Flags().GetString(flagStartingIPAddress)
			numValidators, _ := cmd.Flags().GetInt(flagNumValidators)
			algo, _ := cmd.Flags().GetString(flags.FlagKeyAlgorithm)

			// TODO keyringBackend - test for localnet, file - for testnet

			return generateNetwork(clientCtx, cmd, config, mbm, genBalIterator, outputDir, chainID, minGasPrices,
				nodeDirPrefix, nodeDaemonHome, startingIPAddress, keyringBackend, algo, numValidators)
		},
	}

	cmd.Flags().Int(flagNumValidators, 4, "Number of validators to initialize the testnet with")
	cmd.Flags().String(flagNodeDirPrefix, "node", "Prefix the directory name for each node with (node results in node0, node1, ...)")
	cmd.Flags().String(flagStartingIPAddress, "192.168.0.1", "Starting IP address (192.168.0.1 results in persistent peers list ID0@192.168.0.1:46656, ID1@192.168.0.2:46656, ...)")
	cmd.Flags().StringP(flagOutputDir, "o", "./net", "Directory to store initialization data for the network")
	cmd.Flags().StringP(flagNodeDaemonHome, "d", ".knstld", "Home directory of the node's daemon configuration")
	cmd.Flags().StringP(flagNetConfigFile, "n", "./config/net.json", "Net configuration file")
	cmd.Flags().StringP(flagKeyStorageFile, "k", "./config/keys.json", "Keys file")
	cmd.Flags().String(server.FlagMinGasPrices, fmt.Sprintf("0.000006%s", _const.StakeDenom), "Minimum gas prices to accept for transactions; All fees in a tx must meet this minimum (e.g. 0.01apple,0.001darc)")
	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|test)")
	cmd.Flags().String(flags.FlagKeyAlgorithm, string(hd.Secp256k1Type), "Key signing algorithm to generate keys for")
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func parseKeyStorage(keyStorageFile string) (*types.KeyStorage, error) {
	var keyStorage types.KeyStorage
	if err := utils.ReadJson(keyStorageFile, &keyStorage); err != nil {
		return nil, err
	}

	return &keyStorage, nil
}

func saveKey(nodeDir, keyringBackend, algoStr string, valKey *conf.ValidatorKey, key *types.Key, inBuf io.Reader) error {
	addr, secret, err := keybase.SaveCoinKey(nodeDir, keyringBackend, algoStr, key, true, true, inBuf)
	if err != nil {
		_ = os.RemoveAll(outDir)
		return err
	}

	valKey.AccAddr = addr

	//addr, secret, err := server.GenerateSaveCoinKey(kb, nodeDirName, true, algo)
	//if err != nil {
	//	_ = os.RemoveAll(outputDir)
	//	return  err
	//}

	info := map[string]string{"secret": secret}
	cliPrint, err := json.Marshal(info)
	if err != nil {
		return err
	}

	if err := utils.WriteFile(fmt.Sprintf("%v.json", "key_seed"), nodeDir, cliPrint); err != nil {
		return err
	}

	return nil
}

func parseNetConfig(netConfigFile string) (*conf.Config, error) {
	c, err := conf.ParseFile(netConfigFile)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func genAccounts(accs []conf.Account) ([]banktypes.Balance, []authtypes.GenesisAccount, error) {
	var (
		genAccounts authtypes.GenesisAccounts
		genBalances []banktypes.Balance
	)

	for _, ga := range accs {
		addr, err := sdk.AccAddressFromBech32(ga.Address)
		if err != nil {
			return nil, nil, err
		}
		var genAccount authtypes.GenesisAccount
		genAccount = authtypes.NewBaseAccount(addr, nil, 0, 0)
		if err := genAccount.Validate(); err != nil {
			return nil, nil, fmt.Errorf("failed to validate new genesis account: %w", err)
		}
		if genAccounts.Contains(addr) {
			return nil, nil, fmt.Errorf("cannot add account at existing address %s", addr)
		}
		genAccounts = append(genAccounts, genAccount)

		coins, err := sdk.ParseCoinsNormalized(strings.Join(ga.Coins, ","))
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse coins: %w", err)
		}

		balances := banktypes.Balance{Address: addr.String(), Coins: coins.Sort()}
		genBalances = append(genBalances, balances)
	}

	genAccounts = authtypes.SanitizeGenesisAccounts(genAccounts)
	genBalances = banktypes.SanitizeGenesisBalances(genBalances)

	return genBalances, genAccounts, nil
}

func clientConfig(
	clientCtx client.Context,
	nodeConfig *tmconfig.Config,
	appConfig *srvconfig.Config,
	validators []*conf.Validator) (err error) {
	var addressesIPs []string

	for _, validator := range validators {
		addressesIPs = append(addressesIPs, validator.Memo)
	}

	sort.Strings(addressesIPs)
	nodeConfig.SetRoot(configDir)
	nodeConfig.Moniker = ""
	nodeConfig.RPC.CORSAllowedOrigins = []string{"*"}
	nodeConfig.P2P.PersistentPeers = strings.Join(addressesIPs, ",")

	if err := os.MkdirAll(filepath.Join(configDir, "config"), nodeDirPerm); err != nil {
		_ = os.RemoveAll(outDir)
		return err
	}

	tmconfig.WriteConfigFile(filepath.Join(nodeConfig.RootDir, "config", "config.toml"), nodeConfig)
	srvconfig.WriteConfigFile(filepath.Join(nodeConfig.RootDir, "config", "app.toml"), appConfig)
	return nil
}

func configValidator(nodeConfig *tmconfig.Config,
	appConfig *srvconfig.Config,
	valInfo conf.ValidatorInfo,
	key *types.Key,
	genAccount *authtypes.GenesisAccount,
	keyringBackend,
	algoStr string,
	inBuf io.Reader,
) (*conf.Validator, error) {
	nodeDir := filepath.Join(outDir, valInfo.Name, app.NodeDir, nodeDaemonHomeName)
	if err := os.MkdirAll(filepath.Join(nodeDir, "config"), nodeDirPerm); err != nil {
		_ = os.RemoveAll(outDir)
		return nil, err
	}

	valNodeConfig := conf.ValNodeConfig{
		DirName:   valInfo.Name,
		DaemonDir: nodeDir,
	}

	nodeConfig.SetRoot(nodeDir)
	nodeConfig.Moniker = valInfo.Description.Moniker
	nodeConfig.RPC.ListenAddress = "tcp://0.0.0.0:26657"
	if valInfo.Config != nil {
		if err := mapstructure.Decode(valInfo.Config, nodeConfig); err != nil {
			return nil, err
		}
	}

	srvconfig.WriteConfigFile(filepath.Join(nodeDir, "config/app.toml"), appConfig)

	nodeID, valPubKey, err := genutil.InitializeNodeValidatorFiles(nodeConfig)
	if err != nil {
		_ = os.RemoveAll(outDir)
		return nil, err
	}

	memo := fmt.Sprintf("%s@%s:26656", nodeID, valInfo.IP)

	if err := saveKey(nodeDir, algoStr, keyringBackend, valInfo.Key, key, inBuf); err != nil {
		return nil, err
	}

	return &conf.Validator{
		Index:         valInfo.Index,
		Moniker:       valInfo.Description.Moniker,
		ValNodeConfig: valNodeConfig,
		GenFile:       nodeConfig.GenesisFile(),
		Memo:          memo,
		ID:            nodeID,
		Cors:          valInfo.Cors,
		ValPubKey:     valPubKey,
		IP:            valInfo.IP,
		Key:           valInfo.Key,
		Description:   valInfo.Description,
		GenAccount:    genAccount,
	}, nil
}

func configValidators(
	nodeConfig *tmconfig.Config,
	appConfig *srvconfig.Config,
	keyStorage *types.KeyStorage,
	genAccounts []authtypes.GenesisAccount,
	c *conf.Config,
	keyringBackend,
	algoStr string,
	inBuf io.Reader,
) (validators []*conf.Validator, err error) {
	if c.GlobalConfig != nil {
		if err := mapstructure.Decode(c.GlobalConfig, nodeConfig); err != nil {
			return nil, err
		}
	}

	for _, valInfo := range c.Validators {
		key, err := keyStorage.GetKey(valInfo.Key.Address)
		if err != nil {
			return nil, err
		}

		addr, err := sdk.AccAddressFromBech32(valInfo.Key.Address)
		if err != nil {
			return nil, err
		}

		var genAccount authtypes.GenesisAccount
		for _, gacc := range genAccounts {
			if gacc.GetAddress().Equals(addr) {
				genAccount = gacc
			}
		}

		node, err := configValidator(nodeConfig, appConfig, valInfo, key, &genAccount, algoStr, keyringBackend, inBuf)
		if err != nil {
			return nil, err
		}

		validators = append(validators, node)
	}

	return
}

// Override genesis app state according to chain.yml
func updateGenesisState(clientCtx client.Context, genState *map[string]json.RawMessage, appState map[string]interface{}) error {
	for mod, updates := range appState {
		var currState map[string]interface{}
		err := json.Unmarshal((*genState)[mod], &currState)
		if err != nil {
			return err
		}
		if err := mergo.Merge(&currState, updates, mergo.WithOverride); err != nil {
			return err
		}
		state, err := json.Marshal(currState)
		if err != nil {
			return err
		}
		(*genState)[mod] = state
	}

	return nil
}

func initGenFiles(
	clientCtx client.Context,
	mbm module.BasicManager,
	genesis conf.Genesis,
	validators []*conf.Validator,
	genAccounts []authtypes.GenesisAccount,
	genBalances []banktypes.Balance,
	chainID string,
) error {
	appGenState := mbm.DefaultGenesis(clientCtx.JSONMarshaler)
	// set the accounts in the genesis state
	var authGenState authtypes.GenesisState
	clientCtx.JSONMarshaler.MustUnmarshalJSON(appGenState[authtypes.ModuleName], &authGenState)

	accounts, err := authtypes.PackAccounts(genAccounts)
	if err != nil {
		return err
	}
	authGenState.Accounts = accounts
	appGenState[authtypes.ModuleName] = clientCtx.JSONMarshaler.MustMarshalJSON(&authGenState)

	// set the balances in the genesis state
	var bankGenState banktypes.GenesisState
	clientCtx.JSONMarshaler.MustUnmarshalJSON(appGenState[banktypes.ModuleName], &bankGenState)

	bankGenState.Balances = genBalances
	appGenState[banktypes.ModuleName] = clientCtx.JSONMarshaler.MustMarshalJSON(&bankGenState)

	if err := updateGenesisState(clientCtx, &appGenState, genesis.AppState); err != nil {
		return err
	}

	if err := mbm.ValidateGenesis(clientCtx.JSONMarshaler, clientCtx.TxConfig, appGenState); err != nil {
		return fmt.Errorf("error validating genesis: %s", err.Error())
	}

	appGenStateJSON, err := json.MarshalIndent(appGenState, "", "  ")
	if err != nil {
		return err
	}

	genDoc := tmtypes.GenesisDoc{
		ChainID:    chainID,
		AppState:   appGenStateJSON,
		Validators: nil,
	}

	// generate empty genesis files for each validator and save
	for _, validator := range validators {
		if err := genDoc.SaveAs(validator.GenFile); err != nil {
			return err
		}
	}

	return nil
}

func genTxs(
	clientCtx client.Context,
	validators []*conf.Validator,
	keyStorage *types.KeyStorage,
	keyringBackend string,
	chainID string,
	inBuf io.Reader,
) error {
	for _, validator := range validators {
		//genDoc, err := tmtypes.GenesisDocFromFile(validator.GenFile)
		//if err != nil {
		//	return err
		//}
		//
		//var genesisState map[string]json.RawMessage
		//if err = clientCtx.JSONMarshaler.UnmarshalJSON(genDoc.AppState, &genesisState); err != nil {
		//	return err
		//}

		//kb, err := clkeys.NewLegacyKeyBaseFromDir(validator.NodeConfig.CliDir)
		//if err != nil {
		//	return err
		//}
		//

		//c := sdk.NewCoin(types.StakeDenom, sdk.NewInt(validator.Key.CoinDelegate))
		//coins := sdk.NewCoins(c)
		//if err := genutil.ValidateAccountInGenesis(genesisState, genBalIterator, key.GetAddress(), coins, *cdc); err != nil {
		//	return err
		//}

		// TODO if staking tokens not specified in net config
		//accTokens := sdk.TokensFromConsensusPower(1000)
		//accStakingTokens := sdk.TokensFromConsensusPower(500)
		//coins := sdk.Coins{
		//	sdk.NewCoin(fmt.Sprintf("%stoken", nodeDirName), accTokens),
		//	sdk.NewCoin(sdk.DefaultBondDenom, accStakingTokens),
		//}

		createValMsg, err := stakingtypes.NewMsgCreateValidator(
			sdk.ValAddress(validator.Key.AccAddr),
			validator.ValPubKey,
			sdk.NewCoin(_const.StakeDenom, sdk.TokensFromConsensusPower(validator.Key.CoinDelegate)),
			stakingtypes.NewDescription(validator.ValNodeConfig.DaemonDir, "", "", "", ""),
			stakingtypes.NewCommissionRates(sdk.OneDec(), sdk.OneDec(), sdk.OneDec()),
			sdk.OneInt(),
		)
		if err != nil {
			return err
		}

		txBuilder := clientCtx.TxConfig.NewTxBuilder()
		txBuilder.SetMemo(validator.Memo)
		if err := txBuilder.SetMsgs(createValMsg); err != nil {
			return err
		}

		kb, err := keyring.New(sdk.KeyringServiceName(), keyringBackend, validator.ValNodeConfig.DaemonDir, inBuf)
		if err != nil {
			return err
		}

		valKey, err := keyStorage.GetKey(validator.Key.Address)
		if err != nil {
			return err
		}

		txFactory := tx.Factory{}
		txFactory = txFactory.
			WithChainID(chainID).
			WithMemo(validator.Memo).
			WithKeybase(kb).
			WithTxConfig(clientCtx.TxConfig)

		if err := tx.Sign(txFactory, valKey.GetName(), txBuilder, true); err != nil {
			return err
		}

		txBz, err := clientCtx.TxConfig.TxJSONEncoder()(txBuilder.GetTx())
		if err != nil {
			return err
		}

		if err := utils.WriteFile(fmt.Sprintf("%v.json", validator.ValNodeConfig.DirName), gentxsDir, txBz); err != nil {
			return err
		}
	}

	return nil
}

func collectGenFiles(
	clientCtx client.Context,
	config *tmconfig.Config,
	genBalIterator banktypes.GenesisBalancesIterator,
	validators []*conf.Validator,
	chainID string,
	outputDir string,
) error {
	var appState json.RawMessage
	genTime := tmtime.Now()

	for _, validator := range validators {
		//nodeDir := filepath.Join(outputDir, validator.Moniker, nodeDaemonHomeName)
		gentxsDir := filepath.Join(outputDir, "gentxs")

		config.SetRoot(validator.ValNodeConfig.DaemonDir)

		config.Moniker = validator.Moniker
		initCfg := genutiltypes.NewInitConfig(chainID, gentxsDir, validator.ID, validator.ValPubKey)

		genDoc, err := tmtypes.GenesisDocFromFile(validator.GenFile)
		if err != nil {
			return err
		}

		// TODO - nil?
		nodeAppState, err := genutil.GenAppStateFromConfig(clientCtx.JSONMarshaler, clientCtx.TxConfig, config, initCfg, *genDoc, genBalIterator)
		if err != nil {
			return err
		}

		// set the canonical application state (they should not differ)
		if appState == nil {
			appState = nodeAppState
		}

		// overwrite each validator's genesis file to have a canonical genesis time
		genFile := config.GenesisFile()
		if err := genutil.ExportGenesisFileWithTime(genFile, chainID, nil, appState, genTime); err != nil {
			return err
		}
	}

	// genesis client
	config.SetRoot(configDir)
	return genutil.ExportGenesisFileWithTime(config.GenesisFile(), chainID, nil, appState, genTime)
}

func generateNetwork(
	clientCtx client.Context,
	cmd *cobra.Command,
	nodeConfig *tmconfig.Config,
	mbm module.BasicManager,
	genBalIterator banktypes.GenesisBalancesIterator,
	outputDir,
	chainID,
	minGasPrices,
	nodeDirPrefix,
	nodeDaemonHome,
	startingIPAddress,
	keyringBackend,
	algoStr string,
	numValidators int,
) error {
	inBuf := bufio.NewReader(cmd.InOrStdin())

	chainID = viper.GetString(flags.FlagChainID)
	if chainID == "" {
		chainID = fmt.Sprintf("test-chain-%v", time.Now().Unix())
	}

	// TODO if net files empty - generate from num validators

	//depCdc := clientCtx.JSONMarshaler
	//cdc := depCdc.(codec.Marshaler)

	appConfig := srvconfig.DefaultConfig()
	appConfig.MinGasPrices = minGasPrices
	appConfig.API.Enable = true
	appConfig.Telemetry.Enabled = true
	appConfig.Telemetry.PrometheusRetentionTime = 60
	appConfig.Telemetry.EnableHostnameLabel = false
	appConfig.Telemetry.GlobalLabels = [][]string{{"chain_id", chainID}}

	netConfigFile := viper.GetString(flagNetConfigFile)
	keyStorageFile := viper.GetString(flagKeyStorageFile)

	outDir = viper.GetString(flagOutputDir)
	gentxsDir = filepath.Join(outDir, "gentxs")
	configDir = filepath.Join(outDir)

	keyStorage, err := parseKeyStorage(keyStorageFile)
	if err != nil {
		return err
	}

	c, err := parseNetConfig(netConfigFile)
	if err != nil {
		return err
	}

	genBalances, genAccounts, err := genAccounts(c.Accounts)
	if err != nil {
		return err
	}

	validators, err := configValidators(
		nodeConfig,
		appConfig,
		keyStorage,
		genAccounts,
		c,
		keyringBackend,
		algoStr,
		inBuf,
	)
	if err != nil {
		return err
	}

	if err := initGenFiles(clientCtx, mbm, c.Genesis, validators, genAccounts, genBalances, chainID); err != nil {
		return err
	}

	if err := genTxs(
		clientCtx,
		validators,
		keyStorage,
		keyringBackend,
		chainID,
		inBuf,
	); err != nil {
		return err
	}

	if err := clientConfig(clientCtx, nodeConfig, appConfig, validators); err != nil {
		return err
	}

	if err := collectGenFiles(clientCtx, nodeConfig, genBalIterator, validators, chainID, outDir); err != nil {
		return err
	}

	fmt.Printf("Successfully initialized %d node directories\n", len(validators))
	return nil
}
