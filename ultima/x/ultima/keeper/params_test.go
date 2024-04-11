package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "ultima/testutil/keeper"
	"ultima/x/ultima/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.UltimaKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
