package tx

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/konstellation/kn-sdk/x/issue/types"
)

// getTxCmdBurn implements burn function
func getTxCmdBurn(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn [amount]",
		Args:  cobra.ExactArgs(1),
		Short: "Burn tokens from sender",
		Long:  "Burn tokens from sender",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			coins, err := sdk.ParseCoins(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgBurn(cliCtx.GetFromAddress(), coins)
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return validateErr
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
