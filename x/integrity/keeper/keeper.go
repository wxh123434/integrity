package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/wxh123434/integrity/x/integrity/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// this line is used by starport scaffolding # ibc/keeper/import
)

type (
	Keeper struct {
		cdc        codec.AminoMarshaler
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		bankKeeper types.BankKeeper
		mintKeeper types.MintKeeper
		// this line is used by starport scaffolding # ibc/keeper/attribute
	}
)

func NewKeeper(
	cdc codec.AminoMarshaler,
	storeKey,
	memKey sdk.StoreKey,
	bk types.BankKeeper,
	mk types.MintKeeper,
	// this line is used by starport scaffolding # ibc/keeper/parameter
) *Keeper {
	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		bankKeeper: bk,
		mintKeeper: mk,
		// this line is used by starport scaffolding # ibc/keeper/return
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

//
func (k Keeper) MintCoinsForHash(ctx sdk.Context, newCoins sdk.Coins) error {
	if newCoins.Empty() {
		// skip as no coins need to be minted
		return nil
	}

	return k.mintKeeper.MintCoins(ctx, newCoins)
}

//newcoins是啥

// func (k Keeper) SendCoinsFromModuleToModule(ctx sdk.Context, fees sdk.Coins) error {
// 	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, "mint", types.ModuleName, fees)
// } //从mint模块发送到此模块

func (k Keeper) SendCoinsFromMintModuleToAccount(ctx sdk.Context, amt sdk.Coins, addr sdk.AccAddress) error {
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, "mint", addr, amt)
}
