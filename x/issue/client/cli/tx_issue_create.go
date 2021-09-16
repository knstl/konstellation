package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
)

var _ = strconv.Itoa(0)

func CmdIssueCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "issue-create [denom] [symbol] [total-supply] --decimals [decimals] --burnOwnerDisabled [burn-owner-disabled] --burnHolderDisabled [burn-holder-disabled] --burnFromDisabled [burn-from-disabled] --mintDisabled [mint-disabled] --freezeDisabled [freeze-disabled] --description [description]",
		Short:   "Broadcast message IssueCreate",
		Long:    "Issue a new token",
		Example: "$ konstellationcli issue create foocoin FOO 100000000 --from foo",
		Args:    cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			account := clientCtx.GetFromAddress()

			totalSupply, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("Total supply %s not a valid int, please input a valid total supply", args[2])
			}
			//totalSupply = sdk.NewIntWithDecimal(totalSupply.Int64(), cast.ToInt(decimals))
			flagDecimalsStr, _ := cmd.Flags().GetString(flagDecimals)
			decimals, err := strconv.Atoi(flagDecimalsStr)
			if err != nil {
				return err
			}
			burnOwnerDisabled, _ := cmd.Flags().GetBool(flagBurnOwnerDisabled)
			burnHolderDisabled, _ := cmd.Flags().GetBool(flagBurnHolderDisabled)
			burnFromDisabled, _ := cmd.Flags().GetBool(flagBurnFromDisabled)
			mintDisabled, _ := cmd.Flags().GetBool(flagMintDisabled)
			freezeDisabled, _ := cmd.Flags().GetBool(flagFreezeDisabled)
			description, _ := cmd.Flags().GetString(flagDescription)

			issueFeatures := types.IssueFeatures{
				BurnOwnerDisabled:  burnOwnerDisabled,
				BurnHolderDisabled: burnHolderDisabled,
				BurnFromDisabled:   burnFromDisabled,
				MintDisabled:       mintDisabled,
				FreezeDisabled:     freezeDisabled,
			}

			issueParams := types.IssueParams{
				Denom:         args[0],
				Symbol:        strings.ToUpper(args[1]),
				TotalSupply:   totalSupply,
				Decimals:      uint32(decimals),
				Description:   description,
				IssueFeatures: &issueFeatures,
			}

			msg := types.NewMsgIssueCreate(account, account, &issueParams)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	decimalsStr := fmt.Sprint(types.CoinDecimalsMaxValue)
	cmd.Flags().String(flagDecimals, decimalsStr, "Decimals of the token")
	cmd.Flags().Bool(flagBurnOwnerDisabled, false, "Disable token owner burn the token")
	cmd.Flags().Bool(flagBurnHolderDisabled, false, "Disable token holder burn the token")
	cmd.Flags().Bool(flagBurnFromDisabled, false, "Disable token owner burn the token from any holder")
	cmd.Flags().Bool(flagMintDisabled, false, "Token owner can not minting the token")
	cmd.Flags().Bool(flagFreezeDisabled, false, "Token holder can transfer the token in and out")

	return cmd
}
