// package helpers

// import (
//     "embed"
//     "html/template"
//     "io"

//     "github.com/labstack/echo/v5"
// )

// //go:embed templates/*.html
// var tmplFS embed.FS

// type Template struct {
//     templates *template.Template
// }

// func New() *Template {
//     funcMap := template.FuncMap{
//         "inc": inc,
//     }

//     templates := template.Must(template.New("").Funcs(funcMap).ParseFS(tmplFS, "app/views/*/*.html"))
//     return &Template{
//         templates: templates,
//     }
// }


// func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
//     return t.templates.ExecuteTemplate(w, name, data)
// }