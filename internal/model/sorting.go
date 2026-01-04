package model

import "strings"

// alphabeticalPropertyComparison does a case insensitive comparison of the names of the
// two properties, allowing them to be sorted.
func alphabeticalPropertyComparison(left *Property, right *Property) int {
	leftName := strings.ToLower(left.Name)
	rightName := strings.ToLower(right.Name)

	return strings.Compare(leftName, rightName)
}

// alphabeticalFunctionComparison does a case insensitive comparison of the names of the
// two functions, allowing them to be sorted.
func alphabeticalFunctionComparison(left *Function, right *Function) int {
	leftName := strings.ToLower(left.Name)
	rightName := strings.ToLower(right.Name)

	return strings.Compare(leftName, rightName)
}

// alphabeticalInterfaceComparison does a case insensitive comparison of the names of the
// two interfaces, allowing them to be sorted.
func alphabeticalInterfaceComparison(left *Interface, right *Interface) int {
	leftName := strings.ToLower(left.Name())
	rightName := strings.ToLower(right.Name())

	return strings.Compare(leftName, rightName)
}

// alphabeticalObjectComparison does a case insensitive comparison of the names of the
// two objects, allowing them to be sorted.
func alphabeticalObjectComparison(left *Object, right *Object) int {
	leftName := strings.ToLower(left.Name())
	rightName := strings.ToLower(right.Name())

	return strings.Compare(leftName, rightName)
}

// alphabeticalDeclarationComparison does a case insensitive comparison of the names of the
// two declarations, allowing them to be sorted.
func alphabeticalDeclarationComparison(left Declaration, right Declaration) int {
	leftName := strings.ToLower(left.Name())
	rightName := strings.ToLower(right.Name())

	return strings.Compare(leftName, rightName)
}
