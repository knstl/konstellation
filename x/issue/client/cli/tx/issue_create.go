package tx

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/konstellation/konstellation/x/issue/types"
)

const (
	flagDecimals           = "decimals"
	flagBurnOwnerDisabled  = "burn-owner"
	flagBurnHolderDisabled = "burn-holder"
	flagBurnFromDisabled   = "burn-from"
	flagFreezeDisabled     = "freeze"
	flagMintDisabled       = "mint"
	flagDescription        = "description"
)

// getTxCmdCreate implements issue a coin transaction command.
func getTxCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create [denom] [symbol] [owner] [issuer] [total-supply]",
		Args:    cobra.ExactArgs(3),
		Short:   "Issue a new token",
		Long:    "Issue a new token",
		Example: "$ konstellationcli issue create foocoin FOO 100000000 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			account := clientCtx.GetFromAddress()

			decimals := viper.GetUint(flagDecimals)
			totalSupply, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("Total supply %s not a valid int, please input a valid total supply", args[2])
			}
			//totalSupply = sdk.NewIntWithDecimal(totalSupply.Int64(), cast.ToInt(decimals))

			issueFeatures := types.IssueFeatures{
				BurnOwnerDisabled:  viper.GetBool(flagBurnOwnerDisabled),
				BurnHolderDisabled: viper.GetBool(flagBurnHolderDisabled),
				BurnFromDisabled:   viper.GetBool(flagBurnFromDisabled),
				MintDisabled:       viper.GetBool(flagMintDisabled),
				FreezeDisabled:     viper.GetBool(flagFreezeDisabled),
			}

			issueParams := types.IssueParams{
				Denom:         args[0],
				Symbol:        strings.ToUpper(args[1]),
				TotalSupply:   totalSupply,
				Decimals:      uint32(decimals),
				Description:   viper.GetString(flagDescription),
				IssueFeatures: &issueFeatures,
			}

			msg := types.NewMsgIssueCreate(account, account, &issueParams)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().Uint(flagDecimals, types.CoinDecimalsMaxValue, "Decimals of the token")
	cmd.Flags().Bool(flagBurnOwnerDisabled, false, "Disable token owner burn the token")
	cmd.Flags().Bool(flagBurnHolderDisabled, false, "Disable token holder burn the token")
	cmd.Flags().Bool(flagBurnFromDisabled, false, "Disable token owner burn the token from any holder")
	cmd.Flags().Bool(flagMintDisabled, false, "Token owner can not minting the token")
	cmd.Flags().Bool(flagFreezeDisabled, false, "Token holder can transfer the token in and out")

	return cmd
}
