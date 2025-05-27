package functions

func (f *Functions) dependencySplitThreshold() int {
	const defaultThreshold = 6

	if f.cfg.ClassDiagrams == nil ||
		f.cfg.ClassDiagrams.DependencySplitThreshold == nil {
		return defaultThreshold
	}

	return *f.cfg.ClassDiagrams.DependencySplitThreshold
}
