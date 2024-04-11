package types

const (
	// ModuleName defines the module name
	ModuleName = "davail"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_davail"
)

var (
	ParamsKey = []byte("p_davail")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
