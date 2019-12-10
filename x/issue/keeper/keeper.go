package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/params/subspace"
	"github.com/konstellation/konstellation/x/issue/types"
)

// IssueKeeper encodes/decodes accounts using the go-amino (binary)
// encoding/decoding library.
type Keeper struct {
	// The (unexposed) key used to access the store from the Context.
	key sdk.StoreKey

	// The codec codec for binary encoding/decoding of accounts.
	cdc *codec.Codec

	// The reference to the Paramstore to get and set issue specific params
	paramSubspace subspace.Subspace

	// Reserved codespace
	codespace sdk.CodespaceType

	// The reference to the Param Keeper to get and set Global Params
	paramsKeeper params.Keeper
	// The reference to the CoinKeeper to modify balances

	ck types.CoinKeeper
	sk types.SupplyKeeper
	// The reference to the FeeCollectionKeeper to add fee
	//feeCollectionKeeper FeeCollectionKeeper
}

// NewAccountKeeper returns a new sdk.AccountKeeper that uses go-amino to
// (binary) encode and decode concrete sdk.Accounts.
// nolint
func NewKeeper(
	cdc *codec.Codec, key sdk.StoreKey, ck types.CoinKeeper, sk types.SupplyKeeper, codespace sdk.CodespaceType) Keeper {

	return Keeper{
		key: key,
		cdc: cdc,
		//paramSubspace: paramstore.WithKeyTable(types.ParamKeyTable()),
		codespace: codespace,
		//paramsKeeper: paramsKeeper,
		ck: ck,
		sk: sk,
	}
}

func (k Keeper) GetCodec() *codec.Codec {
	return k.cdc
}

//Set address
func (k Keeper) setAddressIssues(ctx sdk.Context, accAddress string, issueIDs []string) {
	store := ctx.KVStore(k.key)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(issueIDs)
	store.Set(KeyAddressIssues(accAddress), bz)
}

func (k Keeper) deleteAddressIssues(ctx sdk.Context, accAddress string) {
	store := ctx.KVStore(k.key)
	store.Delete(KeyAddressIssues(accAddress))
}

//Add address
func (k Keeper) addAddressIssues(ctx sdk.Context, issue *types.CoinIssue) {
	issueIDs := k.GetAddressIssues(ctx, issue.GetOwner().String())
	issueIDs = append(issueIDs, issue.GetIssueId())
	k.setAddressIssues(ctx, issue.GetOwner().String(), issueIDs)
}

//Set symbol
func (k Keeper) setSymbolIssues(ctx sdk.Context, symbol string, issueIDs []string) {
	store := ctx.KVStore(k.key)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(issueIDs)
	store.Set(KeySymbolIssues(symbol), bz)
}

//Get address from a issue
func (k Keeper) GetAddressIssues(ctx sdk.Context, accAddress string) (issueIDs []string) {
	store := ctx.KVStore(k.key)
	bz := store.Get(KeyAddressIssues(accAddress))
	if bz == nil {
		return []string{}
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &issueIDs)
	return issueIDs
}

//Get issueIDs from a issue
func (k Keeper) GetSymbolIssues(ctx sdk.Context, symbol string) (issueIDs []string) {
	store := ctx.KVStore(k.key)
	bz := store.Get(KeySymbolIssues(symbol))
	if bz == nil {
		return []string{}
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &issueIDs)
	return issueIDs
}

//Returns issue by issueID
func (k Keeper) GetIssue(ctx sdk.Context, issueID string) *types.CoinIssue {
	store := ctx.KVStore(k.key)
	bz := store.Get(KeyIssuer(issueID))
	if len(bz) == 0 {
		return nil
	}
	var coinIssue types.CoinIssue
	k.GetCodec().MustUnmarshalBinaryLengthPrefixed(bz, &coinIssue)
	return &coinIssue
}

//Returns issues by accAddress
func (k Keeper) GetIssues(ctx sdk.Context, accAddress string) types.CoinIssues {
	issueIDs := k.GetAddressIssues(ctx, accAddress)
	length := len(issueIDs)
	if length == 0 {
		return nil
	}
	issues := make(types.CoinIssues, 0, length)
	for _, v := range issueIDs {
		issues = append(issues, *k.GetIssue(ctx, v))
	}

	return issues
}

// Gets the next available issueID and increments it
func (k Keeper) getNewIssueID(store sdk.KVStore) (issueID uint64, err sdk.Error) {
	bz := store.Get(KeyNextIssueID)
	if bz == nil {
		bz = k.cdc.MustMarshalBinaryLengthPrefixed(1)
		//return 0, sdk.NewError(k.codespace, types.CodeInvalidGenesis, "InitialIssueID never set")
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &issueID)
	bz = k.cdc.MustMarshalBinaryLengthPrefixed(issueID + 1)
	store.Set(KeyNextIssueID, bz)
	return issueID, nil
}

// Get issue id and return
func (k Keeper) ResolveNextIssueID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.key)
	id, _ := k.getNewIssueID(store)

	return id
}

func (k Keeper) addIssue(ctx sdk.Context, issue *types.CoinIssue) {
	k.addAddressIssues(ctx, issue)

	issueIDs := k.GetSymbolIssues(ctx, issue.GetSymbol())
	issueIDs = append(issueIDs, issue.GetIssueId())
	k.setSymbolIssues(ctx, issue.GetSymbol(), issueIDs)

	store := ctx.KVStore(k.key)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(issue)
	store.Set(KeyIssuer(issue.GetIssueId()), bz)
}

func (k Keeper) allowance(ctx sdk.Context, owner sdk.AccAddress, spender sdk.AccAddress, issueID string) sdk.Coin {
	store := ctx.KVStore(k.key)
	bz := store.Get(KeyAllowance(issueID, owner, spender))
	if bz == nil {
		return sdk.NewCoin(issueID, sdk.ZeroInt())
	}
	var amount sdk.Int
	k.GetCodec().MustUnmarshalBinaryLengthPrefixed(bz, &amount)
	return sdk.NewCoin(issueID, amount)
}

func (k *Keeper) approve(ctx sdk.Context, owner, spender sdk.AccAddress, amount sdk.Coin) {
	store := ctx.KVStore(k.key)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(amount.Amount)
	store.Set(KeyAllowance(amount.Denom, owner, spender), bz)
}

func (k *Keeper) decreaseAllowance(ctx sdk.Context, owner, spender sdk.AccAddress, amount sdk.Coin) {
	allowance := k.allowance(ctx, owner, spender, amount.Denom)
	if allowance.IsGTE(amount) {
		k.approve(ctx, owner, spender, allowance.Sub(amount))
	} else {
		k.approve(ctx, owner, spender, sdk.NewCoin(amount.Denom, sdk.ZeroInt()))
	}
}

func (k *Keeper) transfer(ctx sdk.Context, from, to sdk.AccAddress, coins sdk.Coins) sdk.Error {
	if !k.ck.GetSendEnabled(ctx) {
		return bank.ErrSendDisabled(k.codespace)
	}

	if k.ck.BlacklistedAddr(to) {
		return sdk.ErrUnauthorized(fmt.Sprintf("%s is not allowed to receive transactions", to))
	}

	return k.ck.SendCoins(ctx, from, to, coins)
}

//Create a issue
func (k *Keeper) CreateIssue(ctx sdk.Context, owner, issuer sdk.AccAddress, params *types.IssueParams) *types.CoinIssue {
	id := k.ResolveNextIssueID(ctx)
	issue := types.NewCoinIssue(owner, issuer, params)
	issue.IssueTime = ctx.BlockHeader().Time.Unix()
	issue.IssueId = KeyIssueIdStr(id)

	return issue
}

//Create a issue
func (k *Keeper) Issue(ctx sdk.Context, issue *types.CoinIssue) sdk.Error {
	k.addIssue(ctx, issue)

	if err := k.sk.MintCoins(ctx, types.ModuleName, issue.ToCoins()); err != nil {
		return err
	}

	if err := k.sk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, issue.GetOwner(), issue.ToCoins()); err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeIssue,
			sdk.NewAttribute(sdk.AttributeKeyAmount, issue.ToCoin().String()),
			sdk.NewAttribute(types.AttributeKeyIssuer, issue.GetIssuer().String()),
		),
	)

	return nil
}

func (k *Keeper) Transfer(ctx sdk.Context, from, to sdk.AccAddress, coins sdk.Coins) sdk.Error {
	if err := k.transfer(ctx, from, to, coins); err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeTransfer,
			sdk.NewAttribute(sdk.AttributeKeySender, from.String()),
		),
	)

	return nil
}

func (k *Keeper) TransferFrom(ctx sdk.Context, sender, from, to sdk.AccAddress, coins sdk.Coins) sdk.Error {
	for i, coin := range coins {
		allowance := k.allowance(ctx, from, sender, coin.Denom)
		if allowance.IsGTE(coin) {
			k.decreaseAllowance(ctx, from, sender, coin)
		} else {
			coins = append(coins[:i], coins[i+1:]...)
		}
	}

	if err := k.transfer(ctx, from, to, coins); err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeTransferFrom,
			sdk.NewAttribute(types.AttributeKeySpender, sender.String()),
			sdk.NewAttribute(sdk.AttributeKeySender, from.String()),
			sdk.NewAttribute(types.AttributeKeyRecipient, to.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, coins.String()),
		),
	)

	return nil
}

func (k *Keeper) Approve(ctx sdk.Context, owner, spender sdk.AccAddress, amount sdk.Coin) sdk.Error {
	k.approve(ctx, owner, spender, amount)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeApprove,
			sdk.NewAttribute(types.AttributeKeyIssueId, amount.Denom),
			sdk.NewAttribute(types.AttributeKeyOwner, owner.String()),
			sdk.NewAttribute(types.AttributeKeySpender, spender.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, amount.Amount.String()),
		),
	)

	return nil
}

func (k *Keeper) IncreaseAllowance(ctx sdk.Context, owner, spender sdk.AccAddress, amount sdk.Coin) sdk.Error {
	allowance := k.allowance(ctx, owner, spender, amount.Denom)
	k.approve(ctx, owner, spender, allowance.Add(amount))

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeIncreaseAllowance,
			sdk.NewAttribute(types.AttributeKeyIssueId, amount.Denom),
			sdk.NewAttribute(types.AttributeKeyOwner, owner.String()),
			sdk.NewAttribute(types.AttributeKeySpender, spender.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, amount.Amount.String()),
		),
	)

	return nil
}

func (k *Keeper) DecreaseAllowance(ctx sdk.Context, owner, spender sdk.AccAddress, amount sdk.Coin) sdk.Error {
	k.decreaseAllowance(ctx, owner, spender, amount)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeDecreaseAllowance,
			sdk.NewAttribute(types.AttributeKeyIssueId, amount.Denom),
			sdk.NewAttribute(types.AttributeKeyOwner, owner.String()),
			sdk.NewAttribute(types.AttributeKeySpender, spender.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, amount.Amount.String()),
		),
	)

	return nil
}

func (k Keeper) Allowance(ctx sdk.Context, owner sdk.AccAddress, spender sdk.AccAddress, issueID string) (amount sdk.Coin) {
	return k.allowance(ctx, owner, spender, issueID)
}

func (k Keeper) Iterator(ctx sdk.Context, startIssueId string) sdk.Iterator {
	store := ctx.KVStore(k.key)
	endIssueId := startIssueId

	if startIssueId == "" {
		endIssueId = KeyIssueIdStr(types.CoinIssueMaxId)
		startIssueId = KeyIssueIdStr(types.CoinIssueMinId - 1)
	} else {
		startIssueId = KeyIssueIdStr(types.CoinIssueMinId - 1)
	}

	iterator := store.ReverseIterator(KeyIssuer(startIssueId), KeyIssuer(endIssueId))
	return iterator
}

func (k Keeper) ListAll(ctx sdk.Context) types.CoinIssues {
	iterator := k.Iterator(ctx, "")
	defer iterator.Close()

	list := make(types.CoinIssues, 0)
	for ; iterator.Valid(); iterator.Next() {
		bz := iterator.Value()
		if len(bz) == 0 {
			continue
		}

		var issue types.CoinIssue
		k.GetCodec().MustUnmarshalBinaryLengthPrefixed(bz, &issue)
		list = append(list, issue)
	}
	return list
}

func (k Keeper) List(ctx sdk.Context, params types.IssuesParams) []types.CoinIssue {
	if params.Owner != "" {
		return k.GetIssues(ctx, params.Owner)
	}

	iterator := k.Iterator(ctx, params.StartIssueId)
	defer iterator.Close()

	list := make(types.CoinIssues, 0, params.Limit)
	for ; iterator.Valid(); iterator.Next() {
		bz := iterator.Value()
		if len(bz) == 0 {
			continue
		}

		var issue types.CoinIssue
		k.GetCodec().MustUnmarshalBinaryLengthPrefixed(bz, &issue)
		list = append(list, issue)
		if len(list) >= params.Limit {
			break
		}
	}
	return list
}
