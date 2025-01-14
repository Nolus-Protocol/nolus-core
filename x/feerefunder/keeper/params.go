package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Nolus-Protocol/nolus-core/x/feerefunder/types"
)

// GetParams get all parameters as types.Params.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params, err error) {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := store.Get(types.ParamsKey)
	if err != nil {
		return params, err
	}
	if bz == nil {
		return params, nil
	}

	err = k.cdc.Unmarshal(bz, &params)
	return params, err
}

// SetParams set the params.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := k.cdc.Marshal(&params)
	if err != nil {
		return err
	}

	_ = store.Set(types.ParamsKey, bz)
	return nil
}
