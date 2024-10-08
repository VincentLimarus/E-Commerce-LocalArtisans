package requestsDTO

import "github.com/google/uuid"

type GetAllProductRequestDTO struct {
	Page      int    `json:"page" form:"page" binding:"omitempty"`
	Limit     int    `json:"limit" form:"limit" binding:"omitempty"`
	OrderBy   string `json:"order_by" form:"order_by" binding:"omitempty"`
	OrderType string `json:"order_type" form:"order_type" binding:"omitempty"`
}

type GetProductRequestDTO struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type GetAllProductByArtisanIDRequestDTO struct {
	ArtisanID uuid.UUID `json:"artisan_id" form:"artisan_id" binding:"required"`
}

type GetAllProductByCategoryIDRequestDTO struct {
	CategoryID uuid.UUID `json:"category_id" form:"category_id" binding:"required"`
}

type CreateProductRequestDTO struct {
	CategoryID  uuid.UUID   `json:"category_id" form:"category_id" binding:"required"`
	ArtisanID   uuid.UUID   `json:"artisan_id" form:"artisan_id" binding:"required"`
	Name        string  	`json:"name" form:"name" binding:"required"`
	Price       float64 	`json:"price" form:"price" binding:"required"`
	Description string  	`json:"description" form:"description" binding:"required"`
	Quantity    int     	`json:"quantity" form:"quantity" binding:"required"`
	CreatedBy   string  	`json:"created_by" form:"created_by" binding:"omitempty"`
	IsActive    bool    	`json:"is_active" form:"is_active" binding:"omitempty"`
}

type UpdateProductRequestDTO struct {
	ID          string    `json:"id" form:"id" binding:"required"`
	CategoryID  uuid.UUID `json:"category_id" form:"category_id" binding:"omitempty"`
	Name        string    `json:"name" form:"name" binding:"omitempty"`
	Price       float64   `json:"price" form:"price" binding:"omitempty"`
	Description string    `json:"description" form:"description" binding:"omitempty"`
	Quantity    int       `json:"quantity" form:"quantity" binding:"omitempty"`	
	UpdatedBy   string    `json:"updated_by" form:"updated_by" binding:"omitempty"`
	IsActive    bool      `json:"is_active" form:"is_active" binding:"omitempty"`
}

type DeleteProductRequestDTO struct {
	// Only admin can delete (Artisan).
	ID string `json:"id" form:"id" binding:"required"`
}

type AddProductToCartRequestDTO struct {
	ID 	 			string `json:"id" form:"id" binding:"required"`
	CartID 			string `json:"cart_id" form:"cart_id" binding:"omitempty"`
	Quantity 		int `json:"quantity" form:"quantity" binding:"required"` // Order Quantity
	IsActive 		bool `json:"is_active" form:"is_active" binding:"required"`
	CreatedBy 		string `json:"created_by" form:"created_by" binding:"omitempty"`
}

type CheckOutProductRequestDTO struct {
	ID 	 			string `json:"id" form:"id" binding:"required"`
	Quantity 		int `json:"quantity" form:"quantity" binding:"required"` // Order Quantity
	IsActive 		bool `json:"is_active" form:"is_active" binding:"required"`
	CreatedBy 		string `json:"created_by" form:"created_by" binding:"omitempty"`
}	