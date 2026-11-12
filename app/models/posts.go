package models

type Post struct {
	id	int
}

func GetPost(postId int) Post {
	// var post Post

	// post.id = 12

	post := Post{
		id: 12,
	}

	return post
}