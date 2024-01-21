package repository

import (
	"context"
	"github.com/rs/zerolog/log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/asaphin/static-page-activity-tracker/domain"
)

var (
	awsRegion         = os.Getenv("AWS_REGION")
	activityTableName = os.Getenv("ACTIVITY_TABLE")
)

type DynamoDBActivityRepository struct {
	client *dynamodb.Client
}

func NewDynamoDBActivityRepository() *DynamoDBActivityRepository {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(awsRegion))
	if err != nil {
		panic(err)
	}

	client := dynamodb.NewFromConfig(cfg)

	return &DynamoDBActivityRepository{
		client: client,
	}
}

func (r *DynamoDBActivityRepository) Save(ctx context.Context, activity *domain.Activity) error {
	item, err := attributevalue.MarshalMap(activity)
	if err != nil {
		return err
	}

	log.Debug().Interface("item", item).Msg("item marshaled for dynamo db")

	_, err = r.client.PutItem(ctx, &dynamodb.PutItemInput{TableName: aws.String(activityTableName), Item: item})

	return err
}
