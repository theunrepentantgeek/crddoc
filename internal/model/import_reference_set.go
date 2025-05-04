package model

// ImportReferenceSet is a set of ImportReference objects, indexed by their alias.
type ImportReferenceSet map[string]ImportReference

// NewImportReferenceSet creates a new ImportReferenceSet and adds the provided
// ImportReferences to it.
func NewImportReferenceSet(refs ...ImportReference) ImportReferenceSet {
	result := make(ImportReferenceSet)
	for _, ref := range refs {
		result.Add(ref)
	}

	return result
}

// Add adds an ImportReference to the set, using its alias as the key.
// If the alias already exists, it will be overwritten.
func (set ImportReferenceSet) Add(ref ImportReference) {
	set[ref.Alias] = ref
}

// LookupImportPath looks up the import path for a given TypeReference in the
// ImportReferenceSet. It returns the import path and true if found, false otherwise.
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
