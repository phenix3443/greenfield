package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateStorageProvider = "create_storage_provider"
	TypeMsgEditStorageProvider   = "edit_storage_provider"
	TypeMsgDeposit               = "deposit"
)

var (
	_ sdk.Msg = &MsgCreateStorageProvider{}
	_ sdk.Msg = &MsgEditStorageProvider{}
	_ sdk.Msg = &MsgDeposit{}
)

// NewMsgCreateStorageProvider creates a new MsgCreateStorageProvider instance.
// creator is the module account of gov module
// SpAddress is the account address of storage provider
// fundAddress is another accoutn address of storage provider which used to deposit or rewarding
func NewMsgCreateStorageProvider(creator string, SpAddress sdk.AccAddress, fundingAddress sdk.AccAddress, description Description, deposit sdk.Coin) (*MsgCreateStorageProvider, error) {
	return &MsgCreateStorageProvider{
		Creator:        creator,
		SpAddress:      SpAddress.String(),
		FundingAddress: fundingAddress.String(),
		Description:    description,
		Deposit:        deposit,
	}, nil
}

// Route implements the sdk.Msg interface.
func (msg *MsgCreateStorageProvider) Route() string {
	return RouterKey
}

// Type implements the sdk.Msg interface.
func (msg *MsgCreateStorageProvider) Type() string {
	return TypeMsgCreateStorageProvider
}

// GetSigners implements the sdk.Msg interface.
func (msg *MsgCreateStorageProvider) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromHexUnsafe(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes returns the message bytes to sign over.
func (msg *MsgCreateStorageProvider) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg *MsgCreateStorageProvider) ValidateBasic() error {
	if _, err := sdk.AccAddressFromHexUnsafe(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if _, err := sdk.AccAddressFromHexUnsafe(msg.SpAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid operator address (%s)", err)
	}
	if _, err := sdk.AccAddressFromHexUnsafe(msg.FundingAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid fund address (%s)", err)
	}

	if !msg.Deposit.IsValid() || !msg.Deposit.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid deposit amount")
	}

	if msg.Description == (Description{}) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty description")
	}
	return nil
}

// NewMsgEditStorageProvider creates a new MsgEditStorageProvider instance
// TODO(fynn): add morer modifiable items if needed.
func NewMsgEditStorageProvider(creator string, spAddress string, description Description) *MsgEditStorageProvider {
	return &MsgEditStorageProvider{
		SpAddress:   spAddress,
		Description: description,
	}
}

// Route implements the sdk.Msg interface.
func (msg *MsgEditStorageProvider) Route() string {
	return RouterKey
}

// Type implements the sdk.Msg interface.
func (msg *MsgEditStorageProvider) Type() string {
	return TypeMsgEditStorageProvider
}

// GetSigners implements the sdk.Msg interface.
func (msg *MsgEditStorageProvider) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.SpAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes returns the message bytes to sign over.
func (msg *MsgEditStorageProvider) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg *MsgEditStorageProvider) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.SpAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Description == (Description{}) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empty description")
	}
	return nil
}

// NewMsgDeposit creates a new MsgDeposit instance
func NewMsgDeposit(creator string, spAddress string, deposit sdk.Coin) *MsgDeposit {
	return &MsgDeposit{
		Creator:   creator,
		SpAddress: spAddress,
		Deposit:   deposit,
	}
}

// Route implements the sdk.Msg interface.
func (msg *MsgDeposit) Route() string {
	return RouterKey
}

// Type implements the sdk.Msg interface.
func (msg *MsgDeposit) Type() string {
	return TypeMsgDeposit
}

// GetSigners implements the sdk.Msg interface.
func (msg *MsgDeposit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes returns the message bytes to sign over.
func (msg *MsgDeposit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg *MsgDeposit) ValidateBasic() error {
	if _, err := sdk.AccAddressFromHexUnsafe(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if !msg.Deposit.IsValid() || !msg.Deposit.Amount.IsPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid deposit amount")
	}

	return nil
}
