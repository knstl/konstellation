package tx

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/spf13/cobra"
)

// getTxCmdMint implements transfer function
func getTxCmdMint(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [amount]",
		Args:  cobra.ExactArgs(1),
		Short: "Mint tokens to sender",
		Long:  "Mint tokens to sender",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			coins, err := sdk.ParseCoins(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgMint(cliCtx.GetFromAddress(), coins)
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return validateErr
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
