package domain

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required,min=3,max=20"` // Username min 3, max 20, wajib diisi
	Email    string `json:"email" binding:"required,email"`           // Berdasarkan standar RFC 5321 (jalur SMTP).
	Password string `json:"password" binding:"required,min=8,max=72"` // Minimum 8 untuk keamanan. Maksimum 72 karena algoritma Bcrypt hanya memproses hingga 72 karakter pertama.
	Name     string `json:"name" binding:"required,max=100"`          // Nama maksimal 100
	Role_ID  string `json:"role_id"`
}
