package app

import "context"

type SaveActivityUsecase interface {
	SaveActivity(ctx context.Context, data map[string]interface{}) error
}

type ActivityRepository interface {
	Save(ctx context.Context, event map[string]interface{}) error
}
