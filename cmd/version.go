package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/konstellation/konstellation/app"
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

//
//func displayInfo(cdc codec.JSONMarshaler, info printInfo) error {
//	out, err := cdc.MarshalInterfaceJSON(info)
//	if err != nil {
//		return err
//	}
//
//	_, err = fmt.Fprintf(os.Stderr, "%s\n", string(sdk.MustSortJSON(out)))
//	return err
//}

func AppStatusCmd(ctx *server.Context, cdc codec.JSONMarshaler, mbm module.BasicManager) *cobra.Command { // nolint: golint
	cmd := &cobra.Command{
		Use:   "app-status",
		Short: "Print the konstellation app version and genesis",
		Long:  "Print the konstellation app version and genesis",
		Args:  cobra.ExactArgs(0),
		RunE: func(c *cobra.Command, args []string) error {
			//appState, err := codec.MarshalJSONIndent(mbm.DefaultGenesis(cdc))
			//if err != nil {
			//	return err
			//}
			//
			//return displayInfo(cdc, newPrintInfo(app.Version, appState))

			return nil
		},
	}

	return cmd
}

func AppVersionCmd(ctx *server.Context, cdc codec.JSONMarshaler) *cobra.Command { // nolint: golint
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
