package search

import (
	"go-restful-product-search-webapp/config"
	"go-restful-product-search-webapp/database"
	"go-restful-product-search-webapp/dto"
	"go-restful-product-search-webapp/errors"
	"log"
	"sort"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

	"github.com/aws/aws-sdk-go/aws"
)

// ProductScan scans product documents for matches.
func ProductScan(svc *database.DB, searchTerm string) []dto.TargetProduct {

	filt := expression.Name("title").Contains(searchTerm)

	proj := expression.NamesList(
		expression.Name("tcin"),
		expression.Name("title"),
		expression.Name("price"),
	)

	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	errors.HandleIfError(err)

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(getConfigMap().DynamoDB.ProductsTableName),
	}

	result, err := svc.Client.Scan(params)
	errors.HandleIfError(err)

	products := []dto.TargetProduct{}
	for _, i := range result.Items {
		item := dto.TargetProduct{}
		err = dynamodbattribute.UnmarshalMap(i, &item)
		errors.HandleIfError(err)

		products = append(products, item)
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].Price > products[j].Price
	})

	return products
}

// ProductSearch searches products for matches.
func ProductSearch(db *database.DB, searchTerm string) []dto.TargetProduct {
	params := GetQueryInput(searchTerm)
	result, err := db.Client.Query(params)
	errors.HandleIfError(err)

	var products []dto.TargetProduct
	for _, i := range result.Items {
		item := dto.TargetProduct{}
		err = dynamodbattribute.UnmarshalMap(i, &item)
		errors.HandleIfError(err)

		products = append(products, item)
	}
	return products
}

func GetQueryInput(searchTerm string) *dynamodb.QueryInput {
	params := &dynamodb.QueryInput{
		TableName:              aws.String(getConfigMap().DynamoDB.ProductsTableName),
		KeyConditionExpression: aws.String("tcin = :tcin"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":tcin": {
				S: aws.String(searchTerm),
			},
		},
		ScanIndexForward: aws.Bool(false),
	}
	return params
}

func getConfigMap() config.Config {
	configMap, err := config.GetConfigMap()
	if err != nil {
		log.Fatal(err)
	}
	return configMap
}
