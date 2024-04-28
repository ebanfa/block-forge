package operations

/* i mport (
	"github.com/edward1christian/block-forge/nova/pkg/components/operations/commands"
	"github.com/edward1christian/block-forge/pkg/application/component"
)

// Define a struct to represent an operation to register
type OperationRegistration struct {
	ID        string
	FactoryID string
	Factory   component.ComponentFactoryInterface
}

// GetOperationsToRegister returns the list of operations to register
func GetOperationsToRegister() []OperationRegistration {
	// Create a map to hold the operation names and their corresponding factories
	operationFactories := map[string]component.ComponentFactoryInterface{
		"BuildProjectOp":        &commands.BuildProjectOpFactory{},
		"GenerateArtifactsOp":   &commands.GenerateArtifactsOpFactory{},
		"InitProjectOp":         &commands.InitProjectOpFactory{},
		"LoadConfigurationOp":   &commands.LoadConfigurationOpFactory{},
		"AddMessageOp":          &commands.AddMessageOpFactory{},
		"AddModuleOp":           &commands.AddModuleOpFactory{},
		"AddQueryOp":            &commands.AddQueryOpFactory{},
		"RemoveProjectConfigOp": &commands.RemoveProjectConfigOpFactory{},
		"RunProjectOp":          &commands.RunProjectOpFactory{},
		"AddTypeOp":             &commands.AddTypeOpFactory{},
		"ValidateConfigOp":      &commands.ValidateConfigOpFactory{},
	}

	// Create a slice to hold the operation registrations
	var operationsToRegister []OperationRegistration

	// Iterate over the map to populate operationsToRegister
	for operationName, factory := range operationFactories {
		// Get the factory ID
		factoryID := operationName + "Factory"

		// Add the operation registration to operationsToRegister
		operationsToRegister = append(operationsToRegister, OperationRegistration{
			ID:        operationName,
			FactoryID: factoryID,
			Factory:   factory,
		})
	}

	return operationsToRegister
}
*/
