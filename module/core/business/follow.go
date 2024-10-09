package business

import (
	"blockchain-newsfeed-server/module/core/dto"
	"blockchain-newsfeed-server/module/core/model"
	"blockchain-newsfeed-server/module/core/repository"
	"context"

	"github.com/rs/zerolog/log"
)

type IFollowBiz interface {
	Follow(ctx context.Context, data dto.FollowRequest) error
	UnFollow(ctx context.Context, data dto.FollowRequest) error
}

type followBiz struct {
	followRepo repository.IFollowRepository
}

func NewFollowBiz(
	followRepo repository.IFollowRepository,
) IFollowBiz {
	return &followBiz{
		followRepo: followRepo,
	}
}

func (f *followBiz) Follow(ctx context.Context, data dto.FollowRequest) error {
	log.Info().Ctx(ctx).Interface("data", data).Msg("followBiz Follow")
	return f.followRepo.Insert(ctx, &model.FollowModel{
		FollowerID: data.FollowerID,
		FolloweeID: data.FolloweeID,
	})
}

func (f *followBiz) UnFollow(ctx context.Context, data dto.FollowRequest) error {
	log.Info().Ctx(ctx).Interface("data", data).Msg("followBiz UnFollow")
	return nil
}
