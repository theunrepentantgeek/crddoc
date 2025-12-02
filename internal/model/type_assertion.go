package model

// TypeAssertionInfo represents a type assertion of the form `var _ Interface = &Type{}`.
// These assertions are used to link interfaces to their implementations.
type TypeAssertionInfo struct {
	InterfaceName string // Name of the interface being implemented
	TypeName      string // Name of the type that implements the interface
	IsPointer     bool   // Whether the implementation is pointer-based (e.g., &Type{})
}

// NewTypeAssertionInfo creates a new TypeAssertionInfo.
func NewTypeAssertionInfo(interfaceName, typeName string, isPointer bool) TypeAssertionInfo {
	return TypeAssertionInfo{
		InterfaceName: interfaceName,
		TypeName:      typeName,
		IsPointer:     isPointer,
	}
}
