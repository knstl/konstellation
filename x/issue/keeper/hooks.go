package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

type Hooks struct {
	keeper Keeper
}

// Create new box hooks
func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

func (h Hooks) CanSend(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) (bool, sdk.Error) {
	for _, v := range amt {
		i, err := h.keeper.GetIssue(ctx, v.Denom)
		if err != nil {
			return false, err
		}

		if i != nil {
			//if err := h.keeper.CheckFreeze(ctx, fromAddr, toAddr, v.Denom); err != nil {
			//	return false, err
			//}
		}
	}
	return true, nil
}

type BankHooks struct {
	issueHooks Hooks
}

func NewBankHooks(issueHooks Hooks) BankHooks {
	return BankHooks{
		issueHooks: issueHooks,
	}
}

// nolint
func (bankHooks BankHooks) CanSend(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) (bool, sdk.Error) {
	_, err := bankHooks.issueHooks.CanSend(ctx, fromAddr, toAddr, amt)
	if err != nil {
		return false, err
	}

	return true, nil
}
