package functions

func (f *Functions) applyEdits(content string) string {
	for _, editor := range f.cfg.Editors {
		content = editor.Replace(content)
	}

	return content
}
