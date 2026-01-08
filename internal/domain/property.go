package domain

import "time"

type Property struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Price           int       `json:"price"`
	Status          int       `json:"status"`
	Province        string    `json:"province"`
	Regency         string    `json:"regency"`
	District        string    `json:"district"`
	Address         string    `json:"address"`
	Description     string    `json:"description"`
	BuildingArea    string    `json:"building_area"`
	LandArea        string    `json:"land_area"`
	Electricity     string    `json:"electricity"`
	WaterSource     string    `json:"water_source"`
	Bedrooms        string    `json:"bedrooms"`
	Bathrooms       string    `json:"bathrooms"`
	Floors          int       `json:"floors"`
	Garage          int       `json:"garage"`
	Carport         int       `json:"carport"`
	Certificate     string    `json:"certificate"`
	YearConstructed string    `json:"year_constructed"`
	CreatedAt       time.Time `json:"created_at"`
	PropertyTypeId  int       `json:"property_type_id"`
	UserId          int       `json:"user_id"`
}
