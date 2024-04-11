package ultima_test

import (
	"testing"

	keepertest "ultima/testutil/keeper"
	"ultima/testutil/nullify"
	ultima "ultima/x/ultima/module"
	"ultima/x/ultima/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.UltimaKeeper(t)
	ultima.InitGenesis(ctx, k, genesisState)
	got := ultima.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
