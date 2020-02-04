package main

import (
	"net/http"
	"os"
	"path"

	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/lcd"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	bankrest "github.com/cosmos/cosmos-sdk/x/bank/client/rest"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	distributionrest "github.com/cosmos/cosmos-sdk/x/distribution/client/rest"
	"github.com/cosmos/cosmos-sdk/x/gov"
	mintrest "github.com/cosmos/cosmos-sdk/x/mint/client/rest"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	slashingrest "github.com/cosmos/cosmos-sdk/x/slashing/client/rest"
	stakingrest "github.com/cosmos/cosmos-sdk/x/staking/client/rest"
	supplyrest "github.com/cosmos/cosmos-sdk/x/supply/client/rest"
	issuerest "github.com/konstellation/kn-sdk/x/issue/client/rest"

	"github.com/konstellation/kn-sdk/types"
	"github.com/konstellation/konstellation/app"
)

func main() {
	cobra.EnableCommandSorting = false

	cdc := app.MakeCodec()

	types.RegisterNativeCoinUnits()
	types.RegisterBech32Prefix()

	// rootCmd is the entry point for this binary
	rootCmd := &cobra.Command{
		Use:   "konstellationlcd",
		Short: "konstellation lcd server",
	}
	// Add --chain-id to persistent flags and mark it required
	rootCmd.PersistentFlags().String(client.FlagChainID, "", "Chain ID of tendermint node")
	rootCmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
		return initConfig(rootCmd)
	}
	rootCmd.AddCommand(
		lcd.ServeCommand(cdc, registerRoutes),
	)

	// prepare and add flags
	executor := cli.PrepareMainCmd(rootCmd, app.EnvPrefixLCD, app.DefaultLCDHome)
	err := executor.Execute()
	if err != nil {
		// handle with #870
		panic(err)
	}
}

func setupResponse(w *http.ResponseWriter, _ *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Headers", "header_content_type, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// registerRoutes registers the routes from the different modules for the LCD.
// NOTE: details on the routes added for each module are in the module documentation
// NOTE: If making updates here you also need to update the test helper in client/lcd/test_helper.go
func registerRoutes(rs *lcd.RestServer) {
	rs.Mux.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}
	})
	rs.Mux.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			setupResponse(&w, r)
			h.ServeHTTP(w, r)
		})
	})
	rs.Mux.Use(cors.Default().Handler)

	rpc.RegisterRPCRoutes(rs.CliCtx, rs.Mux)
	authrest.RegisterRoutes(rs.CliCtx, rs.Mux, auth.StoreKey)
	authrest.RegisterTxRoutes(rs.CliCtx, rs.Mux)
	bankrest.RegisterRoutes(rs.CliCtx, rs.Mux)
	distributionrest.RegisterRoutes(rs.CliCtx, rs.Mux, distribution.StoreKey)
	stakingrest.RegisterRoutes(rs.CliCtx, rs.Mux)
	slashingrest.RegisterRoutes(rs.CliCtx, rs.Mux)
	supplyrest.RegisterRoutes(rs.CliCtx, rs.Mux)
	mintrest.RegisterRoutes(rs.CliCtx, rs.Mux)
	govrest := gov.NewAppModuleBasic(paramsclient.ProposalHandler, distribution.ProposalHandler)
	govrest.RegisterRESTRoutes(rs.CliCtx, rs.Mux)
	issuerest.RegisterRoutes(rs.CliCtx, rs.Mux)

	registerSwaggerUI(rs)
}

func registerSwaggerUI(rs *lcd.RestServer) {
	fileSystem, err := fs.New()
	if err != nil {
		panic(err)
	}

	fileServer := http.FileServer(fileSystem)
	rs.Mux.PathPrefix("/swagger-ui/").Handler(http.StripPrefix("/swagger-ui/", fileServer))
}

func initConfig(cmd *cobra.Command) error {
	home, err := cmd.PersistentFlags().GetString(cli.HomeFlag)
	if err != nil {
		return err
	}

	cfgFile := path.Join(home, "config", "config.toml")
	if _, err := os.Stat(cfgFile); err == nil {
		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}
	if err := viper.BindPFlag(client.FlagChainID, cmd.PersistentFlags().Lookup(client.FlagChainID)); err != nil {
		return err
	}
	if err := viper.BindPFlag(cli.EncodingFlag, cmd.PersistentFlags().Lookup(cli.EncodingFlag)); err != nil {
		return err
	}
	return viper.BindPFlag(cli.OutputFlag, cmd.PersistentFlags().Lookup(cli.OutputFlag))
}
