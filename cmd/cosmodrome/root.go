package main

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/konstellation/konstellation/app"
	"github.com/konstellation/konstellation/const"
	"github.com/spf13/cobra"
	"github.com/tendermint/spm/cosmoscmd"
	"os"
)

// NewRootCmd creates a new root command for cosmodrome. It is called once in the
// main function.
func NewRootCmd(
	accountAddressPrefix string,
	moduleBasics module.BasicManager,
) (*cobra.Command, cosmoscmd.EncodingConfig) {
	// Set config for prefixes
	cosmoscmd.SetPrefixes(accountAddressPrefix)
	_const.RegisterNativeCoinUnits()

	encodingConfig := cosmoscmd.MakeEncodingConfig(moduleBasics)

	//types.RegisterBech32Prefix()

	initClientCtx := client.Context{}.
		WithJSONMarshaler(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithBroadcastMode(flags.BroadcastBlock).
		WithHomeDir(app.DefaultNodeHome)

	rootCmd := &cobra.Command{
		Use:   version.AppName,
		Short: "cosmodrome",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if err := client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}
			return server.InterceptConfigsPreRunHandler(cmd)
		},
	}

	initRootCmd(rootCmd, encodingConfig)

	return rootCmd, encodingConfig
}

func initRootCmd(rootCmd *cobra.Command, encodingConfig cosmoscmd.EncodingConfig) {
	authclient.Codec = encodingConfig.Marshaler

	rootCmd.AddCommand(
		GenNetCmd(app.ModuleBasics, banktypes.GenesisBalancesIterator{}),
	)
}

//func addModuleInitFlags(startCmd *cobra.Command) {
//	crisis.AddModuleInitFlags(startCmd)
//	wasm.AddModuleInitFlags(startCmd)
//}
