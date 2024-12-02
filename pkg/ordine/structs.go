package ordine

type ProductBody struct {
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
}
