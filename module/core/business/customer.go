package business

import (
	"blockchain-newsfeed-server/module/core/dto"
	"blockchain-newsfeed-server/module/core/model"
	"blockchain-newsfeed-server/module/core/repository"
	"context"
)

type ICustomerBiz interface {
	GetCustomerProfile(ctx context.Context, userID string) (model.UserModel, error)
}

type customerBiz struct {
	userRepo repository.IUserRepository
}

func NewCustomerBiz(
	userRepo repository.IUserRepository,
) ICustomerBiz {
	return &customerBiz{
		userRepo: userRepo,
	}
}

func (v *customerBiz) GetCustomerProfile(ctx context.Context, userID string) (model.UserModel, error) {
	userDB, err := v.userRepo.FindOne(ctx, dto.FilterUser{
		ID: userID,
	})
	if err != nil {
		return model.UserModel{}, err
	}
	return userDB, nil
}
