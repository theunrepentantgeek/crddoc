package v1

// Additional methods for PersonWithMethods defined in a separate file.

// GetFullInfo returns a formatted string with all person information.
func (p PersonWithMethods) GetFullInfo() string {
	return p.Name
}

// IncrementAge increments the age by one.
func (p *PersonWithMethods) IncrementAge() {
	p.Age++
}
