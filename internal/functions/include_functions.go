package functions

// includeFunctions gives templates access to the configuration
// (exposed as 'includeFunctions()' in templates).
func (f *Functions) includeFunctions() bool {
	return f.cfg.IncludeFunctions
}
