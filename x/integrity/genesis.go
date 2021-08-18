package integrity

import (
	"github.com/SBC/integrity/x/integrity/keeper"
	"github.com/SBC/integrity/x/integrity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the datahash
	for _, elem := range genState.DatahashList {
		k.SetDatahash(ctx, *elem)
	}

	// Set datahash count
	k.SetDatahashCount(ctx, genState.DatahashCount)

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all datahash
	datahashList := k.GetAllDatahash(ctx)
	for _, elem := range datahashList {
		elem := elem
		genesis.DatahashList = append(genesis.DatahashList, &elem)
	}

	// Set the current count
	genesis.DatahashCount = k.GetDatahashCount(ctx)

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
