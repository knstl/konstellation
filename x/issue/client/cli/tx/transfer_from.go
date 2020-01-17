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

// getTxCmdTransferFrom Moves `amount` tokens from `sender` to `recipient` using the
//     * allowance mechanism. `amount` is then deducted from the caller's
//     * allowance.
func getTxCmdTransferFrom(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-from [from_address] [to_address] [amount]",
		Args:  cobra.ExactArgs(3),
		Short: "Transfer from tokens",
		Long:  "Moves `amount` tokens from `sender` to `recipient` using the allowance mechanism. `amount` is then deducted from the caller's allowance.",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			fromAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			toAddr, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoins(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgTransferFrom(cliCtx.GetFromAddress(), fromAddr, toAddr, coins)
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return validateErr
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
