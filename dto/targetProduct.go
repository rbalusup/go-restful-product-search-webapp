package dto

// TargetProduct is a struct to hold information about products
type TargetProduct struct {
	TCIN  string  `json:"tcin" dynamodbav:"tcin"`
	Title string  `json:"title" dynamodbav:"title"`
	Price float64 `json:"price" dynamodbav:"price"`
}

// ProductResponse is an API struct to hold and return a product.
type ProductResponse struct {
	Data TargetProduct `json:"data"`
}

// ProductSearchResult is an API struct to hold and return search results.
type ProductSearchResult struct {
	Data []TargetProduct `json:"data"`
}
