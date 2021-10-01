package main

import (
	"github.com/spf13/cobra"
	"os"

	"github.com/cosmos/cosmos-sdk/server"

	"github.com/konstellation/konstellation/app"
)

func main() {
	cobra.EnableCommandSorting = false

	rootCmd, _ := NewRootCmd(
		app.AccountAddressPrefix,
		app.ModuleBasics,
		// this line is used by starport scaffolding # root/arguments
	)

	if err := Execute(rootCmd, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
