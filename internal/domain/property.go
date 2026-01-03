package domain

type Property struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Status			int			`json:"status"`
	Address	    string  `json:"location"`
	Price       float64 `json:"price"`
}
