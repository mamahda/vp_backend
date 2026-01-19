package domain

// User merepresentasikan entitas user
// yang digunakan di seluruh layer aplikasi
// (handler, service, repository).
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=72"`
	Username string `json:"username" binding:"required,max=100"`
	Phone    string `json:"phone_number"`
	Role_ID  int    `json:"role_id"`
}

// IsAdmin menentukan apakah user
// memiliki role sebagai admin/agent.
//
// Return:
// - true  → jika Role_ID == 1
// - false → jika bukan admin
func (u *User) IsAdmin() bool {
	return u.Role_ID == 1
}

