package model

// Parameter represents a function parameter or return value.
type Parameter struct {
	Name       string        // Name of the parameter (may be empty)
	Type       TypeReference // Type of the parameter
	IsVariadic bool          // Whether the parameter is variadic (e.g., ...string)
}
