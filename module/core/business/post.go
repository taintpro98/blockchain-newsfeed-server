package business

import (
	"blockchain-newsfeed-server/module/core/dto"
	"blockchain-newsfeed-server/module/core/model"
	"blockchain-newsfeed-server/module/core/repository"
	"context"

	"github.com/rs/zerolog/log"
)

type IPostBiz interface {
	ListPosts(ctx context.Context, data dto.ListPostsRequest) (dto.ListPostsResponse, *int64, error)
	CreatePost(ctx context.Context, userID string, data dto.CreatePostRequest) (dto.CreatePostResponse, error)
}

type postBiz struct {
	postRepo repository.IPostRepository
}

func NewPostBiz(
	postRepo repository.IPostRepository,
) IPostBiz {
	return &postBiz{
		postRepo: postRepo,
	}
}

func (p *postBiz) CreatePost(ctx context.Context, userID string, data dto.CreatePostRequest) (dto.CreatePostResponse, error) {
	log.Info().Ctx(ctx).Interface("data", data).Msg("postBiz CreatePost")
	var result dto.CreatePostResponse
	postInsert := model.PostModel{
		Title:    data.Title,
		Content:  data.Content,
		MediaURL: data.MediaURL,
		UserID:   userID,
	}
	err := p.postRepo.Insert(ctx, &postInsert)
	if err != nil {
		return result, err
	}
	result.PostID = postInsert.ID
	return result, nil
}

func (b *postBiz) ListPosts(ctx context.Context, data dto.ListPostsRequest) (dto.ListPostsResponse, *int64, error) {
	log.Info().Ctx(ctx).Interface("data", data).Msg("postBiz ListPosts")
	postDBs, err := b.postRepo.List(ctx, dto.FilterPost{
		CommonFilter: dto.CommonFilter{
			Select: []string{"id", "title", "content"},
		},
	})
	if err != nil {
		return dto.ListPostsResponse{}, nil, err
	}
	response := dto.ListPostsResponse{
		Posts: postDBs,
	}
	count, err := b.postRepo.Count(ctx, dto.FilterPost{})
	if err != nil {
		return response, nil, err
	}
	return response, count, nil
}
