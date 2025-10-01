package models

import (
	"fmt"
	"runtime"
	"time"

	database "github.com/HendricksK/pushrrr/app/database"
	helpers "github.com/HendricksK/pushrrr/app/helpers"
)

type Post struct {
	Id        uint       `gorm:"primaryKey"`
	Name      *string    `json:"name"`
	Data      *string    `json:"data"`
	Uri       *string    `json:"uri"`
	Author    *string    `json:"author"`
	Tags      *string    `json:"tags"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

var __table = "post"

func Ping() string {
	return "pong"
}

func GetPost(id string) Post {

	var post Post

	db := database.Open()

	err := db.QueryRow("SELECT * FROM "+__table+" WHERE id = ?", id).Scan(
		&post.Id,
	)

	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		helpers.Log(err.Error(), filename, line)
	}

	database.Close(db)

	return post
}

func GetPosts() []Post {

	var posts []Post

	db := database.Open()

	rows, err := db.Query("SELECT * FROM " + __table)
	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		helpers.Log(err.Error(), filename, line)
		return posts
	}
	defer rows.Close()

	for rows.Next() {
		var post Post

		err = rows.Scan(
			&post.Id)

		if err != nil {
			_, filename, line, _ := runtime.Caller(1)
			helpers.Log(err.Error(), filename, line)
			panic(err)
		}

		posts = append(posts, post)
	}

	database.Close(db)

	return posts
}

func GetPostsWhereIn(ids []string) []Post {

	fmt.Print(helpers.GetEnvVar("DATABASE"))

	var posts []Post

	db := database.Open()

	rows, err := db.Query("SELECT * FROM "+__table+"WHERE id IN ?", ids)
	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		helpers.Log(err.Error(), filename, line)
		return posts
	}
	defer rows.Close()

	for rows.Next() {
		var post Post

		err = rows.Scan(
			&post.Id)

		if err != nil {
			_, filename, line, _ := runtime.Caller(1)
			helpers.Log(err.Error(), filename, line)
			panic(err)
		}

		posts = append(posts, post)
	}

	database.Close(db)

	return posts
}
