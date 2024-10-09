package repository

import (
	"blockchain-newsfeed-server/module/core/model"
	"context"
)

type IFollowRepository interface {
	Insert(ctx context.Context, data *model.FollowModel) error
}

type followRepository struct {
	commonRepository
}

func NewFollowRepository() IFollowRepository {
	return &followRepository{}
}

func (u *followRepository) tableName() string {
	return model.FollowModel{}.TableName()
}

func (u *followRepository) Insert(ctx context.Context, data *model.FollowModel) error {
	return u.CInsert(ctx, CommonRepositoryParams{
		TableName: u.tableName(),
		Data:      data,
	})
}