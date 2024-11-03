package model

type Resource struct {
	Object
	Spec   *Property
	Status *Property
}

func TryNewResource(object *Object) (*Resource, bool) {
	if object == nil {
		return nil, false
	}

	spec, haveSpec := object.properties["spec"]
	if !haveSpec {
		spec, haveSpec = object.properties["Spec"]
	}

	status, haveStatus := object.properties["status"]
	if !haveStatus {
		status, haveStatus = object.properties["Status"]
	}

	if !haveSpec || !haveStatus {
		return nil, false
	}

	return &Resource{
		Object: *object,
		Spec:   spec,
		Status: status,
	}, true
}

func (*Resource) Kind() DeclarationType {
	return ResourceDeclaration
}
