package domain

type Property struct {
	ID          				int     `json:"id"`
	Title       				string  `json:"title"`
	Address	    				string  `json:"location"`
	Province 						string	`json:"province"`
	City 								string	`json:"city"`
	District 						string	`json:"district"`
	SubDistrict 				string	`json:"sub_district"`
	Description			 		string  `json:"description"`
	Bedroom							string	`json:"bedroom"`
	Bathroom						string	`json:"bathroom"`
	Land_Area						string	`json:"land_area"`
	Building_Area				string	`json:"building_area"`
	Electricity					string	`json:"electricity"`
	Water_Source				string	`json:"water_source"`
	Certificate					string	`json:"certificate"`
	Year_Constructed		string	`json:"year_constructed"`
	Price       				int `json:"price"`
	Status							int			`json:"status"`
}
