package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "cipher/testutil/keeper"
	"cipher/x/cipher/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CipherKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
