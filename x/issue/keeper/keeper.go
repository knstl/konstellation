package keeper

import (
	"fmt"
	"strings"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/konstellation/konstellation/x/issue/types"
	// this line is used by starport scaffolding # ibc/keeper/import
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc      codec.Marshaler
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
		// this line is used by starport scaffolding # ibc/keeper/attribute
		// The reference to the Paramstore to get and set issue specific params
		paramSubspace paramstypes.Subspace

		// The reference to the Param Keeper to get and set Global Params
		paramsKeeper paramskeeper.Keeper

		ak types.AccountKeeper
		// The reference to the CoinKeeper to modify balances
		//sk types.CoinKeeper
		//sk types.SupplyKeeper
		csk types.CoinSupplyKeeper
		// The reference to the FeeCollectionKeeper to add fee
		feeCollectorName string
	}
)

func NewKeeper(
	cdc codec.Marshaler,
	storeKey,
	memKey sdk.StoreKey,
	// this line is used by starport scaffolding # ibc/keeper/parameter
	ak types.AccountKeeper,
	//ck types.CoinKeeper,
	//sk types.SupplyKeeper,
	csk types.CoinSupplyKeeper,
	feeCollectorName string,
	paramsKeeper paramskeeper.Keeper,
	paramSpace paramstypes.Subspace,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		// this line is used by starport scaffolding # ibc/keeper/return
		paramSubspace: paramSpace.WithKeyTable(types.ParamKeyTable()),
		paramsKeeper:  paramsKeeper,
		ak:            ak,
		//ck:               ck,
		//sk:               sk,
		csk:              csk,
		feeCollectorName: feeCollectorName,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k *Keeper) GetCodec() codec.BinaryMarshaler {
	return k.cdc
}

// ----------------------- last id ----------------

func (k *Keeper) getLastId(ctx sdk.Context) sdk.Int {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(KeyLastIssueId)
	if bz == nil {
		return sdk.NewIntFromUint64(types.InitLastId)
	}

	id := sdk.IntProto{}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &id)
	return id.Int
}

func (k *Keeper) setLastId(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	idSdkInt := sdk.NewIntFromUint64(id)
	idIntProto := sdk.IntProto{Int: idSdkInt}
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(&idIntProto)
	store.Set(KeyLastIssueId, bz)
}

func (k *Keeper) incLastId(ctx sdk.Context) {
	k.setLastId(ctx, k.getLastId(ctx).Uint64()+1)
}

func (k *Keeper) SetLastId(ctx sdk.Context, id uint64) {
	k.setLastId(ctx, id)
}

func (k *Keeper) GetLastId(ctx sdk.Context) uint64 {
	return k.getLastId(ctx).Uint64()
}

// ----------------------- boundary denoms ----------------

func (k *Keeper) updateLeftBoundaryDenom(ctx sdk.Context, issue *types.CoinIssue) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(KeyFirstIssueDenom)
	if bz == nil {
		denom := types.CoinIssueDenom{Denom: issue.Denom}
		bz = k.cdc.MustMarshalBinaryLengthPrefixed(&denom)
		store.Set(KeyFirstIssueDenom, bz)
	}
}

func (k *Keeper) getLeftBoundaryDenom(ctx sdk.Context) (denom string) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(KeyFirstIssueDenom)
	if bz == nil {
		return
	}

	issueDenom := types.CoinIssueDenom{}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &issueDenom)
	denom = issueDenom.Denom
	return
}

func (k *Keeper) updateRightBoundaryDenom(ctx sdk.Context, issue *types.CoinIssue) {
	store := ctx.KVStore(k.storeKey)
	denom := types.CoinIssueDenom{Denom: issue.Denom}
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(&denom)
	store.Set(KeyLastIssueDenom, bz)
}

func (k *Keeper) getRightBoundaryDenom(ctx sdk.Context) (denom string) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(KeyLastIssueDenom)
	if bz == nil {
		return
	}

	issueDenom := types.CoinIssueDenom{}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &issueDenom)
	denom = issueDenom.Denom
	return
}

func (k *Keeper) updateBoundaryDenoms(ctx sdk.Context, issue *types.CoinIssue) {
	k.updateLeftBoundaryDenom(ctx, issue)
	k.updateRightBoundaryDenom(ctx, issue)
}

func (k *Keeper) getBoundaryDenoms(ctx sdk.Context) (string, string) {
	left := k.getLeftBoundaryDenom(ctx)
	right := k.getRightBoundaryDenom(ctx)

	return left, right
}

// ----------------------- address:denom pair ----------------

func (k *Keeper) getAddressDenoms(ctx sdk.Context, accAddress string) (denoms []string) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(KeyAddressDenoms(accAddress))
	if bz == nil {
		return []string{}
	}

	issueDenoms := types.CoinIssueDenoms{}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &issueDenoms)
	denoms = []string{}
	for _, issueDenom := range issueDenoms.Denoms {
		denoms = append(denoms, issueDenom.Denom)
	}
	return
}

func (k *Keeper) setAddressDenoms(ctx sdk.Context, accAddress string, denoms []string) {
	store := ctx.KVStore(k.storeKey)
	issueDenoms := types.CoinIssueDenoms{Denoms: []*types.CoinIssueDenom{}}
	for _, denom := range denoms {
		issueDenoms.Denoms = append(issueDenoms.Denoms, &types.CoinIssueDenom{Denom: denom})
	}

	bz := k.cdc.MustMarshalBinaryLengthPrefixed(&issueDenoms)
	store.Set(KeyAddressDenoms(accAddress), bz)
}

func (k *Keeper) addAddressDenom(ctx sdk.Context, issue *types.CoinIssue) {
	denoms := k.getAddressDenoms(ctx, issue.GetOwner())
	denoms = append(denoms, issue.GetDenom())
	k.setAddressDenoms(ctx, issue.GetOwner(), denoms)
}

func (k *Keeper) deleteAllAddressDenoms(ctx sdk.Context, accAddress string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(KeyAddressDenoms(accAddress))
}

func (k *Keeper) deleteAddressDenom(ctx sdk.Context, accAddress string, denom string) {
	denoms := k.getAddressDenoms(ctx, accAddress)
	for i, d := range denoms {
		if d == denom {
			denoms = append(denoms[:i], denoms[i+1:]...)
		}
	}

	k.setAddressDenoms(ctx, accAddress, denoms)
}

// ----------------------- symbol:denom pair ----------------

func (k *Keeper) getSymbolDenom(ctx sdk.Context, symbol string) (denom string) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(KeySymbolDenom(symbol))
	if bz == nil {
		return
	}

	issueDenom := types.CoinIssueDenom{}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &issueDenom)
	denom = issueDenom.Denom
	return
}

func (k *Keeper) setSymbolDenom(ctx sdk.Context, symbol, denom string) {
	store := ctx.KVStore(k.storeKey)
	issueDenom := types.CoinIssueDenom{Denom: denom}
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(&issueDenom)
	store.Set(KeyLastIssueDenom, bz)
}

func (k *Keeper) addSymbolDenom(ctx sdk.Context, issue *types.CoinIssue) {
	k.setSymbolDenom(ctx, issue.GetSymbol(), issue.GetDenom())
}

// ----------------------- id:denom pair ----------------

func (k *Keeper) getIdDenom(ctx sdk.Context, id uint64) (denom string) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(KeyIdDenom(id))
	if bz == nil {
		return
	}

	issueDenom := types.CoinIssueDenom{}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &issueDenom)
	denom = issueDenom.Denom
	return
}

func (k *Keeper) setIdDenom(ctx sdk.Context, id uint64, denom string) {
	store := ctx.KVStore(k.storeKey)
	issueDenom := types.CoinIssueDenom{Denom: denom}
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(&issueDenom)
	store.Set(KeyLastIssueDenom, bz)
}

func (k *Keeper) addIdDenom(ctx sdk.Context, issue *types.CoinIssue) {
	k.setIdDenom(ctx, issue.GetId(), issue.GetDenom())
}

// ----------------------- issue -----------------------

func (k *Keeper) getIssue(ctx sdk.Context, denom string) *types.CoinIssue {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(KeyIssuer(denom))
	if len(bz) == 0 {
		return nil
	}

	var coinIssue types.CoinIssue
	k.GetCodec().MustUnmarshalBinaryLengthPrefixed(bz, &coinIssue)
	return &coinIssue
}

func (k *Keeper) GetIssue(ctx sdk.Context, denom string) (*types.CoinIssue, *sdkerrors.Error) {
	issue := k.getIssue(ctx, denom)
	if issue == nil {
		return nil, types.ErrUnknownIssue(denom)
	}

	return issue, nil
}

func (k *Keeper) checkOwner(_ sdk.Context, issue *types.CoinIssue, owner sdk.AccAddress) (*types.CoinIssue, *sdkerrors.Error) {
	if !sdk.AccAddress(issue.Owner).Equals(owner) {
		return nil, types.ErrOwnerMismatch(issue.Denom)
	}

	return issue, nil
}

func (k *Keeper) getIssueIfOwner(ctx sdk.Context, denom string, owner sdk.AccAddress) (*types.CoinIssue, *sdkerrors.Error) {
	issue := k.getIssue(ctx, denom)
	if issue == nil {
		return nil, types.ErrUnknownIssue(denom)
	}

	return k.checkOwner(ctx, issue, owner)
}

func (k *Keeper) addIssue(ctx sdk.Context, issue *types.CoinIssue) {
	k.addAddressDenom(ctx, issue)
	k.addSymbolDenom(ctx, issue)
	k.addIdDenom(ctx, issue)
	k.setIssue(ctx, issue)
	k.updateBoundaryDenoms(ctx, issue)
	k.incLastId(ctx)
}

func (k *Keeper) setIssue(ctx sdk.Context, issue *types.CoinIssue) {
	store := ctx.KVStore(k.storeKey)
	store.Set(KeyIssuer(issue.GetDenom()), k.GetCodec().MustMarshalBinaryLengthPrefixed(issue))
}

func (k *Keeper) AddIssue(ctx sdk.Context, issue *types.CoinIssue) {
	k.addIssue(ctx, issue)
}

func (k *Keeper) CreateIssue(ctx sdk.Context, owner, issuer sdk.AccAddress, params *types.IssueParams) *types.CoinIssue {
	issue := types.NewCoinIssue(owner, issuer, params)
	issue.SetId(k.getLastId(ctx).Uint64())
	issue.SetIssueTime(ctx.BlockHeader().Time.Unix())

	return issue
}

func (k *Keeper) ChangeFeatures(ctx sdk.Context, owner sdk.AccAddress, denom string, features *types.IssueFeatures) *sdkerrors.Error {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeChangeFeatures,
			sdk.NewAttribute(types.AttributeKeyOwner, owner.String()),
			sdk.NewAttribute(types.AttributeKeyDenom, denom),
			sdk.NewAttribute(types.AttributeKeyFeatures, features.String()),
		),
	)

	i, err := k.getIssueIfOwner(ctx, denom, owner)
	if err != nil {
		return err
	}

	i.SetFeatures(features)
	k.setIssue(ctx, i)

	return nil
}

func (k *Keeper) ChangeDescription(ctx sdk.Context, owner sdk.AccAddress, denom string, description string) *sdkerrors.Error {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeChangeDescription,
			sdk.NewAttribute(types.AttributeKeyOwner, owner.String()),
			sdk.NewAttribute(types.AttributeKeyDenom, denom),
			sdk.NewAttribute(types.AttributeKeyDescription, description),
		),
	)

	i, err := k.getIssueIfOwner(ctx, denom, owner)
	if err != nil {
		return err
	}

	i.Description = description
	k.setIssue(ctx, i)

	return nil
}

// ----------------------- issues -----------------------

func (k *Keeper) getIssues(ctx sdk.Context, denoms []string) types.CoinIssues {
	length := len(denoms)
	issues := make(types.CoinIssues, 0, length)

	for _, v := range denoms {
		issues = append(issues, k.getIssue(ctx, v))
	}

	return issues
}

func (k *Keeper) getIssuesByAddress(ctx sdk.Context, accAddress string) types.CoinIssues {
	return k.getIssues(ctx, k.getAddressDenoms(ctx, accAddress))
}

func (k *Keeper) iterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	//first, last := k.getBoundaryDenoms(ctx)
	//if first == last {
	//	return store.ReverseIterator(KeyIdIssuer(first), nil)
	//}
	//
	//return store.ReverseIterator(KeyIssuer(last), KeyIssuer(first))

	lastId := k.getLastId(ctx).Uint64()
	if lastId == types.InitLastId {
		lastId++
	}

	return store.ReverseIterator(KeyIdDenom(types.InitLastId), KeyIdDenom(lastId))
}

func (k *Keeper) ListAll(ctx sdk.Context) types.CoinIssues {
	iterator := k.iterator(ctx)
	defer iterator.Close()

	denoms := make([]string, 0)
	for ; iterator.Valid(); iterator.Next() {
		bz := iterator.Value()
		if len(bz) == 0 {
			continue
		}

		var coinIssue types.CoinIssue
		k.GetCodec().MustUnmarshalBinaryLengthPrefixed(bz, &coinIssue)
		denoms = append(denoms, coinIssue.Denom)
	}

	return k.getIssues(ctx, denoms)
}

func (k *Keeper) List(ctx sdk.Context, params types.IssuesParams) types.CoinIssues {
	if params.Owner != "" {
		return k.getIssuesByAddress(ctx, params.Owner)
	}

	iterator := k.iterator(ctx)
	defer iterator.Close()

	denoms := make([]string, 0, params.Limit)
	for ; iterator.Valid(); iterator.Next() {
		bz := iterator.Value()
		if len(bz) == 0 {
			continue
		}

		var coinIssue types.CoinIssue
		k.GetCodec().MustUnmarshalBinaryLengthPrefixed(bz, &coinIssue)
		denoms = append(denoms, coinIssue.Denom)
		if len(denoms) >= int(params.Limit) {
			break
		}
	}

	return k.getIssues(ctx, denoms)
}

// ----------------------- allowance -----------------------

func (k *Keeper) setAllowance(ctx sdk.Context, owner, spender sdk.AccAddress, amount sdk.Coin) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(&amount)
	store.Set(KeyAllowance(amount.Denom, owner, spender), bz)
}

func (k *Keeper) setAllowances(ctx sdk.Context, owner, spender sdk.AccAddress, amount sdk.Coin) {
	store := ctx.KVStore(k.storeKey)
	allowances := make(types.Allowances, 0)
	allowanceList := types.AllowanceList{Allowances: []*types.Allowance{}}
	allowance := types.NewAllowance(amount, spender)
	bz := store.Get(KeyAllowances(amount.Denom, owner))
	if bz != nil {
		k.GetCodec().MustUnmarshalBinaryLengthPrefixed(bz, &allowanceList)
	}
	allowances = types.Allowances(allowanceList.Allowances)
	if i := allowances.ContainsI(allowance); i > -1 {
		allowances[i] = allowance
	} else {
		allowances = append(allowances, allowance)
	}
	allowanceSlice := []*types.Allowance{}
	for _, allowance := range allowances {
		allowanceSlice = append(allowanceSlice, allowance)
	}
	allowanceList = types.AllowanceList{Allowances: allowanceSlice}
	bz = k.cdc.MustMarshalBinaryLengthPrefixed(&allowanceList)
	store.Set(KeyAllowances(amount.Denom, owner), bz)
}

func (k *Keeper) allowance(ctx sdk.Context, owner, spender sdk.AccAddress, denom string) sdk.Coin {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(KeyAllowance(denom, owner, spender))
	if bz == nil {
		return sdk.NewCoin(denom, sdk.ZeroInt())
	}

	var amount sdk.Coin
	k.GetCodec().MustUnmarshalBinaryLengthPrefixed(bz, &amount)
	amount.Denom = denom
	return amount
}

func (k *Keeper) allowances(ctx sdk.Context, owner sdk.AccAddress, denom string) types.Allowances {
	store := ctx.KVStore(k.storeKey)
	allowanceList := types.AllowanceList{Allowances: []*types.Allowance{}}
	bz := store.Get(KeyAllowances(denom, owner))
	if bz != nil {
		k.GetCodec().MustUnmarshalBinaryLengthPrefixed(bz, &allowanceList)
	}
	allowances := types.Allowances(allowanceList.Allowances)

	return allowances
}

func (k *Keeper) approve(ctx sdk.Context, owner, spender sdk.AccAddress, amount sdk.Coin) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeApprove,
			sdk.NewAttribute(types.AttributeKeyDenom, amount.Denom),
			sdk.NewAttribute(types.AttributeKeyOwner, owner.String()),
			sdk.NewAttribute(types.AttributeKeySpender, spender.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, amount.Amount.String()),
		),
	)

	k.setAllowance(ctx, owner, spender, amount)
	k.setAllowances(ctx, owner, spender, amount)
}

func (k *Keeper) increaseAllowance(ctx sdk.Context, owner, spender sdk.AccAddress, amount sdk.Coin) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeIncreaseAllowance,
			sdk.NewAttribute(types.AttributeKeyDenom, amount.Denom),
			sdk.NewAttribute(types.AttributeKeyOwner, owner.String()),
			sdk.NewAttribute(types.AttributeKeySpender, spender.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, amount.Amount.String()),
		),
	)

	allowance := k.allowance(ctx, owner, spender, amount.Denom)
	k.approve(ctx, owner, spender, allowance.Add(amount))
}

func (k *Keeper) decreaseAllowance(ctx sdk.Context, owner, spender sdk.AccAddress, amount sdk.Coin) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeDecreaseAllowance,
			sdk.NewAttribute(types.AttributeKeyDenom, amount.Denom),
			sdk.NewAttribute(types.AttributeKeyOwner, owner.String()),
			sdk.NewAttribute(types.AttributeKeySpender, spender.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, amount.Amount.String()),
		),
	)

	allowance := k.allowance(ctx, owner, spender, amount.Denom)
	if allowance.IsGTE(amount) {
		k.approve(ctx, owner, spender, allowance.Sub(amount))
	} else {
		k.approve(ctx, owner, spender, sdk.NewCoin(amount.Denom, sdk.ZeroInt()))
	}
}

// ----------------------- freeze -----------------------

func (k *Keeper) setFreeze(ctx sdk.Context, denom string, holder sdk.AccAddress, freeze *types.Freeze) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(freeze)
	store.Set(KeyFreeze(denom, holder), bz)
}

func (k *Keeper) getFreeze(ctx sdk.Context, denom string, holder sdk.AccAddress) *types.Freeze {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(KeyFreeze(denom, holder))
	if len(bz) == 0 {
		return types.NewFreeze(false, false)
	}
	var freeze types.Freeze
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &freeze)
	return &freeze
}

func (k *Keeper) GetFreezes(ctx sdk.Context, denom string) []*types.AddressFreeze {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, PrefixFreeze(denom))
	defer iterator.Close()
	list := make(types.AddressFreezes, 0)
	for ; iterator.Valid(); iterator.Next() {
		var freeze types.Freeze
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &freeze)
		keys := strings.Split(string(iterator.Key()), KeyDelimiter)
		address := keys[len(keys)-1]
		list = append(list, types.NewAddressFreeze(address, freeze.In, freeze.Out))
	}
	return list
}

func (k *Keeper) CheckFreeze(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, denom string) *sdkerrors.Error {
	freeze := k.GetFreeze(ctx, denom, from)
	if freeze.Out {
		return types.ErrCanNotTransferOut(denom, from.String())
	}

	freeze = k.GetFreeze(ctx, denom, to)
	if freeze.In {
		return types.ErrCanNotTransferIn(denom, to.String())
	}

	return nil
}

// ----------------------- transfers -----------------------

func (k *Keeper) transfer(ctx sdk.Context, from, to sdk.AccAddress, coins sdk.Coins) *sdkerrors.Error {
	for _, coin := range coins {
		if !k.csk.SendEnabledCoin(ctx, coin) {
			return banktypes.ErrSendDisabled
		}
	}

	if k.csk.BlockedAddr(to) {
		return sdkerrors.ErrUnauthorized
	}

	for _, coin := range coins {
		if err := k.CheckFreeze(ctx, from, to, coin.Denom); err != nil {
			return err
		}
	}

	if err := k.csk.SendCoins(ctx, from, to, coins); err != nil {
		return types.ErrCanNotTransferIn(coins[0].Denom, to.String())
	}

	return nil
}

func (k *Keeper) mint(ctx sdk.Context, minter, to sdk.AccAddress, coins sdk.Coins) *sdkerrors.Error {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(sdk.AttributeKeyAmount, coins.String()),
			sdk.NewAttribute(types.AttributeKeyMinter, minter.String()),
			sdk.NewAttribute(types.AttributeKeyRecipient, to.String()),
		),
	)

	for i, coin := range coins {
		issue, err := k.getIssueIfOwner(ctx, coin.Denom, minter)
		if err != nil {
			coins = append(coins[:i], coins[i+1:]...)
		} else {
			issue.AddTotalSupply(coin.Amount)
			if issue.QuoDecimals(issue.TotalSupply).LTE(types.CoinMaxTotalSupply) {
				k.setIssue(ctx, issue)
			}
		}

		if issue != nil {
			if issue.MintDisabled {
				return types.ErrCanNotMint(issue.Denom)
			}
		}
	}

	if err := k.csk.MintCoins(ctx, types.ModuleName, coins); err != nil {
		return types.ErrCanNotMint(coins[0].Denom)
	}

	if err := k.csk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, coins); err != nil {
		return types.ErrCanNotTransferIn(coins[0].Denom, to.String())
	}
	return nil

}

func (k *Keeper) burn(ctx sdk.Context, burner, from sdk.AccAddress, coins sdk.Coins) *sdkerrors.Error {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeBurn,
			sdk.NewAttribute(sdk.AttributeKeyAmount, coins.String()),
			sdk.NewAttribute(types.AttributeKeyBurner, burner.String()),
			sdk.NewAttribute(types.AttributeKeyFrom, from.String()),
		),
	)

	acc := k.ak.GetAccount(ctx, from)

	for i, coin := range coins {
		issue := k.getIssue(ctx, coin.Denom)

		if issue == nil {
			return types.ErrUnknownIssue(coin.Denom)
		}

		if sdk.AccAddress(issue.Owner).Equals(from) && issue.BurnOwnerDisabled {
			return types.ErrCanNotBurnOwner(issue.Denom)
		}

		if burner.Equals(from) {
			if issue.BurnHolderDisabled {
				return types.ErrCanNotBurnHolder(issue.Denom)
			}
		} else {
			if issue.BurnFromDisabled {
				return types.ErrCanNotBurnFrom(issue.Denom)
			}
		}

		issue.SubTotalSupply(coin.Amount)
		k.setIssue(ctx, issue)

		currAmt := sdk.NewInt(int64(acc.GetAccountNumber()))
		if coin.Amount.GT(currAmt) {
			coin.Amount = currAmt
			coins = append(coins[:i], coin)
			coins = append(coins[:i+1], coins[i+1:]...)
		}
	}

	if err := k.csk.SendCoinsFromAccountToModule(ctx, from, types.ModuleName, coins); err != nil {
		return types.ErrCanNotTransferOut(coins[0].Denom, from.String())
	}

	if err := k.csk.BurnCoins(ctx, types.ModuleName, coins); err != nil {
		return types.ErrCanNotBurnFrom(coins[0].Denom)
	}

	return nil
}

func (k *Keeper) freeze(ctx sdk.Context, holder sdk.AccAddress, denom, op string, freeze bool) *sdkerrors.Error {
	f := k.getFreeze(ctx, denom, holder)
	switch op {
	case types.FreezeIn:
		f.In = freeze
	case types.FreezeOut:
		f.Out = freeze
	case types.FreezeInOut:
		f.In = freeze
		f.Out = freeze
	default:
		return types.ErrInvalidFreezeOp(op)
	}

	k.setFreeze(ctx, denom, holder, f)

	return nil
}

// ----------------------- ERC20 -----------------------

func (k *Keeper) Issue(ctx sdk.Context, issue *types.CoinIssue) *sdkerrors.Error {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeIssue,
			sdk.NewAttribute(sdk.AttributeKeyAmount, issue.ToCoin().String()),
			sdk.NewAttribute(types.AttributeKeyIssuer, issue.GetIssuer()),
		),
	)

	i := k.getIssue(ctx, issue.Denom)
	if i != nil {
		return types.ErrIssueAlreadyExists
	}

	k.addIssue(ctx, issue)

	if err := k.csk.MintCoins(ctx, types.ModuleName, issue.ToCoins()); err != nil {
		return types.ErrCanNotMint(issue.Denom)
	}
	owner, _ := sdk.AccAddressFromBech32(issue.GetOwner())
	if err := k.csk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, issue.ToCoins()); err != nil {
		return types.ErrCanNotTransferIn(issue.Denom, owner.String())
	}

	return nil
}

func (k *Keeper) Transfer(ctx sdk.Context, from, to sdk.AccAddress, coins sdk.Coins) *sdkerrors.Error {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeTransfer,
			sdk.NewAttribute(sdk.AttributeKeySender, from.String()),
		),
	)

	return k.transfer(ctx, from, to, coins)
}

func (k *Keeper) TransferFrom(ctx sdk.Context, sender, from, to sdk.AccAddress, coins sdk.Coins) *sdkerrors.Error {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeTransferFrom,
			sdk.NewAttribute(types.AttributeKeySpender, sender.String()),
			sdk.NewAttribute(sdk.AttributeKeySender, from.String()),
			sdk.NewAttribute(types.AttributeKeyRecipient, to.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, coins.String()),
		),
	)

	for _, coin := range coins {
		allowance := k.allowance(ctx, from, sender, coin.Denom)
		if allowance.IsGTE(coin) {
			k.decreaseAllowance(ctx, from, sender, coin)
		} else {
			return types.ErrAmountGreaterThanAllowance(coin, allowance)
		}
	}

	return k.transfer(ctx, from, to, coins)
}

func (k *Keeper) TransferOwnership(ctx sdk.Context, owner, to sdk.AccAddress, denom string) *sdkerrors.Error {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeTransferOwnership,
			sdk.NewAttribute(types.AttributeKeyOwner, owner.String()),
			sdk.NewAttribute(types.AttributeKeyTo, to.String()),
			sdk.NewAttribute(types.AttributeKeyDenom, denom),
		),
	)

	i, err := k.getIssueIfOwner(ctx, denom, owner)
	if err != nil {
		return err
	}

	i.Owner = to.String()
	k.deleteAddressDenom(ctx, owner.String(), denom)
	k.addAddressDenom(ctx, i)
	k.setIssue(ctx, i)

	return nil
}

func (k *Keeper) Approve(ctx sdk.Context, owner, spender sdk.AccAddress, coins sdk.Coins) *sdkerrors.Error {
	for _, coin := range coins {
		k.approve(ctx, owner, spender, coin)
	}

	return nil
}

func (k *Keeper) Allowances(ctx sdk.Context, owner sdk.AccAddress, denom string) types.Allowances {
	return k.allowances(ctx, owner, denom)
}

func (k *Keeper) IncreaseAllowance(ctx sdk.Context, owner, spender sdk.AccAddress, coins sdk.Coins) *sdkerrors.Error {
	for _, coin := range coins {
		k.increaseAllowance(ctx, owner, spender, coin)
	}

	return nil
}

func (k *Keeper) DecreaseAllowance(ctx sdk.Context, owner, spender sdk.AccAddress, coins sdk.Coins) *sdkerrors.Error {
	for _, coin := range coins {
		k.decreaseAllowance(ctx, owner, spender, coin)
	}

	return nil
}

func (k *Keeper) Mint(ctx sdk.Context, minter, to sdk.AccAddress, coins sdk.Coins) *sdkerrors.Error {
	return k.mint(ctx, minter, to, coins)
}

func (k *Keeper) Burn(ctx sdk.Context, burner sdk.AccAddress, coins sdk.Coins) *sdkerrors.Error {
	return k.burn(ctx, burner, burner, coins)
}

func (k *Keeper) BurnFrom(ctx sdk.Context, burner, from sdk.AccAddress, coins sdk.Coins) *sdkerrors.Error {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeBurnFrom,
			sdk.NewAttribute(sdk.AttributeKeyAmount, coins.String()),
			sdk.NewAttribute(types.AttributeKeyBurner, burner.String()),
			sdk.NewAttribute(types.AttributeKeyFrom, from.String()),
		),
	)

	for _, coin := range coins {
		allowance := k.allowance(ctx, from, burner, coin.Denom)
		if allowance.IsGTE(coin) {
			k.decreaseAllowance(ctx, from, burner, coin)
		} else {
			return types.ErrAmountGreaterThanAllowance(coin, allowance)
		}
	}

	return k.burn(ctx, burner, from, coins)
}

func (k *Keeper) Unfreeze(ctx sdk.Context, freezer, holder sdk.AccAddress, denom, op string) *sdkerrors.Error {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUnfreeze,
			sdk.NewAttribute(types.AttributeKeyFreezer, freezer.String()),
			sdk.NewAttribute(types.AttributeKeyHolder, holder.String()),
			sdk.NewAttribute(types.AttributeKeyDenom, denom),
			sdk.NewAttribute(types.AttributeKeyOp, op),
		),
	)
	_, err := k.getIssueIfOwner(ctx, denom, freezer)
	if err != nil {
		return err
	}

	return k.freeze(ctx, holder, denom, op, false)
}

// ----------------------- fee -----------------------

func (k *Keeper) ChargeFee(ctx sdk.Context, sender sdk.AccAddress, fee sdk.Coin) *sdkerrors.Error {
	if fee.IsZero() || fee.IsNegative() {
		return nil
	}

	if err := k.csk.SendCoinsFromAccountToModule(ctx, sender, k.feeCollectorName, sdk.NewCoins(fee)); err != nil {
		return types.ErrNotEnoughFee
	}

	return nil
}
