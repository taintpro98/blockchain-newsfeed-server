package transport

import (
	"blockchain-newsfeed-server/module/core/dto"
	"blockchain-newsfeed-server/pkg/constants"

	"github.com/gin-gonic/gin"
)

func (t *Transport) ListPosts(ctx *gin.Context) {
	var data dto.ListPostsRequest
	if err := ctx.ShouldBindQuery(&data); err != nil {
		dto.HandleResponse(ctx, nil, err)
		return
	}
	posts, total, err := t.postBiz.ListPosts(ctx, data)
	if err != nil {
		dto.HandleResponse(ctx, nil, err)
	} else {
		limit, offset := data.Paginate.InfoPaginate()
		dto.HandleResponse(ctx, posts, nil, dto.PaginateResponse{
			Total:  total,
			Limit:  limit,
			Offset: offset,
		})
	}
}

func (t *Transport) CreatePost(ctx *gin.Context) {
	var data dto.CreatePostRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		dto.HandleResponse(ctx, nil, err)
		return
	}
	userID := ctx.MustGet(constants.XUserID).(string)
	result, err := t.postBiz.CreatePost(ctx, userID, data)
	dto.HandleResponse(ctx, result, err)
}
