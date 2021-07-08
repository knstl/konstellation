package tx

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/konstellation/konstellation/x/issue/types"
)

func getTxCmdFeatures() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "features [denom]",
		Args:  cobra.ExactArgs(1),
		Short: "Enable feature",
		Long:  "Enable feature for token",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			denom := args[0]

			issueFeatures := types.IssueFeatures{
				BurnOwnerDisabled:  viper.GetBool(flagBurnOwnerDisabled),
				BurnHolderDisabled: viper.GetBool(flagBurnHolderDisabled),
				BurnFromDisabled:   viper.GetBool(flagBurnFromDisabled),
				MintDisabled:       viper.GetBool(flagMintDisabled),
				FreezeDisabled:     viper.GetBool(flagFreezeDisabled),
			}

			msg := types.NewMsgFeatures(clientCtx.GetFromAddress(), denom, &issueFeatures)
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return validateErr
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().Bool(flagBurnOwnerDisabled, false, "Disable token owner burn the token")
	cmd.Flags().Bool(flagBurnHolderDisabled, false, "Disable token holder burn the token")
	cmd.Flags().Bool(flagBurnFromDisabled, false, "Disable token owner burn the token from any holder")
	cmd.Flags().Bool(flagMintDisabled, false, "Token owner can not minting the token")
	cmd.Flags().Bool(flagFreezeDisabled, false, "Token holder can transfer the token in and out")

	return cmd
}
