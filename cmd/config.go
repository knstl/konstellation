package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"

	cfg "github.com/tendermint/tendermint/config"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/types/module"
)

func ConfigCmd(ctx *server.Context, _ *codec.Codec, _ module.BasicManager, _ string) *cobra.Command { // nolint: golint
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Configuration commands",
		Long:  "Configuration commands",
	}

	cmd.AddCommand(
		SetConfigCmd(ctx),
		GetConfigCmd(ctx),
	)

	return cmd
}

func SetConfigCmd(ctx *server.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "set [name] [value]",
		Short: "Change configuration file",
		Long:  "Change configuration file",
		Args:  cobra.ExactArgs(2),
		RunE: func(c *cobra.Command, args []string) error {
			config := ctx.Config
			v := make(map[string]interface{})
			_ = mapstructure.Decode(config, &v)
			v[args[0]] = args[1]
			_ = mapstructure.Decode(v, &config)

			cfg.WriteConfigFile(filepath.Join(config.RootDir, "config", "config.toml"), config)
			return nil
		},
	}
}

func GetConfigCmd(ctx *server.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "get [name]",
		Short: "Change configuration file",
		Long:  "Change configuration file",
		Args:  cobra.ExactArgs(1),
		RunE: func(c *cobra.Command, args []string) error {
			config := ctx.Config
			v := make(map[string]interface{})
			_ = mapstructure.Decode(config, &v)
			fmt.Printf("%s: %v\n", args[0], v[args[0]])
			return nil
		},
	}
}
