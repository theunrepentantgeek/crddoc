package functions

// configuration gives templates access to the configuration
// (exposed as 'config()' in templates).
func (f *Functions) includeClassDiagrams() bool {
	return f.cfg.ClassDiagrams
}
