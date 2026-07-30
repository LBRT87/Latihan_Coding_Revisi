package usecase

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Flavour     string  `json:"flavour"`
	PictureUrl  string  `json:"picture_url"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
