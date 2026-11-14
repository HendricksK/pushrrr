package controllers

import (

	"html/template"
    "log"           
    "net/http"
	"fmt"

	models "github.com/HendricksK/pushrrr/app/models"
)

func GetPost(id int, w http.ResponseWriter, r *http.Request) {
	posts := models.GetPost(id)

	fmt.Printf("%+v\n", posts)

	w.Header().Add("Server", "Go")

    // Initialize a slice containing the paths to the two files. It's important
    // to note that the file containing our base template must be the *first*
    // file in the slice.
    files := []string{
        "./templates/shared/base.html",
        "./templates/pages/home.html",
    }

    // Use the template.ParseFiles() function to read the files and store the
    // templates in a template set. Notice that we use ... to pass the contents 
    // of the files slice as variadic arguments.
    ts, err := template.ParseFiles(files...)
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Use the ExecuteTemplate() method to write the content of the "base" 
    // template as the response body.
    err = ts.ExecuteTemplate(w, "base", nil)
    if err != nil {
        log.Print(err.Error())
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

// func GetPosts(ids ...[]uint) []models.Post {

// 	// if ids != nil {
// 	// 	return models.GetPostsWhereIn(ids)
// 	// }
// 	return models.GetPosts()
// }

// func CreatePost(posts ...string) []models.Post {
// 	return models.CreatePost()
// }
