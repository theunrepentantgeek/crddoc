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

	result := &Resource{
		Object: *object,
		Spec:   spec,
		Status: status,
	}

	result.takePropertyOwnership()

	return result, true
}

func (*Resource) Kind() DeclarationType {
	return ResourceDeclaration
}

func (r *Resource) takePropertyOwnership() {
	for _, p := range r.properties {
		p.setContainer(r)
	}
}
