package database

import (
	"go-restful-product-search-webapp/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var svc *dynamodb.DynamoDB
var sess *session.Session

// DB struct enables data access and allows mocking for tests.
type DB struct {
	Client dynamodbiface.DynamoDBAPI
}

func GetSession(configMap config.Config) *session.Session {
	if sess != nil {
		return sess
	}
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(configMap.AWS.Region),
		//Endpoint: aws.String("http://127.0.0.1:8000"),
		Endpoint: aws.String(configMap.DynamoDB.Endpoint),
	}))
	return sess
}

// GetClient provides a dynamodb client.
func GetClient(configMap config.Config) dynamodbiface.DynamoDBAPI {
	if svc != nil {
		return svc
	}
	svc = dynamodb.New(GetSession(configMap), &aws.Config{
		Endpoint: aws.String(configMap.DynamoDB.Endpoint),
	})
	return svc
}
