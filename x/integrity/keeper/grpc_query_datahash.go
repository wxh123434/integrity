package keeper

import (
	"context"

	"github.com/SBC/integrity/x/integrity/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DatahashAll(c context.Context, req *types.QueryAllDatahashRequest) (*types.QueryAllDatahashResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var datahashs []*types.Datahash
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	datahashStore := prefix.NewStore(store, types.KeyPrefix(types.DatahashKey))

	pageRes, err := query.Paginate(datahashStore, req.Pagination, func(key []byte, value []byte) error {
		var datahash types.Datahash
		if err := k.cdc.UnmarshalBinaryBare(value, &datahash); err != nil {
			return err
		}

		datahashs = append(datahashs, &datahash)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDatahashResponse{Datahash: datahashs, Pagination: pageRes}, nil
}

func (k Keeper) Datahash(c context.Context, req *types.QueryGetDatahashRequest) (*types.QueryGetDatahashResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var datahash types.Datahash
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasDatahash(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DatahashKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetDatahashIDBytes(req.Id)), &datahash)

	return &types.QueryGetDatahashResponse{Datahash: &datahash}, nil
}
