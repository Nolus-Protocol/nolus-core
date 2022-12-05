package simapp

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	"github.com/tendermint/tendermint/libs/json"

	"github.com/cosmos/cosmos-sdk/simapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/tendermint/spm/cosmoscmd"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
	tmdb "github.com/tendermint/tm-db"

	"gitlab-nomo.credissimo.net/nomo/nolus-core/app"
)

// New creates application instance with in-memory database and disabled logging.
func New(dir string, withDefaultGenesisState bool) cosmoscmd.App {
	db := tmdb.NewMemDB()
	logger := log.NewNopLogger()

	encoding := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)

	a := app.New(logger, db, nil, true, map[int64]bool{}, dir, 0, encoding,
		simapp.EmptyAppOptions{})
	// InitChain updates deliverState which is required when app.NewContext is called
	genState := []byte("{}")
	if withDefaultGenesisState {
		genStateObj := NewDefaultGenesisState(encoding.Marshaler)
		state, err := json.MarshalIndent(genStateObj, "", " ")
		if err != nil {
			panic(err)
		}
		genState = state
	}
	a.InitChain(abci.RequestInitChain{
		ConsensusParams: defaultConsensusParams,
		AppStateBytes:   genState,
	})
	return a
}

func TestSetup() (*app.App, error) {
	rootApp := New(app.DefaultNodeHome, true)
	nolusApp, ok := rootApp.(*app.App)
	if !ok {
		return nil, fmt.Errorf("invalid simapp created: %v", ok)
	}
	return nolusApp, nil
}

// NewDefaultGenesisState generates the default state for the application.
func NewDefaultGenesisState(cdc codec.JSONCodec) app.GenesisState {
	return app.ModuleBasics.DefaultGenesis(cdc)
}

var defaultConsensusParams = &abci.ConsensusParams{
	Block: &abci.BlockParams{
		MaxBytes: 200000,
		MaxGas:   2000000,
	},
	Evidence: &tmproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &tmproto.ValidatorParams{
		PubKeyTypes: []string{
			tmtypes.ABCIPubKeyTypeEd25519,
		},
	},
}

// NewAppConstructor returns a new simapp AppConstructor
func NewAppConstructor() network.AppConstructor {
	encoding := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)

	return func(val network.Validator) servertypes.Application {
		return app.New(val.Ctx.Logger, tmdb.NewMemDB(), nil, true, map[int64]bool{}, val.Ctx.Config.RootDir, 0, encoding,
			simapp.EmptyAppOptions{},
			baseapp.SetPruning(storetypes.NewPruningOptionsFromString(val.AppConfig.Pruning)),
			baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
		)
	}
}
