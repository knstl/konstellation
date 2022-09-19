package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/konstellation/konstellation/app"
	"github.com/konstellation/konstellation/cmd/knstld/cmd/ibc"
	"github.com/konstellation/konstellation/cmd/knstld/cmd/keys"
	"github.com/tendermint/spm/cosmoscmd"
)

func main() {
	// todo create root.go
	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.New,
		//cosmoscmd.WithWasm(),
		// this line is used by starport scaffolding # root/arguments
	)

	rootCmd.AddCommand(keys.Commands(app.DefaultNodeHome))
	rootCmd.AddCommand(ibc.MigrateGenesisForIBC())

	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
