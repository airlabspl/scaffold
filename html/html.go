package html

import "html/template"

func New(paths ...string) *template.Template {
	return template.Must(template.ParseFS(FS, paths...))
}
