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

// getTxCmdBurnFrom implements burn  function
func getTxCmdBurnFrom(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-from [from_address] [amount]",
		Args:  cobra.ExactArgs(2),
		Short: "Burn tokens from recipient",
		Long:  "Burn tokens from recipient",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			from, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgBurnFrom(cliCtx.GetFromAddress(), from, coins)
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return validateErr
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
