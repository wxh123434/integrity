package keeper

import (
	"testing"

	"github.com/SBC/integrity/x/integrity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

func createNDatahash(keeper *Keeper, ctx sdk.Context, n int) []types.Datahash {
	items := make([]types.Datahash, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendDatahash(ctx, items[i])
	}
	return items
}

func TestDatahashGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNDatahash(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetDatahash(ctx, item.Id))
	}
}

func TestDatahashExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNDatahash(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasDatahash(ctx, item.Id))
	}
}

func TestDatahashRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNDatahash(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDatahash(ctx, item.Id)
		assert.False(t, keeper.HasDatahash(ctx, item.Id))
	}
}

func TestDatahashGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNDatahash(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllDatahash(ctx))
}

func TestDatahashCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNDatahash(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetDatahashCount(ctx))
}
