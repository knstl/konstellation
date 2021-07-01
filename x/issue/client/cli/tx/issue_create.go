package tx

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/konstellation/kn-sdk/x/issue/types"
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
func getTxCmdCreate(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create [denom] [symbol] [owner] [issuer] [total-supply]",
		Args:    cobra.ExactArgs(3),
		Short:   "Issue a new token",
		Long:    "Issue a new token",
		Example: "$ konstellationcli issue create foocoin FOO 100000000 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			account := cliCtx.GetFromAddress()

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
				Decimals:      decimals,
				Description:   viper.GetString(flagDescription),
				IssueFeatures: issueFeatures,
			}

			msg := types.NewMsgIssueCreate(account, account, &issueParams)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
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
