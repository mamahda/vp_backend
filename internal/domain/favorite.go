package domain

type Favorite struct {
	ID          int `json:"id"`
	User_ID     int `json:"user_id"`
	Property_ID int `json:"property_id"`
}
