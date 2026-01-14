package domain

type PropertyImage struct {
	ID          int    `json:"id"`
	Url         string `json:"url"`
	Property_ID int    `json:"property_id"`
}
