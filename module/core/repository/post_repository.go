package repository

import (
	"blockchain-newsfeed-server/config"
	"blockchain-newsfeed-server/module/core/dto"
	"blockchain-newsfeed-server/module/core/model"
	"context"

	"gorm.io/gorm"
)

type IPostRepository interface {
	List(ctx context.Context, filter dto.FilterPost) ([]model.PostModel, error)
	Count(ctx context.Context, filter dto.FilterPost) (*int64, error)
	Insert(ctx context.Context, data *model.PostModel) error
}

type postRepository struct {
	commonRepository
}

func NewPostRepository(cfg *config.Config, db *gorm.DB) IPostRepository {
	return &postRepository{
		commonRepository: commonRepository{
			db:       db,
			configDb: cfg.Database,
		},
	}
}

func (u *postRepository) tableName() string {
	return model.PostModel{}.TableName()
}

func (s *postRepository) BuildQuery(ctx context.Context, filter dto.FilterPost) *gorm.DB {
	query := s.table(ctx, s.tableName())
	return query
}

func (u *postRepository) List(ctx context.Context, filter dto.FilterPost) ([]model.PostModel, error) {
	var result []model.PostModel // khoi tao cho nay ra mang rong
	err := u.CList(ctx, CommonRepositoryParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(ctx, filter),
		Data:         &result,
	})
	return result, err
}

func (u *postRepository) Count(ctx context.Context, filter dto.FilterPost) (*int64, error) {
	return u.CCount(ctx, CommonRepositoryParams{
		TableName:    u.tableName(),
		Filter:       filter,
		CommonFilter: filter.CommonFilter,
		Query:        u.BuildQuery(ctx, filter),
	})
}

func (u *postRepository) Insert(ctx context.Context, data *model.PostModel) error {
	return u.CInsert(ctx, CommonRepositoryParams{
		TableName: u.tableName(),
		Data:      data,
	})
}
