package keeper_test

import (
	"testing"

	"github.com/Nolus-Protocol/nolus-core/app/params"
	testkeeper "github.com/Nolus-Protocol/nolus-core/testutil/keeper"
	"github.com/Nolus-Protocol/nolus-core/x/tax/types"
	"github.com/stretchr/testify/require"
)

func (s *KeeperTestSuite) TestParams() {
	testCases := []struct {
		name      string
		input     types.Params
		expectErr bool
	}{
		{
			name: "set invalid params",
			input: types.Params{
				FeeRate:         0,
				ContractAddress: "a",
				BaseDenom:       "1",
			},
			expectErr: true,
		},
		{
			name: "set full valid params",
			input: types.Params{
				FeeRate:         1,
				ContractAddress: "nolus14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9s0k0puz",
				BaseDenom:       "nolus",
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.SetupTest(false)

		s.Run(tc.name, func() {
			expected := s.app.TaxKeeper.GetParams(s.ctx)
			err := s.app.TaxKeeper.SetParams(s.ctx, tc.input)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				expected = tc.input
				s.Require().NoError(err)
			}

			p := s.app.TaxKeeper.GetParams(s.ctx)
			s.Require().Equal(expected, p)
		})
	}
}

func TestGetParams(t *testing.T) {
	params.SetAddressPrefixes()
	k, ctx := testkeeper.TaxKeeper(t)
	params := types.DefaultParams()

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.FeeRate, k.FeeRate(ctx))
	require.EqualValues(t, params.ContractAddress, k.ContractAddress(ctx))
	require.EqualValues(t, params.BaseDenom, k.BaseDenom(ctx))
}
