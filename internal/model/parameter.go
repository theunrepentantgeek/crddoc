package model

// Parameter represents a function parameter or return value.
type Parameter struct {
	Name       string        // Name of the parameter (may be empty)
	Type       TypeReference // Type of the parameter
	IsVariadic bool          // Whether the parameter is variadic (e.g., ...string)
}

// NewNamedParameter creates a new named Parameter.
// name is the parameter name.
// typeRef is the type of the parameter.
// isVariadic indicates if the parameter is variadic.
func NewNamedParameter(
	name string,
	typeRef TypeReference,
	isVariadic bool,
) Parameter {
	return Parameter{
		Name:       name,
		Type:       typeRef,
		IsVariadic: isVariadic,
	}
}

// NewAnonymousParameter creates a new anonymous Parameter.
// typeRef is the type of the parameter.
// isVariadic indicates if the parameter is variadic.
func NewAnonymousParameter(
	typeRef TypeReference,
	isVariadic bool,
) Parameter {
	return Parameter{
		Name:       "",
		Type:       typeRef,
		IsVariadic: isVariadic,
	}
}
