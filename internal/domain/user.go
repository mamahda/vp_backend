package domain

type User struct {
	ID       	int    `json:"id"`
	Email    	string `json:"email" binding:"required,email"`           
	Password 	string `json:"password" binding:"required,min=8,max=72"` 
	Username  string `json:"name" binding:"required,max=100"`         
	Phone 		string `json:"phone"`            
	Role_ID  	string `json:"role_id"`
}
