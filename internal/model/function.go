package model

import (
	"github.com/dave/dst"
)

// Function represents a method declared on a struct type.
type Function struct {
	Name              string        // Name of the function/method
	Receiver          TypeReference // Type of the receiver (pointer or value)
	IsPointerReceiver bool          // Whether the receiver is a pointer
	Parameters        []Parameter   // Function parameters
	Results           []Parameter   // Return values
	description       []string
	declaredOn        *Object // The object this function is declared on
}

type FunctionList []*Function

// TryNewFunction attempts to create a Function from a dst.FuncDecl.
// Returns the function and true if successful, nil and false otherwise.
func TryNewFunction(
	decl *dst.FuncDecl,
) (*Function, bool) {
	// Only process functions with receivers (methods)
	if decl.Recv == nil || len(decl.Recv.List) == 0 {
		return nil, false
	}

	// Get the receiver type
	recv := decl.Recv.List[0]
	receiverType := NewTypeReferenceFromExpr(recv.Type)

	// Check if the receiver is a pointer
	isPointer := false
	if _, ok := recv.Type.(*dst.StarExpr); ok {
		isPointer = true
	}

	// Parse comments
	description, _ := ParseComments(decl.Decs.Start.All())
	funcName := decl.Name.Name
	description = formatComments(description, funcName)

	// Parse parameters
	var params []Parameter
	if decl.Type.Params != nil {
		params = parseFieldList(decl.Type.Params)
	}

	// Parse results
	var results []Parameter
	if decl.Type.Results != nil {
		results = parseFieldList(decl.Type.Results)
	}

	return &Function{
		Name:              funcName,
		Receiver:          receiverType,
		IsPointerReceiver: isPointer,
		Parameters:        params,
		Results:           results,
		description:       description,
	}, true
}

// parseFieldList converts a dst.FieldList into a slice of Parameters.
func parseFieldList(fields *dst.FieldList) []Parameter {
	var params []Parameter

	for _, field := range fields.List {
		typeRef := NewTypeReferenceFromExpr(field.Type)

		// A field can have multiple names (e.g., "x, y int")
		if len(field.Names) == 0 {
			// Unnamed parameter
			params = append(params, Parameter{
				Name: "",
				Type: typeRef,
			})
		} else {
			// Named parameters
			for _, name := range field.Names {
				params = append(params, Parameter{
					Name: name.Name,
					Type: typeRef,
				})
			}
		}
	}

	return params
}

// Description returns the description of the function.
func (f *Function) Description() []string {
	return f.description
}

// DeclaredOn returns the object this function is declared on.
func (f *Function) DeclaredOn() *Object {
	return f.declaredOn
}

// setDeclaredOn sets the object this function is declared on.
func (f *Function) setDeclaredOn(obj *Object) {
	f.declaredOn = obj
}
