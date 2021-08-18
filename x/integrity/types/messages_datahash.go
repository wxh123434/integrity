package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateDatahash{}

func NewMsgCreateDatahash(creator string, detail string, hash string) *MsgCreateDatahash {
	return &MsgCreateDatahash{
		Creator: creator,
		Detail:  detail,
		Hash:    hash,
	}
}

func (msg *MsgCreateDatahash) Route() string {
	return RouterKey
}

func (msg *MsgCreateDatahash) Type() string {
	return "CreateDatahash"
}

func (msg *MsgCreateDatahash) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDatahash) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDatahash) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateDatahash{}

func NewMsgUpdateDatahash(creator string, id uint64, detail string, hash string) *MsgUpdateDatahash {
	return &MsgUpdateDatahash{
		Id:      id,
		Creator: creator,
		Detail:  detail,
		Hash:    hash,
	}
}

func (msg *MsgUpdateDatahash) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDatahash) Type() string {
	return "UpdateDatahash"
}

func (msg *MsgUpdateDatahash) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDatahash) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDatahash) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteDatahash{}

func NewMsgDeleteDatahash(creator string, id uint64) *MsgDeleteDatahash {
	return &MsgDeleteDatahash{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteDatahash) Route() string {
	return RouterKey
}

func (msg *MsgDeleteDatahash) Type() string {
	return "DeleteDatahash"
}

func (msg *MsgDeleteDatahash) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteDatahash) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteDatahash) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
