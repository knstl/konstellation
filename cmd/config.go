package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"

	cfg "github.com/tendermint/tendermint/config"

	"github.com/cosmos/cosmos-sdk/server"
)

func ConfigCmd() *cobra.Command { // nolint: golint
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Configuration commands",
		Long:  "Configuration commands",
	}

	cmd.AddCommand(
		SetConfigCmd(),
		GetConfigCmd(),
	)

	return cmd
}

func SetConfigCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "set [name] [value]",
		Short: "Change configuration file",
		Long:  "Change configuration file",
		Args:  cobra.ExactArgs(2),
		RunE: func(c *cobra.Command, args []string) error {
			serverCtx := server.GetServerContextFromCmd(c)
			config := serverCtx.Config

			var configValues map[string]interface{}
			if err := mapstructure.Decode(config, &configValues); err != nil {
				return err
			}

			path := strings.Split(args[0], ".")
			if len(path) > 1 {
				var groupConfigs map[string]interface{}
				_ = mapstructure.Decode(configValues[path[0]], &groupConfigs)
				groupConfigs[path[1]] = args[1]
				configValues[path[0]] = groupConfigs
			} else {
				configValues[args[0]] = args[1]
			}
			if err := mapstructure.Decode(configValues, config); err != nil {
				return err
			}

			cfg.WriteConfigFile(filepath.Join(config.RootDir, "config", "config.toml"), config)
			return nil
		},
	}
}

func GetConfigCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get [name]",
		Short: "Change configuration file",
		Long:  "Change configuration file",
		Args:  cobra.ExactArgs(1),
		RunE: func(c *cobra.Command, args []string) error {
			serverCtx := server.GetServerContextFromCmd(c)
			config := serverCtx.Config

			var configValues map[string]interface{}
			_ = mapstructure.Decode(config, &configValues)

			path := strings.Split(args[0], ".")
			if len(path) > 1 {
				_ = mapstructure.Decode(configValues[path[0]], &configValues)

				fmt.Printf("%s: %v\n", args[0], configValues[path[1]])
				return nil
			}

			fmt.Printf("%s: %v\n", args[0], configValues[args[0]])
			return nil
		},
	}
}
