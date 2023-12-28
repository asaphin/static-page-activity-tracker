package repository

import (
	"context"
	"github.com/asaphin/static-page-activity-tracker/app"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"os"
)

var (
	awsRegion         = os.Getenv("AWS_REGION")
	activityTableName = os.Getenv("ACTIVITY_TABLE")
)

type DynamoDBActivityRepository struct {
	db *dynamo.DB
}

func NewDynamoDBActivityRepository() app.ActivityRepository {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(awsRegion)}))

	return &DynamoDBActivityRepository{
		db: dynamo.New(sess),
	}
}

func (r *DynamoDBActivityRepository) Save(ctx context.Context, event map[string]interface{}) error {
	return r.db.Table(activityTableName).Put(event).RunWithContext(ctx)
}
