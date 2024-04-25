package tree

// ModuleConfig represents the configuration for a module node.
type ModuleConfig struct {
	Dependencies []string
	// Other module-specific configurations
}

// QueryConfig represents the configuration for a query or entity node.
type QueryConfig struct {
	Inputs   []FieldConfig
	Response FieldConfig
	// Other query-specific configurations
}

// TransactionConfig represents the configuration for a transaction node.
type TransactionConfig struct {
	Inputs  []FieldConfig
	Outputs []FieldConfig
	Handler string
	// Other transaction-specific configurations
}

// FieldConfig represents the configuration for a field node.
type FieldConfig struct {
	Name string
	Type string
	// Other field-specific configurations
}
