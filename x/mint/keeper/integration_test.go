package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/stretchr/testify/require"

	"github.com/Nolus-Protocol/nolus-core/app/params"
	"github.com/Nolus-Protocol/nolus-core/testutil/simapp"
	"github.com/Nolus-Protocol/nolus-core/x/mint/types"
)

// returns context and an app with updated mint keeper.
func TestSetAndRetrieveParamsAndMinter(t *testing.T) {
	denom := "unls"
	maxMintableNanoseconds := uint64(2000)

	params.SetAddressPrefixes()
	app, err := simapp.TestSetup(t)
	require.NoError(t, err)

	isCheckTx := false
	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})

	_ = app.MintKeeper.SetParams(ctx, types.NewParams(denom, sdkmath.NewUint(maxMintableNanoseconds)))
	app.MintKeeper.SetMinter(ctx, types.DefaultInitialMinter())

	require.Equal(t, denom, app.MintKeeper.GetParams(ctx).MintDenom)
	require.Equal(t, sdkmath.NewUint(maxMintableNanoseconds), app.MintKeeper.GetParams(ctx).MaxMintableNanoseconds)
	require.Equal(t, types.DefaultInitialMinter(), app.MintKeeper.GetMinter(ctx))
}
