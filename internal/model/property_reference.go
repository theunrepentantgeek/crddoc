package model

type PropertyReference struct {
	// Reference to the object that contains this property
	Object *Object

	// Name of the property being referenced
	Property string
}

func NewPropertyReference(
	object *Object,
	property string,
) PropertyReference {
	return PropertyReference{
		Object:   object,
		Property: property,
	}
}

func ComparePropertyReferences(
	left PropertyReference,
	right PropertyReference,
) int {
	if left.Object.Name() < right.Object.Name() {
		return -1
	}

	if left.Object.Name() > right.Object.Name() {
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
