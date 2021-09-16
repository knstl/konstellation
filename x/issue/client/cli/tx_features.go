package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/konstellation/konstellation/x/issue/types"
)

var _ = strconv.Itoa(0)

func CmdFeatures() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "features [denom] --burnOwnerDisabled [burn-owner-disabled] --burnHolderDisabled [burn-holder-disabled] --burnFromDisabled [burn-from-disabled] --mintDisabled [mint-disabled] --freezeDisabled [freeze-disabled]",
		Short: "Broadcast message Features",
		Long:  "Enable feature for token",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			denom := args[0]

			burnOwnerDisabled, _ := cmd.Flags().GetBool(flagBurnOwnerDisabled)
			burnHolderDisabled, _ := cmd.Flags().GetBool(flagBurnHolderDisabled)
			burnFromDisabled, _ := cmd.Flags().GetBool(flagBurnFromDisabled)
			mintDisabled, _ := cmd.Flags().GetBool(flagMintDisabled)
			freezeDisabled, _ := cmd.Flags().GetBool(flagFreezeDisabled)

			issueFeatures := types.IssueFeatures{
				BurnOwnerDisabled:  burnOwnerDisabled,
				BurnHolderDisabled: burnHolderDisabled,
				BurnFromDisabled:   burnFromDisabled,
				MintDisabled:       mintDisabled,
				FreezeDisabled:     freezeDisabled,
			}

			msg := types.NewMsgFeatures(clientCtx.GetFromAddress(), denom, &issueFeatures)
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return validateErr
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
