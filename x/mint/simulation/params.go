package simulation

// DONTCOVER

// import (
// 	"fmt"
// 	"math/rand"

// 	"github.com/Nolus-Protocol/nolus-core/x/mint/types"
// 	"github.com/cosmos/cosmos-sdk/x/simulation"

// 	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
// )

// refactor: fix when simulation refactor is done
// // ParamChanges defines the parameters that can be modified by param change proposals
// // on the simulation.
// func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
// 	return []simtypes.ParamChange{
// 		simulation.NewSimParamChange(types.ModuleName, string(types.KeyMaxMintableNanoseconds),
// 			func(r *rand.Rand) string {
// 				return fmt.Sprintf("\"%s\"", GenMaxMintableNanoseconds(r).String())
// 			},
// 		),
// 	}
// }
