package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	tmcli "github.com/tendermint/tendermint/libs/cli"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"

	"github.com/konstellation/kn-sdk/types"
	"github.com/konstellation/konstellation/common/utils"
)

// InitCmd returns a command that initializes all files needed for Tendermint
// and the respective application.
func InitCmd(
	mbm module.BasicManager,
	txEncCfg client.TxEncodingConfig,
	defaultNodeHome string,
	gus types.GenesisUpdaters,
) *cobra.Command { // nolint: golint
	initCmd := genutilcli.InitCmd(mbm, defaultNodeHome)

	cmd := &cobra.Command{
		Use:   initCmd.Use,
		Short: initCmd.Short,
		Long:  initCmd.Long,
		Args:  initCmd.Args,
		RunE: func(c *cobra.Command, args []string) error {
			err := initCmd.RunE(c, args)
			if err != nil {
				return err
			}

			clientCtx := client.GetClientContextFromCmd(c)
			cdc := clientCtx.JSONMarshaler

			serverCtx := server.GetServerContextFromCmd(c)
			config := serverCtx.Config
			config.SetRoot(viper.GetString(tmcli.HomeFlag))

			nodeID, _, err := genutil.InitializeNodeValidatorFiles(config)
			if err != nil {
				return err
			}

			genFile := config.GenesisFile()
			genDoc, err := tmtypes.GenesisDocFromFile(genFile)
			if err != nil {
				return err
			}

			var appState map[string]json.RawMessage
			if err := json.Unmarshal(genDoc.AppState, &appState); err != nil {
				return err
			}

			// Update default genesis
			for _, gu := range gus {
				gu.UpdateGenesis(cdc, appState)
			}

			if err = mbm.ValidateGenesis(cdc, txEncCfg, appState); err != nil {
				return fmt.Errorf("error validating genesis: %s", err.Error())
			}

			genDoc.AppState, err = json.MarshalIndent(appState, "", " ")

			if err = genutil.ExportGenesisFile(genDoc, genFile); err != nil {
				return errors.Wrap(err, "Failed to export genesis file")
			}

			toPrint := utils.NewPrintInfo(config.Moniker, genDoc.ChainID, nodeID, "", genDoc.AppState)
			return utils.DisplayInfo(toPrint)
		},
	}

	initCmd.Flags().VisitAll(func(f *pflag.Flag) {
		cmd.Flags().String(f.Name, f.Value.String(), f.Usage)
	})

	return cmd
}
