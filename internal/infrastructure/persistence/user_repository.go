package persistence

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/KeiKom083/Portfolio-Backend/internal/domain/model"
	"github.com/KeiKom083/Portfolio-Backend/internal/domain/repository"
)

// userRepository は repository.UserRepository の PostgreSQL 実装。
type userRepository struct {
	pool *pgxpool.Pool
}

// NewUserRepository は UserRepository を生成する。
func NewUserRepository(pool *pgxpool.Pool) repository.UserRepository {
	return &userRepository{pool: pool}
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	user := &model.User{}
	err := r.pool.QueryRow(ctx,
		"SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1", id,
	).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("find user by id: %w", err)
	}
	return user, nil
}

func (r *userRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	rows, err := r.pool.Query(ctx,
		"SELECT id, name, email, created_at, updated_at FROM users ORDER BY created_at DESC",
	)
	if err != nil {
		return nil, fmt.Errorf("find all users: %w", err)
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		user := &model.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan user: %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	now := time.Now()
	err := r.pool.QueryRow(ctx,
		"INSERT INTO users (name, email, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Name, user.Email, now, now,
	).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}
	user.CreatedAt = now
	user.UpdatedAt = now
	return nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	now := time.Now()
	_, err := r.pool.Exec(ctx,
		"UPDATE users SET name = $1, email = $2, updated_at = $3 WHERE id = $4",
		user.Name, user.Email, now, user.ID,
	)
	if err != nil {
		return fmt.Errorf("update user: %w", err)
	}
	user.UpdatedAt = now
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	_, err := r.pool.Exec(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}
	return nil
}
