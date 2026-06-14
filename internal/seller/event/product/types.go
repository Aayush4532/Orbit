package product

type ProductItemRequest struct {
	Title       string  `json:"title" binding:"required,min=2,max=100"`
	Description string  `json:"description" binding:"required,min=10,max=500"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Frequency   int     `json:"frequency" binding:"required,gt=0"`
	Image       string  `json:"image" binding:"required,url"`
}

type RegisterProductsRequestBody struct {
	Products []ProductItemRequest `json:"products" binding:"required,dive"`
}

type UpdateProductRequestBody struct {
	Title       *string  `json:"title" binding:"omitempty,min=2,max=100"`
	Description *string  `json:"description" binding:"omitempty,min=10,max=500"`
	Price       *float64 `json:"price" binding:"omitempty,gt=0"`
	Frequency   *int     `json:"frequency" binding:"omitempty,gt=0"`
	Image       *string  `json:"image" binding:"omitempty,url"`
}