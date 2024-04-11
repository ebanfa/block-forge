package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "davail/testutil/keeper"
	"davail/x/davail/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.DavailKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
