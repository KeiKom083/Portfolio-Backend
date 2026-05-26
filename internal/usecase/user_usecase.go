package usecase

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/KeiKom083/Portfolio-Backend/internal/domain/model"
	"github.com/KeiKom083/Portfolio-Backend/internal/domain/repository"
)

var ErrInvalidCredentials = errors.New("invalid email or password")

// UserUsecase はユーザーに関するビジネスロジックを担う。
type UserUsecase struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
}

// NewUserUsecase は UserUsecase を生成する。
func NewUserUsecase(userRepo repository.UserRepository, sessionRepo repository.SessionRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo, sessionRepo: sessionRepo}
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

// SignUp はパスワードをハッシュ化してユーザーを新規登録する。
func (u *UserUsecase) SignUp(ctx context.Context, name, email, password string) (*model.User, error) {
	user, err := model.NewUser(name, email, password)
	if err != nil {
		return nil, err
	}
	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

// Login はメールアドレスとパスワードを検証し、セッションIDを返す。
func (u *UserUsecase) Login(ctx context.Context, email, password string) (*model.User, string, error) {
	user, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, "", ErrInvalidCredentials
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, "", ErrInvalidCredentials
	}
	sessionID, err := u.sessionRepo.Create(user.ID)
	if err != nil {
		return nil, "", fmt.Errorf("create session: %w", err)
	}
	return user, sessionID, nil
}

// GetUserBySessionID はセッションIDからユーザーを取得する。
func (u *UserUsecase) GetUserBySessionID(ctx context.Context, sessionID string) (*model.User, error) {
	userID, ok := u.sessionRepo.Get(sessionID)
	if !ok {
		return nil, errors.New("session not found")
	}
	return u.userRepo.FindByID(ctx, userID)
}

// Logout はセッションを削除する。
func (u *UserUsecase) Logout(sessionID string) {
	u.sessionRepo.Delete(sessionID)
}
