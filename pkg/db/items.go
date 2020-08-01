package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/jyotishp/go-orders/pkg/models"
)

func GetItem(tableName string, restaurantId int32, itemId int32) (models.Item, error) {
	type Input struct {
		ItemId int32
		RestaurantId int32
	}

	key, err := dynamodbattribute.MarshalMap(Input{ItemId: itemId, RestaurantId: restaurantId})
	if err != nil {
		printError(err)
		return models.Item{}, err
	}

	ip := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: key,
	}
	svc := createSession()
	res, err := svc.GetItem(ip)
	if err != nil {
		printError(err)
		return models.Item{}, err
	}
	opItem := models.Item{}
	err = dynamodbattribute.UnmarshalMap(res.Item, &opItem)
	if err != nil {
		printError(err)
		return models.Item{}, err
	}
	return opItem, nil
}

func InsertItem(tableName string, restaurantId int32, item models.Item) (models.Item ,error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		printError(err)
		return models.Item{}, err
	}

	item.Id = int32(uid.ID())

	ip, err := dynamodbattribute.MarshalMap(createItem(restaurantId, item))
	if err != nil {
		printError(err)
		return models.Item{}, err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: ip,
	}

	svc := createSession()
	_, err = svc.PutItem(input)
	if err != nil {
		printError(err)
		return models.Item{}, err
	}
	return item, nil
}

func UpdateItem(tableName string, restaurantId int32, item models.Item) (models.Item, error) {
	type KeyInput struct {
		RestaurantId int32
		ItemId int32
	}

	key, err := dynamodbattribute.MarshalMap(KeyInput{ItemId: item.Id, RestaurantId: restaurantId})
	if err != nil {
		printError(err)
		return models.Item{}, err
	}

	imap, err := dynamodbattribute.MarshalMap(updateItemMap(createItem(restaurantId, item)))

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: key,
		ExpressionAttributeValues: imap,
		UpdateExpression: aws.String("set Cuisine=:ic, Amount=:iamt, Discount=:idc"),
	}

	svc := createSession()
	_, err = svc.UpdateItem(input)
	if err != nil {
		printError(err)
		return models.Item{}, err
	}

	return item, nil
}