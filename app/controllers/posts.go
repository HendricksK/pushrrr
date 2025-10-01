package controllers

import (
	models "github.com/HendricksK/pushrrr/app/models"
)

func GetPost(id string) models.Post {
	return models.GetPost(id)
}

func GetPosts(ids ...string) []models.Post {

	if ids != nil {
		return models.GetPostsWhereIn(ids)
	}
	return models.GetPosts()
}
