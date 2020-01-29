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

	ak types.AccountKeeper
	ck types.CoinKeeper
	sk types.SupplyKeeper
	// The reference to the FeeCollectionKeeper to add fee
	//feeCollectionKeeper FeeCollectionKeeper
}

// NewAccountKeeper returns a new sdk.AccountKeeper that uses go-amino to
// (binary) encode and decode concrete sdk.Accounts.
// nolint
func NewKeeper(
	cdc *codec.Codec,
	key sdk.StoreKey,
	ak types.AccountKeeper,
	ck types.CoinKeeper,
	sk types.SupplyKeeper,
	codespace sdk.CodespaceType) Keeper {

	return Keeper{
		key: key,
		cdc: cdc,
		//paramSubspace: paramstore.WithKeyTable(types.ParamKeyTable()),
		codespace: codespace,
		//paramsKeeper: paramsKeeper,
		ak: ak,
		ck: ck,
		sk: sk,
	}
}

func (k *Keeper) GetCodec() *codec.Codec {
	return k.cdc
}

// ----------------------- last id ----------------

func (k *Keeper) getLastId(ctx sdk.Context) (id uint64) {
	store := ctx.KVStore(k.key)
	bz := store.Get(KeyLastIssueId)
	if bz == nil {
		return types.InitLastId
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &id)
	return
}

func (k *Keeper) setLastId(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.key)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(id)
	store.Set(KeyLastIssueId, bz)
}

func (k *Keeper) incLastId(ctx sdk.Context) {
	k.setLastId(ctx, k.getLastId(ctx)+1)
}

// ----------------------- boundary denoms ----------------

func (k *Keeper) updateLeftBoundaryDenom(ctx sdk.Context, issue *types.CoinIssue) {
	store := ctx.KVStore(k.key)
	bz := store.Get(KeyFirstIssueDenom)
	if bz == nil {
		bz = k.cdc.MustMarshalBinaryLengthPrefixed(issue.Denom)
		store.Set(KeyFirstIssueDenom, bz)
	}
}

func (k *Keeper) getLeftBoundaryDenom(ctx sdk.Context) (denom string) {
	store := ctx.KVStore(k.key)
	bz := store.Get(KeyFirstIssueDenom)
	if bz == nil {
		return
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &denom)
	return
}

func (k *Keeper) updateRightBoundaryDenom(ctx sdk.Context, issue *types.CoinIssue) {
	store := ctx.KVStore(k.key)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(issue.Denom)
	store.Set(KeyLastIssueDenom, bz)
}

func (k *Keeper) getRightBoundaryDenom(ctx sdk.Context) (denom string) {
	store := ctx.KVStore(k.key)
	bz := store.Get(KeyLastIssueDenom)
	if bz == nil {
		return
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &denom)
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

func (k *Keeper) getAddressDenoms(ctx sdk.Context, accAddress string) (issues []string) {
	store := ctx.KVStore(k.key)
	bz := store.Get(KeyAddressDenoms(accAddress))
	if bz == nil {
		return []string{}
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &issues)
	return
}

func (k *Keeper) setAddressDenoms(ctx sdk.Context, accAddress string, denoms []string) {
	store := ctx.KVStore(k.key)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(denoms)
	store.Set(KeyAddressDenoms(accAddress), bz)
}

func (k *Keeper) addAddressDenoms(ctx sdk.Context, issue *types.CoinIssue) {
	denoms := k.getAddressDenoms(ctx, issue.GetOwner().String())
	denoms = append(denoms, issue.GetDenom())
	k.setAddressDenoms(ctx, issue.GetOwner().String(), denoms)
}

func (k *Keeper) deleteAllAddressDenoms(ctx sdk.Context, accAddress string) {
	store := ctx.KVStore(k.key)
	store.Delete(KeyAddressDenoms(accAddress))
}

func (k *Keeper) deleteAddressIssues(ctx sdk.Context, accAddress string, _ string) {
	denoms := k.getAddressDenoms(ctx, accAddress)
	//  todo delete denom from denoms
	k.setAddressDenoms(ctx, accAddress, denoms)
}

// ----------------------- symbol:denom pair ----------------

func (k *Keeper) getSymbolDenom(ctx sdk.Context, symbol string) (denom string) {
	store := ctx.KVStore(k.key)
	bz := store.Get(KeySymbolDenom(symbol))
	if bz == nil {
		return
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &denom)
	return
}

func (k *Keeper) setSymbolDenom(ctx sdk.Context, symbol, denom string) {
	store := ctx.KVStore(k.key)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(denom)
	store.Set(KeySymbolDenom(symbol), bz)
}

func (k *Keeper) addSymbolDenom(ctx sdk.Context, issue *types.CoinIssue) {
	k.setSymbolDenom(ctx, issue.GetSymbol(), issue.GetDenom())
}

// ----------------------- id:denom pair ----------------

func (k *Keeper) getIdDenom(ctx sdk.Context, id uint64) (denom string) {
	store := ctx.KVStore(k.key)
	bz := store.Get(KeyIdDenom(id))
	if bz == nil {
		return
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &denom)
	return
}

func (k *Keeper) setIdDenom(ctx sdk.Context, id uint64, denom string) {
	store := ctx.KVStore(k.key)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(denom)
	store.Set(KeyIdDenom(id), bz)
}

func (k *Keeper) addIdDenom(ctx sdk.Context, issue *types.CoinIssue) {
	k.setIdDenom(ctx, issue.GetId(), issue.GetDenom())
}

// ----------------------- issue -----------------------

func (k *Keeper) getIssue(ctx sdk.Context, denom string) *types.CoinIssue {
	store := ctx.KVStore(k.key)
	bz := store.Get(KeyIssuer(denom))
	if len(bz) == 0 {
		return nil
	}

	var coinIssue types.CoinIssue
	k.GetCodec().MustUnmarshalBinaryLengthPrefixed(bz, &coinIssue)
	return &coinIssue
}

func (k *Keeper) GetIssue(ctx sdk.Context, denom string) (*types.CoinIssue, sdk.Error) {
	issue := k.getIssue(ctx, denom)
	if issue == nil {
		return nil, types.ErrUnknownIssue(denom)
	}

	return issue, nil
}

func (k *Keeper) checkOwner(_ sdk.Context, issue *types.CoinIssue, owner sdk.AccAddress) (*types.CoinIssue, sdk.Error) {
	if !issue.Owner.Equals(owner) {
		return nil, types.ErrOwnerMismatch(issue.Denom)
	}

	return issue, nil
}

func (k *Keeper) getIssueIfOwner(ctx sdk.Context, denom string, owner sdk.AccAddress) (*types.CoinIssue, sdk.Error) {
	issue := k.getIssue(ctx, denom)
	if issue == nil {
		return nil, types.ErrUnknownIssue(denom)
	}

	return k.checkOwner(ctx, issue, owner)
}

func (k *Keeper) addIssue(ctx sdk.Context, issue *types.CoinIssue) {
	k.addAddressDenoms(ctx, issue)
	k.addSymbolDenom(ctx, issue)
	k.addIdDenom(ctx, issue)
	k.setIssue(ctx, issue)
	k.updateBoundaryDenoms(ctx, issue)
	k.incLastId(ctx)
}

func (k *Keeper) setIssue(ctx sdk.Context, issue *types.CoinIssue) {
	store := ctx.KVStore(k.key)
	store.Set(KeyIssuer(issue.GetDenom()), k.GetCodec().MustMarshalBinaryLengthPrefixed(issue))
}

// ----------------------- issues -----------------------

func (k *Keeper) getIssues(ctx sdk.Context, denoms []string) types.CoinIssues {
	length := len(denoms)
	issues := make(types.CoinIssues, 0, length)

	for _, v := range denoms {
		issues = append(issues, *k.getIssue(ctx, v))
	}

	return issues
}

func (k *Keeper) getIssuesByAddress(ctx sdk.Context, accAddress string) types.CoinIssues {
	return k.getIssues(ctx, k.getAddressDenoms(ctx, accAddress))
}

func (k *Keeper) iterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.key)
	//first, last := k.getBoundaryDenoms(ctx)
	//if first == last {
	//	return store.ReverseIterator(KeyIdIssuer(first), nil)
	//}
	//
	//return store.ReverseIterator(KeyIssuer(last), KeyIssuer(first))

	lastId := k.getLastId(ctx)
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

		var denom string
		k.GetCodec().MustUnmarshalBinaryLengthPrefixed(bz, &denom)
		denoms = append(denoms, denom)
	}

	return k.getIssues(ctx, denoms)
}

func (k *Keeper) List(ctx sdk.Context, params types.IssuesParams) []types.CoinIssue {
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

		var denom string
		k.GetCodec().MustUnmarshalBinaryLengthPrefixed(bz, &denom)
		denoms = append(denoms, denom)
		if len(denoms) >= params.Limit {
			break
		}
	}

	return k.getIssues(ctx, denoms)
}

// ----------------------- allowance -----------------------

func (k *Keeper) allowance(ctx sdk.Context, owner sdk.AccAddress, spender sdk.AccAddress, issueID string) sdk.Coin {
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

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeApprove,
			sdk.NewAttribute(types.AttributeKeyIssueId, amount.Denom),
			sdk.NewAttribute(types.AttributeKeyOwner, owner.String()),
			sdk.NewAttribute(types.AttributeKeySpender, spender.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, amount.Amount.String()),
		),
	)
}

func (k *Keeper) increaseAllowance(ctx sdk.Context, owner, spender sdk.AccAddress, amount sdk.Coin) {
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
}

func (k *Keeper) decreaseAllowance(ctx sdk.Context, owner, spender sdk.AccAddress, amount sdk.Coin) {
	allowance := k.allowance(ctx, owner, spender, amount.Denom)
	if allowance.IsGTE(amount) {
		k.approve(ctx, owner, spender, allowance.Sub(amount))
	} else {
		k.approve(ctx, owner, spender, sdk.NewCoin(amount.Denom, sdk.ZeroInt()))
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeDecreaseAllowance,
			sdk.NewAttribute(types.AttributeKeyIssueId, amount.Denom),
			sdk.NewAttribute(types.AttributeKeyOwner, owner.String()),
			sdk.NewAttribute(types.AttributeKeySpender, spender.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, amount.Amount.String()),
		),
	)
}

// ----------------------- transfers -----------------------

func (k *Keeper) transfer(ctx sdk.Context, from, to sdk.AccAddress, coins sdk.Coins) sdk.Error {
	if !k.ck.GetSendEnabled(ctx) {
		return bank.ErrSendDisabled(k.codespace)
	}

	if k.ck.BlacklistedAddr(to) {
		return sdk.ErrUnauthorized(fmt.Sprintf("%s is not allowed to receive transactions", to))
	}

	return k.ck.SendCoins(ctx, from, to, coins)
}

func (k *Keeper) mint(ctx sdk.Context, minter, to sdk.AccAddress, coins sdk.Coins) sdk.Error {
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

		//if coinIssueInfo.IsMintingFinished() {
		//	return nil, errors.ErrCanNotMint(issueID)
		//}
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(sdk.AttributeKeyAmount, coins.String()),
			sdk.NewAttribute(types.AttributeKeyMinter, minter.String()),
			sdk.NewAttribute(types.AttributeKeyRecipient, to.String()),
		),
	)

	if err := k.sk.MintCoins(ctx, types.ModuleName, coins); err != nil {
		return err
	}

	return k.sk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, coins)
}

func (k *Keeper) burn(ctx sdk.Context, burner, from sdk.AccAddress, coins sdk.Coins) sdk.Error {
	acc := k.ak.GetAccount(ctx, from)

	for i, coin := range coins {
		issue := k.getIssue(ctx, coin.Denom)
		issue.SubTotalSupply(coin.Amount)
		k.setIssue(ctx, issue)

		currAmt := acc.GetCoins().AmountOf(coin.Denom)
		if coin.Amount.GT(currAmt) {
			coin.Amount = currAmt
			coins = append(coins[:i], coin)
			coins = append(coins[:i+1], coins[i+1:]...)
		}
	}

	if err := k.sk.SendCoinsFromAccountToModule(ctx, from, types.ModuleName, coins); err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeBurn,
			sdk.NewAttribute(sdk.AttributeKeyAmount, coins.String()),
			sdk.NewAttribute(types.AttributeKeyBurner, burner.String()),
			sdk.NewAttribute(types.AttributeKeyFrom, from.String()),
		),
	)

	return k.sk.BurnCoins(ctx, types.ModuleName, coins)
}

// ----------------------- ERC20 -----------------------

func (k *Keeper) CreateIssue(ctx sdk.Context, owner, issuer sdk.AccAddress, params *types.IssueParams) *types.CoinIssue {
	issue := types.NewCoinIssue(owner, issuer, params)
	issue.SetId(k.getLastId(ctx))
	issue.SetIssueTime(ctx.BlockHeader().Time.Unix())

	return issue
}

func (k *Keeper) Issue(ctx sdk.Context, issue *types.CoinIssue) sdk.Error {
	k.addIssue(ctx, issue)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeIssue,
			sdk.NewAttribute(sdk.AttributeKeyAmount, issue.ToCoin().String()),
			sdk.NewAttribute(types.AttributeKeyIssuer, issue.GetIssuer().String()),
		),
	)

	if err := k.sk.MintCoins(ctx, types.ModuleName, issue.ToCoins()); err != nil {
		return err
	}

	return k.sk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, issue.GetOwner(), issue.ToCoins())
}

func (k *Keeper) Transfer(ctx sdk.Context, from, to sdk.AccAddress, coins sdk.Coins) sdk.Error {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeTransfer,
			sdk.NewAttribute(sdk.AttributeKeySender, from.String()),
		),
	)

	return k.transfer(ctx, from, to, coins)
}

func (k *Keeper) TransferFrom(ctx sdk.Context, sender, from, to sdk.AccAddress, coins sdk.Coins) sdk.Error {
	for _, coin := range coins {
		allowance := k.allowance(ctx, from, sender, coin.Denom)
		if allowance.IsGTE(coin) {
			k.decreaseAllowance(ctx, from, sender, coin)
		} else {
			return types.ErrAmountGreaterThanAllowance(coin, allowance)
		}
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

	return k.transfer(ctx, from, to, coins)
}

func (k *Keeper) Approve(ctx sdk.Context, owner, spender sdk.AccAddress, coins sdk.Coins) sdk.Error {
	for _, coin := range coins {
		k.approve(ctx, owner, spender, coin)
	}

	return nil
}

func (k *Keeper) Allowance(ctx sdk.Context, owner sdk.AccAddress, spender sdk.AccAddress, denom string) (amount sdk.Coin) {
	return k.allowance(ctx, owner, spender, denom)
}

func (k *Keeper) IncreaseAllowance(ctx sdk.Context, owner, spender sdk.AccAddress, coins sdk.Coins) sdk.Error {
	for _, coin := range coins {
		k.increaseAllowance(ctx, owner, spender, coin)
	}

	return nil
}

func (k *Keeper) DecreaseAllowance(ctx sdk.Context, owner, spender sdk.AccAddress, coins sdk.Coins) sdk.Error {
	for _, coin := range coins {
		k.decreaseAllowance(ctx, owner, spender, coin)
	}

	return nil
}

func (k *Keeper) Mint(ctx sdk.Context, minter sdk.AccAddress, coins sdk.Coins) sdk.Error {
	return k.mint(ctx, minter, minter, coins)
}

func (k *Keeper) MintTo(ctx sdk.Context, minter, to sdk.AccAddress, coins sdk.Coins) sdk.Error {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMintTo,
			sdk.NewAttribute(sdk.AttributeKeyAmount, coins.String()),
			sdk.NewAttribute(types.AttributeKeyMinter, minter.String()),
			sdk.NewAttribute(types.AttributeKeyTo, to.String()),
		),
	)

	return k.mint(ctx, minter, to, coins)
}

func (k *Keeper) Burn(ctx sdk.Context, burner sdk.AccAddress, coins sdk.Coins) sdk.Error {
	return k.burn(ctx, burner, burner, coins)
}

func (k *Keeper) BurnFrom(ctx sdk.Context, burner, from sdk.AccAddress, coins sdk.Coins) sdk.Error {
	for _, coin := range coins {
		allowance := k.allowance(ctx, from, burner, coin.Denom)
		if allowance.IsGTE(coin) {
			k.decreaseAllowance(ctx, from, burner, coin)
		} else {
			return types.ErrAmountGreaterThanAllowance(coin, allowance)
		}
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeBurnFrom,
			sdk.NewAttribute(sdk.AttributeKeyAmount, coins.String()),
			sdk.NewAttribute(types.AttributeKeyBurner, burner.String()),
			sdk.NewAttribute(types.AttributeKeyFrom, from.String()),
		),
	)

	return k.burn(ctx, burner, from, coins)
}
