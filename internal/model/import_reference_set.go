package model

type ImportReferenceSet map[string]ImportReference

func NewImportReferenceSet(refs ...ImportReference) ImportReferenceSet {
	result := make(ImportReferenceSet)
	for _, ref := range refs {
		result.Add(ref)
	}

	return result
}

func (set ImportReferenceSet) Add(ref ImportReference) {
	set[ref.Alias] = ref
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
