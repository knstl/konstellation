package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	tmcli "github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

type printInfo struct {
	AppMessage json.RawMessage `json:"app_message" yaml:"app_message"`
}

func newPrintInfo(appMessage json.RawMessage) printInfo {

	return printInfo{
		AppMessage: appMessage,
	}
}

func displayInfo(cdc *codec.Codec, info printInfo) error {
	fmt.Println()
	out, err := codec.MarshalJSONIndent(cdc, info)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(os.Stderr, "%s\n", string(sdk.MustSortJSON(out)))
	return err
}

// InitCmd returns a command that initializes all files needed for Tendermint
// and the respective application.
func InitCmd(ctx *server.Context, cdc *codec.Codec, mbm module.BasicManager,
	defaultNodeHome string) *cobra.Command { // nolint: golint
	initCmd := genutilcli.InitCmd(ctx, cdc, mbm, defaultNodeHome)

	cmd := &cobra.Command{
		Use:   initCmd.Use,
		Short: initCmd.Short,
		Long:  initCmd.Long,
		Args:  initCmd.Args,
		RunE: func(c *cobra.Command, args []string) error {
			err := initCmd.RunE(c, args)
			if err != nil {
				panic(err)
			}

			config := ctx.Config
			config.SetRoot(viper.GetString(tmcli.HomeFlag))

			genFile := config.GenesisFile()
			genDoc, err := types.GenesisDocFromFile(genFile)
			if err != nil {
				panic(err)
			}

			var appState map[string]json.RawMessage
			err = cdc.UnmarshalJSON(genDoc.AppState, &appState)
			if err != nil {
				panic(err)
			}

			// TODO move init genesis to konstellation staking module
			// -------------

			var genesisState staking.GenesisState
			err = cdc.UnmarshalJSON(appState[staking.ModuleName], &genesisState)
			if err != nil {
				panic(err)
			}

			genesisState.Params.BondDenom = "darc"

			appState[staking.ModuleName] = cdc.MustMarshalJSON(genesisState)

			// --------------

			genDoc.AppState = cdc.MustMarshalJSON(appState)
			if err = genutil.ExportGenesisFile(genDoc, genFile); err != nil {
				return errors.Wrap(err, "Failed to export genesis file")
			}

			toPrint := newPrintInfo(genDoc.AppState)
			return displayInfo(cdc, toPrint)
		},
	}

	initCmd.Flags().VisitAll(func(f *pflag.Flag) {
		cmd.Flags().String(f.Name, f.Value.String(), f.Usage)
	})

	return cmd
}
