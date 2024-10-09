package dto

import "blockchain-newsfeed-server/module/core/model"

type CreatePostRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	MediaURL string `json:"media_url"`
}

type CreatePostResponse struct {
	PostID string `json:"post_id"`
}

type FilterPost struct {
	CommonFilter CommonFilter
}

type ListPostsRequest struct {
	Paginate
}

type ListPostsResponse struct {
	Posts []model.PostModel `json:"posts"`
}
