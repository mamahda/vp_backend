package domain

// PropertyType merepresentasikan jenis properti
// dalam sistem (misalnya: rumah, apartemen, kontrakan).
type PropertyType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

