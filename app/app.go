package app

import (
	"encoding/json"
	"os"

	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
)

const appName = "konstellation"

var (
	// Default home directories for the application CLI
	DefaultCLIHome = os.ExpandEnv("$HOME/.konstellationcli")

	// DefaultNodeHome sets the folder where the application data and configuration will be stored
	DefaultNodeHome = os.ExpandEnv("$HOME/.konstellation")

	// ModuleBasicManager is in charge of setting up basic module elements
	ModuleBasics = module.NewBasicManager(
		genaccounts.AppModuleBasic{},
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		params.AppModuleBasic{},
		staking.AppModuleBasic{},
		distribution.AppModuleBasic{},
		slashing.AppModuleBasic{},
		supply.AppModuleBasic{},
	)

	// Account permissions
	maccPerms = map[string][]string{
		auth.FeeCollectorName:   nil,
		distribution.ModuleName: nil,
		staking.BondedPoolName: {
			supply.Burner,
			supply.Staking,
		},
		staking.NotBondedPoolName: {
			supply.Burner,
			supply.Staking,
		},
	}
)

// MakeCodec generates the necessary codecs for Amino
func MakeCodec() *codec.Codec {
	var cdc = codec.New()
	ModuleBasics.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	return cdc
}

type KonstellationApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	// Keys to access the substores
	keyMain         *sdk.KVStoreKey
	keyAccount      *sdk.KVStoreKey
	keySupply       *sdk.KVStoreKey
	keyStaking      *sdk.KVStoreKey
	tkeyStaking     *sdk.TransientStoreKey
	keyDistribution *sdk.KVStoreKey
	// tkeyDistribution *sdk.TransientStoreKey
	keyParams   *sdk.KVStoreKey
	tkeyParams  *sdk.TransientStoreKey
	keySlashing *sdk.KVStoreKey

	// Keepers
	accountKeeper      auth.AccountKeeper
	bankKeeper         bank.Keeper
	supplyKeeper       supply.Keeper
	stakingKeeper      staking.Keeper
	slashingKeeper     slashing.Keeper
	distributionKeeper distribution.Keeper
	paramsKeeper       params.Keeper

	// Module Manager
	mm *module.Manager
}

// NewKonstellationApp is a constructor function for KonstellationApp
func NewKonstellationApp(logger log.Logger, db dbm.DB) *KonstellationApp {

	// First define the top level codec that will be shared by the different modules
	cdc := MakeCodec()

	// BaseApp handles interactions with Tendermint through the ABCI protocol
	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

	// Here you initialize your application with the store keys it requires
	var app = &KonstellationApp{
		BaseApp: bApp,
		cdc:     cdc,

		keyMain:         sdk.NewKVStoreKey(bam.MainStoreKey),
		keyAccount:      sdk.NewKVStoreKey(auth.StoreKey),
		keySupply:       sdk.NewKVStoreKey(supply.StoreKey),
		keyStaking:      sdk.NewKVStoreKey(staking.StoreKey),
		tkeyStaking:     sdk.NewTransientStoreKey(staking.TStoreKey),
		keyDistribution: sdk.NewKVStoreKey(distribution.StoreKey),
		// tkeyDistribution: sdk.NewTransientStoreKey(distribution.TStoreKey),
		keyParams:   sdk.NewKVStoreKey(params.StoreKey),
		tkeyParams:  sdk.NewTransientStoreKey(params.TStoreKey),
		keySlashing: sdk.NewKVStoreKey(slashing.StoreKey),
	}

	// The ParamsKeeper handles parameter storage for the application
	app.paramsKeeper = params.NewKeeper(
		app.cdc,
		app.keyParams,
		app.tkeyParams,
		params.DefaultCodespace,
	)

	// The AccountKeeper handles address -> account lookups
	app.accountKeeper = auth.NewAccountKeeper(
		app.cdc,
		app.keyAccount,
		app.paramsKeeper.Subspace(auth.DefaultParamspace),
		auth.ProtoBaseAccount,
	)

	// The BankKeeper allows you perform sdk.Coins interactions
	app.bankKeeper = bank.NewBaseKeeper(
		app.accountKeeper,
		app.paramsKeeper.Subspace(bank.DefaultParamspace),
		bank.DefaultCodespace,
		map[string]bool{},
	)

	// The SupplyKeeper collects transaction fees and renders them to the fee distribution module
	app.supplyKeeper = supply.NewKeeper(
		app.cdc,
		app.keySupply,
		app.accountKeeper,
		app.bankKeeper,
		// supply.DefaultCodespace,
		maccPerms,
	)

	// The staking keeper
	stakingKeeper := staking.NewKeeper(
		app.cdc,
		app.keyStaking,
		app.tkeyStaking,
		app.supplyKeeper,
		app.paramsKeeper.Subspace(staking.DefaultParamspace),
		staking.DefaultCodespace,
	)

	app.distributionKeeper = distribution.NewKeeper(
		app.cdc,
		app.keyDistribution,
		app.paramsKeeper.Subspace(distribution.DefaultParamspace),
		&stakingKeeper,
		app.supplyKeeper,
		distribution.DefaultCodespace,
		auth.FeeCollectorName,
		map[string]bool{},
	)

	app.slashingKeeper = slashing.NewKeeper(
		app.cdc,
		app.keySlashing,
		&stakingKeeper,
		app.paramsKeeper.Subspace(slashing.DefaultParamspace),
		slashing.DefaultCodespace,
	)

	app.stakingKeeper = *stakingKeeper.SetHooks(
		staking.NewMultiStakingHooks(
			app.distributionKeeper.Hooks(),
			app.slashingKeeper.Hooks()),
	)

	app.mm = module.NewManager(
		genaccounts.NewAppModule(app.accountKeeper),
		genutil.NewAppModule(app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(app.accountKeeper),
		bank.NewAppModule(app.bankKeeper, app.accountKeeper),
		supply.NewAppModule(app.supplyKeeper, app.accountKeeper),
		distribution.NewAppModule(app.distributionKeeper, app.supplyKeeper),
		slashing.NewAppModule(app.slashingKeeper, app.stakingKeeper),
		staking.NewAppModule(app.stakingKeeper, app.distributionKeeper, app.accountKeeper, app.supplyKeeper),
	)

	app.mm.SetOrderBeginBlockers(distribution.ModuleName, slashing.ModuleName)
	app.mm.SetOrderEndBlockers(staking.ModuleName)

	// Sets the order of Genesis - Order matters, genutil is to always come last
	app.mm.SetOrderInitGenesis(
		genaccounts.ModuleName,
		distribution.ModuleName,
		staking.ModuleName,
		auth.ModuleName,
		bank.ModuleName,
		slashing.ModuleName,
		genutil.ModuleName,
	)

	// register all module routes and module queriers
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())

	// The initChainer handles translating the genesis.json file into initial state for the network
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	// The AnteHandler handles signature verification and transaction pre-processing
	app.SetAnteHandler(
		auth.NewAnteHandler(
			app.accountKeeper,
			app.supplyKeeper,
			auth.DefaultSigVerificationGasConsumer,
		),
	)

	app.MountStores(
		app.keyMain,
		app.keyAccount,
		app.keySupply,
		app.keyStaking,
		app.tkeyStaking,
		app.keyDistribution,
		// app.tkeyDistribution,
		app.keySlashing,
		app.keyParams,
		app.tkeyParams,
	)

	err := app.LoadLatestVersion(app.keyMain)
	if err != nil {
		cmn.Exit(err.Error())
	}

	return app
}

// GenesisState represents chain state at the start of the chain. Any initial state (account balances) are stored here.
type GenesisState map[string]json.RawMessage

// func NewDefaultGenesisState() GenesisState {
// 	return ModuleBasics.DefaultGenesis()
// }

func (app *KonstellationApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState

	err := app.cdc.UnmarshalJSON(req.AppStateBytes, &genesisState)
	if err != nil {
		panic(err)
	}

	return app.mm.InitGenesis(ctx, genesisState)
}

func (app *KonstellationApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}
func (app *KonstellationApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}
func (app *KonstellationApp) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keyMain)
}

// _________________________________________________________

func (app *KonstellationApp) ExportAppStateAndValidators(forZeroHeight bool, jailWhiteList []string,
) (appState json.RawMessage, validators []tmtypes.GenesisValidator, err error) {

	// as if they could withdraw from the start of the next block
	ctx := app.NewContext(true, abci.Header{Height: app.LastBlockHeight()})

	genState := app.mm.ExportGenesis(ctx)
	appState, err = codec.MarshalJSONIndent(app.cdc, genState)
	if err != nil {
		return nil, nil, err
	}

	validators = staking.WriteValidators(ctx, app.stakingKeeper)

	return appState, validators, nil
}
