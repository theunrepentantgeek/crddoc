package functions

// configuration gives templates access to the configuration
// (exposed as 'config()' in templates).
func (f *Functions) includeClassDiagrams() bool {
	// If no class diagram config, return false
	if f.cfg.ClassDiagrams == nil {
		return false
	}

	// If class diagram specifies 'enabled', return that
	if f.cfg.ClassDiagrams.Enabled != nil {
		return *f.cfg.ClassDiagrams.Enabled
	}

	// Otherwise default based on whether other class diagram options are set
	return f.cfg.ClassDiagrams.Empty()
}
