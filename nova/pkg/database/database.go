package database

// Database provides methods for interacting with a persistent storage
// mechanism, such as an IAVL tree.
type Database interface {
	// Get retrieves the value associated with the given key from the tree.
	Get(key []byte) ([]byte, error)

	// Set stores the key-value pair in the tree. If the key already exists,
	// its value will be updated.
	Set(key, value []byte) error

	// Delete removes the key-value pair from the tree.
	Delete(key []byte) error

	// Has returns true if the key exists in the tree, otherwise false.
	Has(key []byte) (bool, error)

	// Iterate iterates over all keys of the tree and calls the given function
	// for each key-value pair. Iteration stops if the function returns true.
	Iterate(fn func(key, value []byte) bool) error

	// IterateRange iterates over all key-value pairs with keys in the range
	// [start, end) and calls the given function for each pair. Iteration stops
	// if the function returns true.
	IterateRange(start, end []byte, ascending bool, fn func(key, value []byte) bool) error

	// Hash returns the root hash of the tree.
	Hash() []byte

	// Version returns the version of the tree.
	Version() int64

	// Load loads the latest versioned tree from disk.
	Load() (int64, error)

	// LoadVersion loads a specific version of the tree from disk.
	LoadVersion(targetVersion int64) (int64, error)

	// SaveVersion saves a new tree version to disk.
	SaveVersion() ([]byte, int64, error)

	// Rollback resets the working tree to the latest saved version, discarding
	// any unsaved modifications.
	Rollback()

	// Close closes the tree.
	Close() error

	// String returns a string representation of the tree.
	String() (string, error)

	// WorkingVersion returns the current working version of the tree.
	WorkingVersion() int64

	// WorkingHash returns the root hash of the current working tree.
	WorkingHash() []byte

	// AvailableVersions returns a list of available versions.
	AvailableVersions() []int

	// IsEmpty checks if the database is empty.
	IsEmpty() bool
}
