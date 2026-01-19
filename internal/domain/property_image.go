package domain

// PropertyImage merepresentasikan data gambar
// yang terkait dengan suatu properti.
type PropertyImage struct {
	ID          int    `json:"id"`
	Url         string `json:"url"`
	Property_ID int    `json:"property_id"`
}

