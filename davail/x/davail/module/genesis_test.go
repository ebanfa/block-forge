package davail_test

import (
	"testing"

	keepertest "davail/testutil/keeper"
	"davail/testutil/nullify"
	davail "davail/x/davail/module"
	"davail/x/davail/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DavailKeeper(t)
	davail.InitGenesis(ctx, k, genesisState)
	got := davail.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
