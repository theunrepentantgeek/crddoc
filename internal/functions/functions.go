package functions

import "text/template"

func CreateFuncMap() template.FuncMap {
	return template.FuncMap{
		"renderType": renderType,
		"unwrap":     unwrap,
	}
}
