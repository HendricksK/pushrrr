package templates

import (
    "embed"
    "html/template"
)

type Template struct {
    templates *template.Template
}

func New() *Template {
    funcMap := template.FuncMap{
        "inc": inc,
    }

    templates := template.Must(template.New("").Funcs(funcMap).ParseFS(tmplFS, "views/*.html"))
    return &Template{
        templates: templates,
    }
}


func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}