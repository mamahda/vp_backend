package domain

// PropertyImage merepresentasikan data gambar
// yang terkait dengan suatu properti.
type PropertyImage struct {
	ID         int    `json:"id"`
	Url        string `json:"url"`
	PropertyID int    `json:"property_id"`
}
