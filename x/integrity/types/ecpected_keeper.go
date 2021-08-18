package types // noalias

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BankKeeper defines the contract needed to be fulfilled for banking and supply
// dependencies.
type BankKeeper interface {
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
} //mint 的keeper实现了以上接口

type MintKeeper interface {
	MintCoins(ctx sdk.Context, newCoins sdk.Coins) error
}

//调用这个实现币的生成
