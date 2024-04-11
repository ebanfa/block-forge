package types

const (
	// ModuleName defines the module name
	ModuleName = "ultima"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ultima"
)

var (
	ParamsKey = []byte("p_ultima")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
