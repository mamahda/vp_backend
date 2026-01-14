package domain

import "time"

type Property struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Price           int       `json:"price"`
	Status          int       `json:"status"`
	Province        string    `json:"province"`
	Regency         string    `json:"regency"`
	District        string    `json:"district"`
	Address         string    `json:"address"`
	BuildingArea    int    `json:"building_area"`
	LandArea        int		    `json:"land_area"`
	Electricity     int 	    `json:"electricity"`
	WaterSource     int		    `json:"water_source"`
	Bedrooms        int		    `json:"bedrooms"`
	Bathrooms       int		    `json:"bathrooms"`
	Floors          int       `json:"floors"`
	Garage          int       `json:"garage"`
	Carport         int       `json:"carport"`
	Certificate     string    `json:"certificate"`
	YearConstructed int	 		  `json:"year_constructed"`
	CreatedAt       time.Time `json:"created_at"`
	PropertyTypeId  int       `json:"property_type_id"`
	UserId          int       `json:"user_id"`
}
