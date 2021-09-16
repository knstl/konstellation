package cli

import (
	"github.com/spf13/cobra"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/oracle/types"
)

var _ = strconv.Itoa(0)

func CmdSetExchangeRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-exchange-rate [pair] [rate] [denoms]",
		Short: "Set exchange rate",
		Long:  `Set exchange rate for pair: kbtckusd 5000000000000 kbtc,kusd`,
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			pair, rateStr, denomsStr := args[0], args[1], args[2]
			if pair == "" {
				return types.ErrInvalidPair
			}

			denoms := strings.Split(denomsStr, ",")
			if len(denoms) != 2 {
				return types.ErrInvalidDenoms
			}

			rate, err := sdk.NewDecFromStr(rateStr)
			if err != nil {
				return types.ErrInvalidRate
			}

			exchangeRate := types.ExchangeRate{
				Pair:   pair,
				Rate:   rate,
				Denoms: denoms,
			}

			msg := types.NewMsgSetExchangeRate(clientCtx.GetFromAddress().String(), &exchangeRate)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
