package resolver

import "github.com/KeiKom083/Portfolio-Backend/internal/usecase"

// Resolver はすべてのリゾルバのルート。
// ユースケースへの依存を保持する。
type Resolver struct {
	UserUsecase *usecase.UserUsecase
}

// NewResolver は Resolver を生成する。
func NewResolver(userUsecase *usecase.UserUsecase) *Resolver {
	return &Resolver{
		UserUsecase: userUsecase,
	}
}
