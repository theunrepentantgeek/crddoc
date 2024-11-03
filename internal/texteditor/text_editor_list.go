package texteditor

import "github.com/theunrepentantgeek/crddoc/internal/config"

type List struct {
	editors []*Editor
}

func New(cfg *config.Config) (*List, error) {
	editors := make([]*Editor, 0, len(cfg.Editors))

	for _, f := range cfg.Editors {
		editor, err := NewEditor(f)
		if err != nil {
			return nil, err
		}

		editors = append(editors, editor)
	}

	return &List{
		editors: editors,
	}, nil
}

func (list *List) Replace(input string) string {
	result := input
	for _, editor := range list.editors {
		result = editor.Replace(result)
	}

	return result
}
