package mocks

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type MockedScanOutput struct {
	dynamodbiface.DynamoDBAPI
	Resp    dynamodb.ScanOutput
	GetResp dynamodb.GetItemOutput
}

func (m MockedScanOutput) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return &m.GetResp, nil
}

func (m MockedScanOutput) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	return &m.Resp, nil
}

type MockedSearchOutput struct {
	dynamodbiface.DynamoDBAPI
	Resp dynamodb.QueryOutput
}

func (m MockedSearchOutput) Search(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	return &m.Resp, nil
}
