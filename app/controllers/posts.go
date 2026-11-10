package controllers

import (
	models "github.com/HendricksK/pushrrr/app/models"
)

func GetPost(id uint) models.Post {
	return models.GetPost(id)
}

func GetPosts(ids ...[]uint) []models.Post {

	// if ids != nil {
	// 	return models.GetPostsWhereIn(ids)
	// }
	return models.GetPosts()
}

func CreatePost(posts ...string) []models.Post {
	return models.CreatePost()
}
