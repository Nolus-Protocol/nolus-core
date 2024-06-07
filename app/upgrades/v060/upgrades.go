package v060

import (
	"context"
	"fmt"

	"github.com/Nolus-Protocol/nolus-core/app/keepers"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	keepers *keepers.AppKeepers,
	codec codec.Codec,
) upgradetypes.UpgradeHandler {
	return func(c context.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx := sdk.UnwrapSDKContext(c)

		ctx.Logger().Info("Starting module migrations...")
		vm, err := mm.RunMigrations(ctx, configurator, vm)
		if err != nil {
			return vm, err
		}

		// TODO
		// ctx.Logger().Info("Setting consensus params...")
		// err = enableVoteExtensions(ctx, keepers.ConsensusKeeper)
		// if err != nil {
		// 	return nil, err
		// }

		ctx.Logger().Info(`
$$\   $$\           $$\                                       $$$$$$\      $$$$$$\      $$$$$$\  
$$$\  $$ |          $$ |                                     $$$ __$$\    $$  __$$\    $$$ __$$\ 
$$$$\ $$ | $$$$$$\  $$ |$$\   $$\  $$$$$$$\       $$\    $$\ $$$$\ $$ |   $$ /  \__|   $$$$\ $$ |
$$ $$\$$ |$$  __$$\ $$ |$$ |  $$ |$$  _____|      \$$\  $$  |$$\$$\$$ |   $$$$$$$\     $$\$$\$$ |
$$ \$$$$ |$$ /  $$ |$$ |$$ |  $$ |\$$$$$$\         \$$\$$  / $$ \$$$$ |   $$  __$$\    $$ \$$$$ |
$$ |\$$$ |$$ |  $$ |$$ |$$ |  $$ | \____$$\         \$$$  /  $$ |\$$$ |   $$ /  $$ |   $$ |\$$$ |
$$ | \$$ |\$$$$$$  |$$ |\$$$$$$  |$$$$$$$  |         \$  /   \$$$$$$  /$$\ $$$$$$  |$$\\$$$$$$  /
\__|  \__| \______/ \__| \______/ \_______/           \_/     \______/ \__|\______/ \__|\______/     
																											   
		
$$$$$$$$\      $$\                                                                               
$$  _____|     $$ |                                                                              
$$ |      $$$$$$$ | $$$$$$\  $$$$$$$\                                                            
$$$$$\   $$  __$$ |$$  __$$\ $$  __$$\                                                           
$$  __|  $$ /  $$ |$$$$$$$$ |$$ |  $$ |                                                          
$$ |     $$ |  $$ |$$   ____|$$ |  $$ |                                                          
$$$$$$$$\\$$$$$$$ |\$$$$$$$\ $$ |  $$ |                                                          
\________|\_______| \_______|\__|  \__|
														 
  $$\ $$\   $$\       $$$$$$$$\ $$\       
  $$ \$$ \  $$ |      $$  _____|$$ |      
$$$$$$$$$$\ $$ |      $$ |      $$ |      
\_$$  $$   |$$ |      $$$$$\    $$ |      
$$$$$$$$$$\ $$ |      $$  __|   $$ |      
\_$$  $$  _|$$ |      $$ |      $$ |      
  $$ |$$ |  $$$$$$$$\ $$ |      $$$$$$$$\ 
  \__|\__|  \________|\__|      \________|
`)

		ctx.Logger().Info(fmt.Sprintf("Migration {%s} applied", UpgradeName))
		return vm, nil
	}
}

// func enableVoteExtensions(ctx sdk.Context, consensusKeeper *consensuskeeper.Keeper) error {
// 	oldParams, err := consensusKeeper.Params(ctx, &types.QueryParamsRequest{})
// 	if err != nil {
// 		return err
// 	}

// 	// we need to enable VoteExtensions for Slinky
// 	oldParams.Params.Abci = &comettypes.ABCIParams{VoteExtensionsEnableHeight: ctx.BlockHeight() + 4}

// 	updateParamsMsg := types.MsgUpdateParams{
// 		Authority: authtypes.NewModuleAddress(adminmoduletypes.ModuleName).String(),
// 		Block:     oldParams.Params.Block,
// 		Evidence:  oldParams.Params.Evidence,
// 		Validator: oldParams.Params.Validator,
// 		Abci:      oldParams.Params.Abci,
// 	}

// 	_, err = consensusKeeper.UpdateParams(ctx, &updateParamsMsg)
// 	return err
// }
