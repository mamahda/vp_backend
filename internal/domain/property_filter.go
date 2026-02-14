package domain

// PropertyFilters merepresentasikan kumpulan parameter
// query untuk melakukan pencarian dan filtering properti.
//
// Struct ini digunakan untuk:
// - Binding query parameter dari request HTTP
// - Membentuk WHERE, ORDER BY, dan pagination di repository
type PropertyFilters struct {

	// Jenis transaksi properti (jual / sewa)
	SaleType string `form:"sale_type"`

	// ID jenis properti (rumah, apartemen, dll)
	PropertyTypeID int `form:"property_type_id"`

	// Lokasi properti
	Province string `form:"province"`
	Regency  string `form:"regency"`

	// Filter harga
	MinPrice int64 `form:"min_price"`
	MaxPrice int64 `form:"max_price"`

	// Filter luas bangunan (m²)
	MinBuildingArea int `form:"min_building_area"`
	MaxBuildingArea int `form:"max_building_area"`

	// Filter luas tanah (m²)
	MinLandArea int `form:"min_land_area"`
	MaxLandArea int `form:"max_land_area"`

	// Jenis sertifikat properti
	Certificate string `form:"certificate"`

	// Filter keyword judul/deksripsi
	Keyword string `form:"keyword"`

	// Opsi pengurutan data
	// Contoh:
	// - price_asc
	// - price_desc
	// - oldest
	SortBy string `form:"sort"`

	// Pagination
	Page  int `form:"page"`
	Limit int `form:"limit"`

	// Offset digunakan secara internal
	// (tidak diambil dari request)
	Offset int `form:"-"`
}
