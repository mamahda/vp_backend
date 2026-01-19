package domain

type PropertyFilters struct {
	SaleType        string `form:"sale_type" binding:"required"`
	PropertyTypeID  int    `form:"property_type_id"`
	Province        string `form:"province"`
	Regency         string `form:"regency"`
	MinPrice        int64  `form:"min_price"`
	MaxPrice        int64  `form:"max_price"`
	MinBuildingArea int    `form:"min_building_area"`
	MaxBuildingArea int    `form:"max_building_area"`
	MinLandArea     int    `form:"min_land_area"`
	MaxLandArea     int    `form:"max_land_area"`
	Certificate     string `form:"certificate"`
	SortBy          string `form:"sort"`
	Page            int    `form:"page" binding:"required"`
	Limit           int    `form:"limit" binding:"required"`
	Offset          int    `form:"-"`
}
