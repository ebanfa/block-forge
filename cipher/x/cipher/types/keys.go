package types

const (
	// ModuleName defines the module name
	ModuleName = "cipher"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cipher"
)

var (
	ParamsKey = []byte("p_cipher")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
