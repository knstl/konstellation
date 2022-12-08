package ibc

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/genutil/types"
	ibcv100 "github.com/cosmos/ibc-go/v3/modules/core/legacy/v100"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmtypes "github.com/tendermint/tendermint/types"
)

const chainUpgradeGuide = "https://github.com/cosmos/ibc-go/blob/release/v2.0.x/docs/migrations/ibc-migration-043.md"

func MigrateGenesisForIBC() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate_ibc [genesis-file] [max-expected-time-per-block]",
		Short: "Migrate genesis to support IBC version 1.",
		Long:  `Migrate genesis to support IBC version 1.`,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			var err error

			importGenesis := args[0]
			expectedTimePerBlock := args[1]
			defaultMaxExpectedTimePerBlock := 30 * time.Second

			maxExpectedTimePerBlock, err := strconv.ParseUint(expectedTimePerBlock, 10, 64)
			if err != nil {
				maxExpectedTimePerBlock = uint64(defaultMaxExpectedTimePerBlock)
			}

			genDoc, err := validateGenDoc(importGenesis)
			if err != nil {
				return err
			}

			// Since some default values are valid values, we just print to
			// make sure the user didn't forget to update these values.
			if genDoc.ConsensusParams.Evidence.MaxBytes == 0 {
				fmt.Printf("Warning: consensus_params.evidence.max_bytes is set to 0. If this is"+
					" deliberate, feel free to ignore this warning. If not, please have a look at the chain"+
					" upgrade guide at %s.\n", chainUpgradeGuide)
			}

			var initialState types.AppMap
			if err := json.Unmarshal(genDoc.AppState, &initialState); err != nil {
				return errors.Wrap(err, "failed to JSON unmarshal initial genesis state")
			}

			// add in migrate cmd function
			// expectedTimePerBlock is a new connection parameter
			// https://github.com/cosmos/ibc-go/blob/release/v1.0.x/docs/ibc/proto-docs.md#params-2
			newGenState, err := ibcv100.MigrateGenesis(initialState, clientCtx, *genDoc, maxExpectedTimePerBlock)
			if err != nil {
				return err
			}

			genDoc.AppState, err = json.Marshal(newGenState)
			if err != nil {
				return errors.Wrap(err, "failed to JSON marshal migrated genesis state")
			}

			bz, err := tmjson.Marshal(genDoc)
			if err != nil {
				return errors.Wrap(err, "failed to marshal genesis doc")
			}

			sortedBz, err := sdk.SortJSON(bz)
			if err != nil {
				return errors.Wrap(err, "failed to sort JSON genesis doc")
			}

			cmd.Println(string(sortedBz))
			return nil
		},
	}
}

// validateGenDoc reads a genesis file and validates that it is a correct
// Tendermint GenesisDoc. This function does not do any cosmos-related
// validation.
func validateGenDoc(importGenesisFile string) (*tmtypes.GenesisDoc, error) {
	genDoc, err := tmtypes.GenesisDocFromFile(importGenesisFile)
	if err != nil {
		return nil, fmt.Errorf("%s. Make sure that"+
			" you have correctly migrated all Tendermint consensus params, please see the"+
			" chain migration guide at %s for more info",
			err.Error(), chainUpgradeGuide,
		)
	}

	return genDoc, nil
}
