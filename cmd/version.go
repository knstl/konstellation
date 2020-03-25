package cmd

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/konstellation/konstellation/app"
	"os"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/spf13/cobra"
)

type printInfo struct {
	Version    string          `json:"version" yaml:"version"`
	AppMessage json.RawMessage `json:"app_message" yaml:"app_message"`
}

func newPrintInfo(version string, appMessage json.RawMessage) printInfo {

	return printInfo{
		Version:    version,
		AppMessage: appMessage,
	}
}

func displayInfo(cdc *codec.Codec, info printInfo) error {
	out, err := codec.MarshalJSONIndent(cdc, info)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(os.Stderr, "%s\n", string(sdk.MustSortJSON(out)))
	return err
}

func AppStatusCmd(ctx *server.Context, cdc *codec.Codec, mbm module.BasicManager) *cobra.Command { // nolint: golint
	cmd := &cobra.Command{
		Use:   "app-status",
		Short: "Print the konstellation app version and genesis",
		Long:  "Print the konstellation app version and genesis",
		Args:  cobra.ExactArgs(0),
		RunE: func(c *cobra.Command, args []string) error {
			appState, err := codec.MarshalJSONIndent(cdc, mbm.DefaultGenesis())
			if err != nil {
				return err
			}

			return displayInfo(cdc, newPrintInfo(app.Version, appState))
		},
	}

	return cmd
}

func AppVersionCmd(ctx *server.Context, cdc *codec.Codec) *cobra.Command { // nolint: golint
	cmd := &cobra.Command{
		Use:   "app-version",
		Short: "Print the konstellation app version",
		Long:  "Print the konstellation app version",
		Args:  cobra.ExactArgs(0),
		RunE: func(c *cobra.Command, args []string) error {
			fmt.Println(app.Version)
			return nil
		},
	}

	return cmd
}
