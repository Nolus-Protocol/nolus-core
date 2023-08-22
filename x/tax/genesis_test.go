package tax_test

import (
	"testing"

	keepertest "github.com/Nolus-Protocol/nolus-core/testutil/keeper"
	"github.com/Nolus-Protocol/nolus-core/testutil/nullify"
	"github.com/Nolus-Protocol/nolus-core/x/tax"
	"github.com/Nolus-Protocol/nolus-core/x/tax/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k, ctx := keepertest.TaxKeeper(t)
	tax.InitGenesis(ctx, *k, genesisState)
	got := tax.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
