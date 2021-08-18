package keeper

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/SBC/integrity/x/integrity/types"
)

func TestDatahashMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateDatahash(ctx, &types.MsgCreateDatahash{Creator: creator})
		require.NoError(t, err)
		assert.Equal(t, i, int(resp.Id))
	}
}

func TestDatahashMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateDatahash
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateDatahash{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateDatahash{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateDatahash{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateDatahash(ctx, &types.MsgCreateDatahash{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateDatahash(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDatahashMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteDatahash
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteDatahash{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteDatahash{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteDatahash{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateDatahash(ctx, &types.MsgCreateDatahash{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteDatahash(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
