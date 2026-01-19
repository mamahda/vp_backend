package repository

import (
	"context"
	"database/sql"

	"vp_backend/internal/domain"

	"github.com/go-sql-driver/mysql"
)

// UserRepository bertanggung jawab untuk
// seluruh operasi database terkait user.
type UserRepository struct {
	DB *sql.DB
}

// Create menyimpan user baru ke database.
//
// Secara default role_id diset ke 3 (user biasa).
//
// Return:
// - domain.ErrEmailAlreadyExists jika email duplikat
// - error lain jika query gagal
func (r *UserRepository) Create(
	ctx context.Context,
	user *domain.User,
) error {

	query := `
		INSERT INTO users (email, password, username, phone_number, role_id)
		VALUES (?, ?, ?, ?, ?)
	`

	_, err := r.DB.ExecContext(
		ctx,
		query,
		user.Email,
		user.Password,
		user.Username,
		user.Phone,
		3, // default role user
	)

	if err != nil {
		// Handle duplicate email (MySQL error 1062)
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return domain.ErrEmailAlreadyExists
			}
		}
		return err
	}

	return nil
}

// FindByEmail mengambil data user
// berdasarkan email.
//
// Digunakan untuk proses login/authentication.
func (r *UserRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*domain.User, error) {

	query := `
		SELECT id, email, password, username, role_id
		FROM users
		WHERE email = ?
	`

	user := domain.User{}
	err := r.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Username,
		&user.Role_ID,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindByID mengambil data user
// berdasarkan ID.
//
// Biasanya digunakan untuk:
// - get profile
// - authorization
func (r *UserRepository) FindByID(
	ctx context.Context,
	id int,
) (*domain.User, error) {

	query := `
		SELECT id, email, password, username, role_id
		FROM users
		WHERE id = ?
	`

	user := domain.User{}
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Username,
		&user.Role_ID,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update memperbarui data user
// (username, email, dan phone number).
//
// Password tidak diubah di sini.
func (r *UserRepository) Update(
	ctx context.Context,
	user *domain.User,
) error {

	query := `
		UPDATE users
		SET username = ?, email = ?, phone_number = ?
		WHERE id = ?
	`

	_, err := r.DB.ExecContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.Phone,
		user.ID,
	)

	return err
}

