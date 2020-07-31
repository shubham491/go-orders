package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/jyotishp/go-orders/pkg/models"
)

func GetRestaurant(tableName string, id int32) (models.Restaurant, error) {
	type Input struct {
		Id int32
	}

	key, err := dynamodbattribute.MarshalMap(Input{Id: id})
	if err != nil {
		printError(err)
		return models.Restaurant{}, err
	}

	ip := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: key,
	}

	svc := createSession()

	res, err := svc.GetItem(ip)
	if err != nil {
		printError(err)
		return models.Restaurant{}, err
	}

	restaurant := models.Restaurant{}

	err = dynamodbattribute.UnmarshalMap(res.Item, &restaurant)
	if err != nil {
		printError(err)
		return models.Restaurant{}, err
	}

	return restaurant, nil
}