package repository

import (
	"database/sql"
	"vp_backend/internal/domain"

	"github.com/go-sql-driver/mysql"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) Create(user *domain.User) error {
	_, err := r.DB.Exec(
		"INSERT INTO users (email, password, username, phone, role_id) VALUES (?, ?, ?, ?, ?)",
		user.Email, user.Password, user.Username, user.Phone, 3,
	)
	if err != nil {
		// Type assertion untuk menangkap error spesifik MySQL
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return domain.ErrEmailAlreadyExists // Kembalikan error dari domain
			}
		}
		return err
	}
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	row := r.DB.QueryRow(
		"SELECT id, email, password, username, role_id FROM users WHERE email = ?",
		email,
	)

	user := domain.User{}
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Username, &user.Role_ID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByID(id int) (*domain.User, error) {
	row := r.DB.QueryRow(
		"SELECT id, email, password, username, role_id FROM users WHERE id = ?",
		id,
	)

	user := domain.User{}
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Username, &user.Role_ID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
