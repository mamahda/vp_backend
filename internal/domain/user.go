package domain

type User struct {
	ID       	int    `json:"id"`
	Email    	string `json:"email" binding:"required,email"`           // Berdasarkan standar RFC 5321 (jalur SMTP).
	Password 	string `json:"password" binding:"required,min=8,max=72"` // Minimum 8 untuk keamanan. Maksimum 72 karena algoritma Bcrypt hanya memproses hingga 72 karakter pertama.
	Username  string `json:"name" binding:"required,max=100"`          // Nama maksimal 100
	Phone 		string `json:"phone"`            // Format E.164 untuk nomor telepon internasional.
	Role_ID  	string `json:"role_id"`
}
