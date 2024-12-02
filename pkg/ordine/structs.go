package ordine

type Product struct {
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
}

type Ordine struct {
	Table      int       `json:"table"`
	ClientName string    `json:"client_name"`
	Status     bool      `json:"status"`
	Products   []Product `json:"products"`
}
