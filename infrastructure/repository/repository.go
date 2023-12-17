package repository

import (
	"context"
	"github.com/asaphin/static-page-activity-tracker/app"
)

type DynamoDBActivityRepository struct {
}

func NewDynamoDBActivityRepository() app.ActivityRepository {
	return &DynamoDBActivityRepository{}
}

func (r *DynamoDBActivityRepository) Save(ctx context.Context, event map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
}
