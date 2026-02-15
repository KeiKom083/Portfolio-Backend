package usecase

import (
	"context"

	"github.com/KeiKom083/Portfolio-Backend/internal/domain/model"
	"github.com/KeiKom083/Portfolio-Backend/internal/domain/repository"
)

// UserUsecase はユーザーに関するビジネスロジックを担う。
type UserUsecase struct {
	userRepo repository.UserRepository
}

// NewUserUsecase は UserUsecase を生成する。
func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

// GetUser は ID でユーザーを取得する。
func (u *UserUsecase) GetUser(ctx context.Context, id string) (*model.User, error) {
	return u.userRepo.FindByID(ctx, id)
}

// ListUsers は全ユーザーを取得する。
func (u *UserUsecase) ListUsers(ctx context.Context) ([]*model.User, error) {
	return u.userRepo.FindAll(ctx)
}

// CreateUser は新しいユーザーを作成する。
func (u *UserUsecase) CreateUser(ctx context.Context, name, email string) (*model.User, error) {
	user := &model.User{
		Name:  name,
		Email: email,
	}
	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}
