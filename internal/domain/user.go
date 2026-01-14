package domain

type User struct {
	ID       	int    `json:"id"`
	Email    	string `json:"email" binding:"required,email"`           
	Password 	string `json:"password" binding:"required,min=8,max=72"` 
	Username  string `json:"username" binding:"required,max=100"`         
	Phone 		string `json:"phone_number"`            
	Role_ID  	int `json:"role_id"`
}

func (u *User) IsAdmin() bool {
	return u.Role_ID == 1
}
