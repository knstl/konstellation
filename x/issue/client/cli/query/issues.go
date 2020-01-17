package query

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagAddress = "address"
	flagLimit   = "limit"
	flagSymbol  = "symbol"
)

func pathQueryIssues() string {
	return fmt.Sprintf("%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryIssues)
}

func getIssuesList(cliCtx context.CLIContext, params types.IssuesParams) ([]byte, int64, error) {
	bz, err := cliCtx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, 0, err
	}
	return cliCtx.QueryWithData(pathQueryIssues(), bz)
}

// getCmdQueryIssues implements the query issue command.
func getQueryCmdIssues(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Query issue list",
		Long:  "Query all or one of the account issue list, the limit default is 30",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			address, err := sdk.AccAddressFromBech32(viper.GetString(flagAddress))
			if err != nil {
				return err
			}
			qp := types.NewIssuesParams(
				address.String(),
				viper.GetInt(flagLimit),
			)

			// Query the issues
			res, _, err := getIssuesList(cliCtx, qp)
			if err != nil {
				return err
			}

			var issues types.CoinIssues
			cdc.MustUnmarshalJSON(res, &issues)
			return cliCtx.PrintOutput(issues)
		},
	}

	cmd.Flags().String(flagAddress, "", "Token owner address")
	cmd.Flags().Int32(flagLimit, 30, "Query number of issue results per page returned")
	return cmd
}
