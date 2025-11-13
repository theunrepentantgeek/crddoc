package model

// Parameter represents a function parameter or return value.
type Parameter struct {
	Name string        // Name of the parameter (may be empty)
	Type TypeReference // Type of the parameter
}

// NewNamedParameter creates a new named Parameter.
// name is the parameter name.
// typeRef is the type of the parameter.
func NewNamedParameter(
	name string,
	typeRef TypeReference,
) Parameter {
	return Parameter{
		Name: name,
		Type: typeRef,
	}
}

// NewAnonymousParameter creates a new anonymous Parameter.
// typeRef is the type of the parameter.
func NewAnonymousParameter(
	typeRef TypeReference,
) Parameter {
	return Parameter{
		Name: "",
		Type: typeRef,
	}
}
