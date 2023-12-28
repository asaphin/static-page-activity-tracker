package app

import (
	"context"
	"time"
)

type saveActivityUsecase struct {
	repository ActivityRepository
}

func NewSaveActivityUsecase(repository ActivityRepository) SaveActivityUsecase {
	return &saveActivityUsecase{
		repository: repository}
}

func (u *saveActivityUsecase) SaveActivity(ctx context.Context, data map[string]interface{}) error {
	data["timestamp"] = time.Now().UnixMilli()

	return u.repository.Save(ctx, data)
}
