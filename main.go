package main

import (
	"go-restful-product-search-webapp/config"
	"go-restful-product-search-webapp/data"
	"go-restful-product-search-webapp/database"
	"go-restful-product-search-webapp/server"
	"log"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func main() {

	configMap := getConfigMap()
	productsFilePath := configMap.FileContextPath.ProductsFile
	isDataSetupNeeded := configMap.DynamoDB.EnableDataSetup
	log.Printf("Product File Path is %s", productsFilePath)
	log.Printf("Is Data setup needed %t", isDataSetupNeeded)

	db := &database.DB{
		Client: database.GetClient(configMap),
	}
	if isDataSetupNeeded {
		data.Setup(db, configMap)
	}

	r := server.SetupRouter(db)

	r.Run()
}

func getConfigMap() config.Config {
	configMap, err := config.GetConfigMap()
	if err != nil {
		log.Fatal(err)
	}
	return configMap
}
