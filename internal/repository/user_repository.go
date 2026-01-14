package repository

import (
	"context"
	"database/sql"
	"vp_backend/internal/domain"

	"github.com/go-sql-driver/mysql"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (email, password, username, phone_number, role_id) VALUES (?, ?, ?, ?, ?)`

	_, err := r.DB.ExecContext(ctx, query, user.Email, user.Password, user.Username, user.Phone, 3)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return domain.ErrEmailAlreadyExists // Kembalikan error dari domain
			}
		}
		return err
	}
	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, email, password, username, role_id FROM users WHERE email = ?`

	user := domain.User{}
	err := r.DB.QueryRowContext(ctx, query,	email).Scan(&user.ID, &user.Email, &user.Password, &user.Username, &user.Role_ID)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id int) (*domain.User, error) {
	query := `SELECT id, email, password, username, role_id FROM users WHERE id = ?`

	user := domain.User{}
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Email, &user.Password, &user.Username, &user.Role_ID)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
