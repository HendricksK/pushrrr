package models

import (
	"runtime"

	helpers "github.com/HendricksK/pushrrr/app/helpers"
)

type Post struct {
	id	int
}

func GetPost(postId int) Post {
	var post Post


	// eventually use helpers to get all aggregate data if needed, for now just use sql golang native lib

	db := helpers.OpenSQL()

	err := db.QueryRow("SELECT * FROM post WHERE id = 2").Scan(&post.id)

	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		helpers.Log(err.Error(), filename, line)
	}

	helpers.CloseSQL(db)

	return post
}