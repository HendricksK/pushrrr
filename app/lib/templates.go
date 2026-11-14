package lib

import (

	"github.com/labstack/echo/v5"
	"io"
	"embed"
)

//go:embed templates/*.html
var tmplFS embed.FS

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func New() *Template {
    funcMap := template.FuncMap{
        // "inc": inc,
    }

    templates := template.Must(template.New("").Funcs(funcMap).ParseFS(tmplFS, "*.html"))
    return &Template{
        templates: templates,
    }
}