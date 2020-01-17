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

// getTxCmdMintTo implements mint function to recipient
func getTxCmdMintTo(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-to [to_address] [amount]",
		Args:  cobra.ExactArgs(2),
		Short: "Mint tokens to address",
		Long:  "Mint tokens to address",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			to, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgMintTo(cliCtx.GetFromAddress(), to, coins)
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return validateErr
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
