package controller

import (
	"log"
	"net/http"

	"go-restful-product-search-webapp/dto"
	"go-restful-product-search-webapp/search"

	"github.com/gin-gonic/gin"
)

// SearchProducts godoc
// @Param q query string true "Search Term"
// @Success 200 {object} dto.ProductSearchResult
// @Router /products/search?q={q} [get]
func (c *Controller) SearchProducts(ctx *gin.Context) {
	searchQuery := ctx.Query("q")
	log.Printf("searchQuery is %s", searchQuery)
	ctx.JSON(http.StatusOK, &dto.ProductSearchResult{
		Data: search.ProductSearch(&c.DB, searchQuery),
	})
}

// ScanProducts godoc
// @Summary Scans for products using DynamoDB
// @Description Uses DynamoDB to find products by Title
// @Accept  json
// @Produce  json
// @Param q query string true "Search Term"
// @Success 200 {object} dto.ProductSearchResult
// @Router /products/scan?q={q} [get]
func (c *Controller) ScanProducts(ctx *gin.Context) {
	searchQuery := ctx.Query("q")
	ctx.JSON(http.StatusOK, &dto.ProductSearchResult{
		Data: search.ProductScan(&c.DB, searchQuery),
	})
}
