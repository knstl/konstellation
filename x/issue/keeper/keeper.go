package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

//Keys add
//Add a issue
func (k Keeper) AddIssue(ctx sdk.Context, issue *types.CoinIssue) {
	k.addAddressIssues(ctx, issue)

	issueIDs := k.GetSymbolIssues(ctx, issue.GetSymbol())
	issueIDs = append(issueIDs, issue.GetIssueId())
	k.setSymbolIssues(ctx, issue.GetSymbol(), issueIDs)

	store := ctx.KVStore(k.key)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(issue)
	store.Set(KeyIssuer(issue.GetIssueId()), bz)
}

//Create a issue
func (k *Keeper) CreateIssue(ctx sdk.Context, issue *types.CoinIssue) sdk.Error {
	store := ctx.KVStore(k.key)
	id, err := k.getNewIssueID(store)
	if err != nil {
		return err
	}
	issue.IssueTime = ctx.BlockHeader().Time.Unix()
	issue.IssueId = KeyIssueIdStr(id)

	k.AddIssue(ctx, issue)

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

	return err
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
