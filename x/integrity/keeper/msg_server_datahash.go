package keeper

import (
	"context"
	"fmt"

	"github.com/SBC/integrity/x/integrity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

//FIXME 在此处加入mintcoin
func (k msgServer) CreateDatahash(goCtx context.Context, msg *types.MsgCreateDatahash) (*types.MsgCreateDatahashResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var datahash = types.Datahash{
		Creator: msg.Creator,
		Detail:  msg.Detail,
		Hash:    msg.Hash,
	}

	id := k.AppendDatahash(
		ctx,
		datahash,
	)

	txgas := sdk.NewInt(1)
	coin := sdk.NewCoin("stake", txgas)
	coins := sdk.NewCoins(coin)

	err := k.MintCoinsForHash(ctx, coins)
	if err != nil {
		fmt.Print("1 err")
		panic("mintcoins err")
	}

	creatorAddr, err2 := sdk.AccAddressFromBech32(datahash.Creator)
	if err2 != nil {
		panic("address err")
	}
	err3 := k.SendCoinsFromMintModuleToAccount(ctx, coins, creatorAddr)
	if err3 != nil {
		panic("sendta err")
	}

	return &types.MsgCreateDatahashResponse{
		Id: id,
	}, nil

}

func (k msgServer) UpdateDatahash(goCtx context.Context, msg *types.MsgUpdateDatahash) (*types.MsgUpdateDatahashResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var datahash = types.Datahash{
		Creator: msg.Creator,
		Id:      msg.Id,
		Detail:  msg.Detail,
		Hash:    msg.Hash,
	}

	// Checks that the element exists
	if !k.HasDatahash(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetDatahashOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetDatahash(ctx, datahash)

	return &types.MsgUpdateDatahashResponse{}, nil

}

func (k msgServer) DeleteDatahash(goCtx context.Context, msg *types.MsgDeleteDatahash) (*types.MsgDeleteDatahashResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasDatahash(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetDatahashOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveDatahash(ctx, msg.Id)

	return &types.MsgDeleteDatahashResponse{}, nil
}
