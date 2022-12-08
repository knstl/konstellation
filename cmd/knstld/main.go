package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/konstellation/konstellation/app"
	"github.com/konstellation/konstellation/app/params"
	"github.com/konstellation/konstellation/cmd/knstld/cmd"
	"github.com/konstellation/konstellation/cmd/knstld/cmd/ibc"
	"github.com/konstellation/konstellation/cmd/knstld/cmd/keys"
	//"github.com/tendermint/spm/cosmoscmd"
)

func main() {
	params.SetAddressPrefixes()
	rootCmd, _ := cmd.NewRootCmd()

	rootCmd.AddCommand(keys.Commands(app.DefaultNodeHome))
	rootCmd.AddCommand(ibc.MigrateGenesisForIBC())

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
