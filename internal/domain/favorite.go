package domain

// Favorite merepresentasikan properti yang disimpan (wishlist)
// oleh user.
//
// Relasi:
// - User 1..* Favorite
// - Property 1..* Favorite
type Favorite struct {

	// Primary key
	ID int `json:"id"`

	// Relasi
	UserID     int `json:"user_id"`
	PropertyID int `json:"property_id"`
}

