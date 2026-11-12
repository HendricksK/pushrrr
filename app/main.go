package main

import (
	"fmt"
	"html/template"
	"io"	
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"embed"


	"github.com/labstack/echo/v5"
	// templates "github.com/HendricksK/pushrrr/app/helpers"
)

type Post struct {
	Post   string
	Data   []string
	IsRead bool
	Sent   string
	Read   string
}

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


func main() {

	// t := templates.New()

    // e := echo.New()
    // e.Renderer = t

	http.HandleFunc("/healthz", healthz)
	// http.HandleFunc("/assets/*", fail)
	http.HandleFunc("/", index)
	http.HandleFunc("/test", test)
	http.HandleFunc("/upload", fileUploadHandler)
	http.HandleFunc("/push", push)

	http.HandleFunc("/assets/theme/images/", themeImages)

	fmt.Println("Server running on :9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func fileUploadHandler(w http.ResponseWriter, r *http.Request) {
	// Limit file size to 10MB. This line saves you from those accidental 100MB uploads!
	r.ParseMultipartForm(10 << 20)

	// Retrieve the file from form data
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Fprintf(w, "Uploaded File: %s\n", handler.Filename)
	fmt.Fprintf(w, "File Size: %d\n", handler.Size)
	fmt.Fprintf(w, "MIME Header: %v\n", handler.Header)

	// Now let’s save it locally
	dst, err := createFile(handler.Filename)
	if err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the destination file
	if _, err := dst.ReadFrom(file); err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
	}
}

func createFile(filename string) (*os.File, error) {
	// Create an uploads directory if it doesn’t exist
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", 0755)
	}

	// Build the file path and create it
	dst, err := os.Create(filepath.Join("uploads", filename))
	if err != nil {
		return nil, err
	}

	return dst, nil
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pushrrr - App version: %s", "0.0.1")
}

func index(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	posts := []Post{
		{
			Post: "here lies a post",
			Data: []string{
				"wqewe",
				"qweqwe",
			},
			IsRead: true,
			Sent:   t.String(),
			Read:   t.String(),
		},
	}

	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	// page1Templ := template.Must(template.New("p1").Parse("./views/index.html"))
	// tmpl.Must(page1Templ.Parse("./views/shared/head.html"))

	// return c.Render(http.StatusOK, "index.html", posts)

	tmpl.Execute(w, posts)
}

func test(w http.ResponseWriter, r *http.Request) {

	t := time.Now()

	posts := []Post{
		{
			Post: "here lies a post",
			Data: []string{
				"wqewe",
				"qweqwe",
			},
			IsRead: true,
			Sent:   t.Format("2006-01-02 15:04:05"),
			Read:   t.Format("2006-01-02 15:04:05"),
		},
	}

	tmpl := template.Must(template.ParseFiles("./views/index.html"))

	tmpl.Execute(w, posts)
}

func themeImages(w http.ResponseWriter, r *http.Request) {

	pwd, _ := os.Getwd()

	fmt.Printf("%s\n", pwd+r.URL.RequestURI())

	buf, err := os.ReadFile(pwd + r.URL.RequestURI())

	if err != nil {

		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(buf)
}

func fail(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(401)
}

func push(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Pushrrr - App version: %s ", r.Method)

	data, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Printf("server: could not read request body: %s\n", err)
	}

	fmt.Printf("server: request body: %s\n", data)

	// posts := controllers.CreatePost(data['data'])

	// fmt.Fprintf(w, "%d ", posts[0].Id)
}
