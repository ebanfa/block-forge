package cipher_test

import (
	"testing"

	keepertest "cipher/testutil/keeper"
	"cipher/testutil/nullify"
	cipher "cipher/x/cipher/module"
	"cipher/x/cipher/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CipherKeeper(t)
	cipher.InitGenesis(ctx, k, genesisState)
	got := cipher.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
