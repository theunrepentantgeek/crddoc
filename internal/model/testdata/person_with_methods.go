package v1

// PersonWithMethods is a test type with methods for testing function parsing.
type PersonWithMethods struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// GetName returns the name of the person.
func (p PersonWithMethods) GetName() string {
	return p.Name
}

// SetName sets the name of the person.
func (p *PersonWithMethods) SetName(name string) {
	p.Name = name
}

// IsAdult checks if the person is an adult.
func (p PersonWithMethods) IsAdult() bool {
	return p.Age >= 18
}

// UpdateAge updates the age and returns the old age.
func (p *PersonWithMethods) UpdateAge(newAge int) int {
	oldAge := p.Age
	p.Age = newAge
	return oldAge
}

// Compare compares this person with another.
func (p PersonWithMethods) Compare(other PersonWithMethods) (equal bool, ageDiff int) {
	equal = p.Name == other.Name
	ageDiff = p.Age - other.Age
	return
}

// Lookup finds a person attribute by path using variadic parameters.
func (p PersonWithMethods) Lookup(path ...string) (string, bool) {
	if len(path) == 0 {
		return "", false
	}
	if path[0] == "name" {
		return p.Name, true
	}
	return "", false
}
