package repository

import (
	"context"

	"github.com/KeiKom083/Portfolio-Backend/internal/domain/model"
)

// UserRepository はユーザーの永続化に関するインターフェース。
// ドメイン層で定義し、インフラ層で実装する（依存性逆転の原則）。
type UserRepository interface {
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindAll(ctx context.Context) ([]*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}
