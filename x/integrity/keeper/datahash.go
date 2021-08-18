package keeper

import (
	"encoding/binary"
	"github.com/SBC/integrity/x/integrity/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

// GetDatahashCount get the total number of datahash
func (k Keeper) GetDatahashCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DatahashCountKey))
	byteKey := types.KeyPrefix(types.DatahashCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to iint64
		panic("cannot decode count")
	}

	return count
}

// SetDatahashCount set the total number of datahash
func (k Keeper) SetDatahashCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DatahashCountKey))
	byteKey := types.KeyPrefix(types.DatahashCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendDatahash appends a datahash in the store with a new id and update the count
func (k Keeper) AppendDatahash(
	ctx sdk.Context,
	datahash types.Datahash,
) uint64 {
	// Create the datahash
	count := k.GetDatahashCount(ctx)

	// Set the ID of the appended value
	datahash.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DatahashKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&datahash)
	store.Set(GetDatahashIDBytes(datahash.Id), appendedValue)

	// Update datahash count
	k.SetDatahashCount(ctx, count+1)

	return count
}

// SetDatahash set a specific datahash in the store
func (k Keeper) SetDatahash(ctx sdk.Context, datahash types.Datahash) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DatahashKey))
	b := k.cdc.MustMarshalBinaryBare(&datahash)
	store.Set(GetDatahashIDBytes(datahash.Id), b)
}

// GetDatahash returns a datahash from its id
func (k Keeper) GetDatahash(ctx sdk.Context, id uint64) types.Datahash {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DatahashKey))
	var datahash types.Datahash
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetDatahashIDBytes(id)), &datahash)
	return datahash
}

// HasDatahash checks if the datahash exists in the store
func (k Keeper) HasDatahash(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DatahashKey))
	return store.Has(GetDatahashIDBytes(id))
}

// GetDatahashOwner returns the creator of the datahash
func (k Keeper) GetDatahashOwner(ctx sdk.Context, id uint64) string {
	return k.GetDatahash(ctx, id).Creator
}

// RemoveDatahash removes a datahash from the store
func (k Keeper) RemoveDatahash(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DatahashKey))
	store.Delete(GetDatahashIDBytes(id))
}

// GetAllDatahash returns all datahash
func (k Keeper) GetAllDatahash(ctx sdk.Context) (list []types.Datahash) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DatahashKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Datahash
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetDatahashIDBytes returns the byte representation of the ID
func GetDatahashIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetDatahashIDFromBytes returns ID in uint64 format from a byte array
func GetDatahashIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
