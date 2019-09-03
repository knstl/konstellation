package main

import (
	"encoding/json"
	"io"

	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
	genaccountscli "github.com/cosmos/cosmos-sdk/x/genaccounts/client/cli"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/konstellation/konstellation/app"
)

func main() {
	cobra.EnableCommandSorting = false

	cdc := app.MakeCodec()

	// kinit.Init()

	ctx := server.NewDefaultContext()

	rootCmd := &cobra.Command{
		Use:               "konstellation",
		Short:             "Konstellation App Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}
	// CLI commands to initialize the chain
	rootCmd.AddCommand(
		genutilcli.InitCmd(ctx, cdc, app.ModuleBasics, app.DefaultNodeHome),
		genutilcli.GenTxCmd(ctx, cdc, app.ModuleBasics, staking.AppModuleBasic{}, genaccounts.AppModuleBasic{}, app.DefaultNodeHome, app.DefaultCLIHome),
		genutilcli.CollectGenTxsCmd(ctx, cdc, genaccounts.AppModuleBasic{}, app.DefaultNodeHome),
		genaccountscli.AddGenesisAccountCmd(ctx, cdc, app.DefaultNodeHome, app.DefaultCLIHome),
		genutilcli.ValidateGenesisCmd(ctx, cdc, app.ModuleBasics),
	)

	server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, "KONSTELLATION", app.DefaultNodeHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, _ io.Writer) abci.Application {
	return app.NewKonstellationApp(logger, db)
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, _ io.Writer, height int64, forZeroHeight bool,
	jailWhiteList []string) (json.RawMessage, []tmtypes.GenesisValidator, error) {
	if height != -1 {
		kApp := app.NewKonstellationApp(logger, db)
		err := kApp.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}

		return kApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}

	kApp := app.NewKonstellationApp(logger, db)

	return kApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
