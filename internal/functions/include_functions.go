package functions

// includeFunctions gives templates access to the configuration
// (exposed as 'includeFunctions()' in templates).
func (f *Functions) includeFunctions() bool {
	if f.cfg == nil {
		return false
	}

	return f.cfg.IncludeFunctions
}
