package domain

import "errors"

// Domain-level errors digunakan untuk merepresentasikan error bisnis (business logic)
// yang bersifat umum dan reusable di seluruh layer aplikasi (service, handler, dll).
//
// Keuntungan menggunakan domain error:
// - Konsisten di seluruh aplikasi
// - Mudah dipetakan ke HTTP status code
// - Tidak mengekspos error internal (DB, library, dsb)

var (
	// ErrEmailAlreadyExists dikembalikan ketika user mencoba
	// mendaftar menggunakan email yang sudah terdaftar di sistem.
	//
	// Biasanya dipetakan ke:
	// HTTP 409 Conflict
	ErrEmailAlreadyExists = errors.New("email already registered")

	// ErrUserNotFound dikembalikan ketika data user tidak ditemukan,
	// misalnya saat login atau fetch user by ID.
	//
	// Biasanya dipetakan ke:
	// HTTP 404 Not Found atau 401 Unauthorized (tergantung konteks)
	ErrUserNotFound = errors.New("user not found")

	// ErrInvalidCredentials dikembalikan ketika email atau password
	// yang diberikan tidak sesuai.
	//
	// Digunakan saat proses login.
	// Biasanya dipetakan ke:
	// HTTP 401 Unauthorized
	ErrInvalidCredentials = errors.New("invalid email or password")

	// ErrInternalServer digunakan sebagai error fallback
	// ketika terjadi kesalahan internal yang tidak boleh
	// diekspos ke client (misalnya error signing JWT, DB error, dll).
	//
	// Biasanya dipetakan ke:
	// HTTP 500 Internal Server Error
	ErrInternalServer = errors.New("internal server error")
)

