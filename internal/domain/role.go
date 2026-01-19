package domain

// Role merepresentasikan peran (role)
// yang dimiliki oleh user dalam sistem.
type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

