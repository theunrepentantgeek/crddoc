package model

type PropertyReference struct {
	Host     string // Name of the container of this property
	Property string // Name of the property being referenced
}

func NewPropertyReference(
	host string,
	property string,
) PropertyReference {
	return PropertyReference{
		Host:     host,
		Property: property,
	}
}

func ComparePropertyReferences(
	left PropertyReference,
	right PropertyReference,
) int {
	if left.Host < right.Host {
		return -1
	}

	if left.Host > right.Host {
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
