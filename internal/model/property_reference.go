package model

import "strings"

type PropertyReference struct {
	HostName string // Name of the container of this property
	HostId   string // Id of the container of this property
	Property string // Name of the property being referenced
}

func NewPropertyReference(
	host string,
	id string,
	property string,
) PropertyReference {
	return PropertyReference{
		HostName: host,
		HostId:   id,
		Property: property,
	}
}

func ComparePropertyReferences(
	left PropertyReference,
	right PropertyReference,
) int {
	leftHostName := strings.ToLower(left.HostName)
	rightHostName := strings.ToLower(right.HostName)

	result := strings.Compare(leftHostName, rightHostName)
	if result != 0 {
		return result
	}

	leftProperty := strings.ToLower(left.Property)
	rightProperty := strings.ToLower(right.Property)

	return strings.Compare(leftProperty, rightProperty)
}
