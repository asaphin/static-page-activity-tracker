package app

import (
	"context"

	"github.com/asaphin/static-page-activity-tracker/domain"
)

type ActivityRepository interface {
	Save(ctx context.Context, activity *domain.Activity) error
}
