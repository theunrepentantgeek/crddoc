package functions

func (f *Functions) applyEdits(content string) string {
	return f.editors.Replace(content)
}
