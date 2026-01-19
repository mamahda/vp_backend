package domain

import "time"

// Property merepresentasikan entitas utama properti
// yang dipublikasikan di dalam sistem Victoria Property.
//
// Digunakan untuk:
// - Response API
// - Mapping data dari database
// - Business logic di service layer
type Property struct {

	// Primary key
	ID int `json:"id"`

	// Informasi utama
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`

	// Status properti
	// Contoh:
	// 0 = Draft
	// 1 = Published
	// 2 = Sold / Rented
	Status int `json:"status"`

	// Lokasi
	Province string `json:"province"`
	Regency  string `json:"regency"`
	District string `json:"district"`
	Address  string `json:"address"`

	// Spesifikasi properti
	BuildingArea int `json:"building_area"` // m²
	LandArea     int `json:"land_area"`     // m²
	Electricity  int `json:"electricity"`   // VA
	WaterSource  int `json:"water_source"`  // enum / code
	Bedrooms     int `json:"bedrooms"`
	Bathrooms    int `json:"bathrooms"`
	Floors       int `json:"floors"`
	Garage       int `json:"garage"`
	Carport      int `json:"carport"`

	// Legalitas & detail tambahan
	Certificate     string `json:"certificate"`
	YearConstructed int    `json:"year_constructed"`

	// Jenis transaksi
	// jual / sewa
	SaleType string `json:"sale_type"`

	// Relasi
	PropertyTypeId int `json:"property_type_id"`
	UserId         int `json:"user_id"`

	// Metadata
	CreatedAt time.Time `json:"created_at"`
}

