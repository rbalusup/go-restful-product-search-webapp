package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/gin-gonic/gin"

	"go-restful-product-search-webapp/database"
	"go-restful-product-search-webapp/dto"
	"go-restful-product-search-webapp/mocks"
	"go-restful-product-search-webapp/server"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	response := parseSimpleJSONResponse(t, doGetRequest(t, "/"))

	value, exists := response["hello"]
	assert.True(t, exists)

	body := gin.H{
		"hello": "SHIPT",
	}
	assert.Equal(t, body["hello"], value)
}

func TestProductScan(t *testing.T) {
	response := doProductScan(t, "Yellow")
	product := response[0]
	assert.Equal(t, "76695887", product.TCIN)
	assert.Equal(t, "Yellow", product.Title)
	assert.Equal(t, 451.47, product.Price)
}

func doRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func doGetRequest(t *testing.T, path string) *httptest.ResponseRecorder {
	db := &database.DB{
		Client: mocks.MockedScanOutput{
			Resp: dynamodb.ScanOutput{
				Count: aws.Int64(1),
				Items: []map[string]*dynamodb.AttributeValue{
					{
						"TCIN": {
							S: aws.String("76695887"),
						},
						"Title": {
							S: aws.String("Yellow"),
						},
						"Price": {
							N: aws.String("451.47"),
						},
					},
				},
			},
			GetResp: dynamodb.GetItemOutput{
				Item: map[string]*dynamodb.AttributeValue{
					"TCIN": {
						S: aws.String("76695886"),
					},
					"Title": {
						S: aws.String("Red"),
					},
					"Price": {
						N: aws.String("457.57"),
					},
				},
			},
		},
	}
	r := server.SetupRouter(db)
	w := doRequest(r, "GET", path)
	assert.Equal(t, http.StatusOK, w.Code)
	return w
}

func parseSimpleJSONResponse(t *testing.T, w *httptest.ResponseRecorder) map[string]string {
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	return response
}

func doProductScan(t *testing.T, searchTerm string) []dto.TargetProduct {
	w := doGetRequest(t, "/api/v1/products/scan?q="+searchTerm)
	var response map[string][]dto.TargetProduct
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)

	value, exists := response["data"]
	assert.True(t, exists)
	return value
}

func doProductSearch(t *testing.T, searchTerm string) []dto.TargetProduct {
	w := doGetRequest(t, "/api/v1/products/search?q="+searchTerm)
	var response map[string][]dto.TargetProduct
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)

	value, exists := response["data"]
	assert.True(t, exists)
	return value
}
