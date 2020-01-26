package data

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"go-restful-product-search-webapp/config"
	"go-restful-product-search-webapp/database"
	"go-restful-product-search-webapp/dto"
	"go-restful-product-search-webapp/errors"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// dataSetup creates the products table, populates it with an initial data set, and makes sure a product is present.
func Setup(svc *database.DB, configMap config.Config) {
	result, _ := svc.Client.ListTables(&dynamodb.ListTablesInput{})
	tableNames := result.TableNames

	waitCounter := 0
	for !contains(tableNames, configMap.DynamoDB.ProductsTableName) {
		createTable(svc, configMap)
		result, _ := svc.Client.ListTables(&dynamodb.ListTablesInput{})
		tableNames = result.TableNames
		time.Sleep(2 * time.Second)
		waitCounter++
		if waitCounter >= 10 {
			break
		}
	}

	populateTable(svc, configMap)
	verifyData(svc, configMap)
}

// createTable creates the configured products table.
func createTable(svc *database.DB, configMap config.Config) {

	createTableInput := &dynamodb.CreateTableInput{
		TableName: aws.String(configMap.DynamoDB.ProductsTableName),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("tcin"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("tcin"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	_, err := svc.Client.CreateTable(createTableInput)
	errors.HandleIfError(err)

	fmt.Println("Created Table " + configMap.DynamoDB.ProductsTableName)
}

// populateTable puts a small set of initial documents in the products table.
func populateTable(svc *database.DB, configMap config.Config) {

	csvFile, _ := os.Open(configMap.FileContextPath.ProductsFile)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var products []dto.TargetProduct
	for {
		line, error2 := reader.Read()
		if error2 == io.EOF {
			break
		} else if error2 != nil {
			log.Fatal(error2)
		}

		price, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			price = 1.0
		}
		products = append(products, dto.TargetProduct{
			TCIN:  line[0],
			Title: line[1],
			Price: price,
		})
	}

	for _, product := range products {
		av, err := dynamodbattribute.MarshalMap(product)
		errors.HandleIfError(err)

		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String(configMap.DynamoDB.ProductsTableName),
		}

		_, err = svc.Client.PutItem(input)
		errors.HandleIfError(err)
		fmt.Println("Put product: " + product.Title)
	}
}

// verifyData makes sure the first product is present in the table.
func verifyData(svc *database.DB, configMap config.Config) {
	getItemResult, err := svc.Client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(configMap.DynamoDB.ProductsTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"tcin": {
				S: aws.String("76695884"),
			},
		},
	})
	errors.HandleIfError(err)

	retrievedProduct := dto.TargetProduct{}

	err = dynamodbattribute.UnmarshalMap(getItemResult.Item, &retrievedProduct)
	errors.HandleIfError(err)

	if retrievedProduct.Title == "" {
		fmt.Println("Could not find Product 76695884")
		return
	}
}

// contains returns true if the name is found in aRange and returns false otherwise.
func contains(aRange []*string, name string) bool {
	valueFound := false

	for _, n := range aRange {
		if name == *n {
			valueFound = true
			break
		}
	}
	return valueFound
}
