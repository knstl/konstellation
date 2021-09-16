package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/konstellation/konstellation/x/issue/types"
)

func TestCoinIssueDenomQuerySingle(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCoinIssueDenom(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetCoinIssueDenomRequest
		response *types.QueryGetCoinIssueDenomResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetCoinIssueDenomRequest{Id: msgs[0].Id},
			response: &types.QueryGetCoinIssueDenomResponse{CoinIssueDenom: &msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetCoinIssueDenomRequest{Id: msgs[1].Id},
			response: &types.QueryGetCoinIssueDenomResponse{CoinIssueDenom: &msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetCoinIssueDenomRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.CoinIssueDenom(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestCoinIssueDenomQueryPaginated(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCoinIssueDenom(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllCoinIssueDenomRequest {
		return &types.QueryAllCoinIssueDenomRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.CoinIssueDenomAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			for j := i; j < len(msgs) && j < i+step; j++ {
				assert.Equal(t, &msgs[j], resp.CoinIssueDenom[j-i])
			}
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.CoinIssueDenomAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			for j := i; j < len(msgs) && j < i+step; j++ {
				assert.Equal(t, &msgs[j], resp.CoinIssueDenom[j-i])
			}
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.CoinIssueDenomAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.CoinIssueDenomAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
