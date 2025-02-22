package model

type ImportReferenceSet map[string]ImportReference

func (set ImportReferenceSet) Add(ref ImportReference) {
	set[ref.ImportPath] = ref
}

func (set ImportReferenceSet) LookupImportPath(
	typ TypeReference,
) (string, bool) {
	if typ.pkg == "" {
		return "", false
	}

	if ref, ok := set[typ.pkg]; ok {
		return ref.ImportPath, true
	}

	return "", false
}
