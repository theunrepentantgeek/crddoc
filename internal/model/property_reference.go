package model

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
	if left.HostName < right.HostName {
		return -1
	}

	if left.HostName > right.HostName {
		return 1
	}

	if left.Property < right.Property {
		return -1
	}

	if left.Property > right.Property {
		return 1
	}

	return 0
}
