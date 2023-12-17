package app

import (
	"context"
)

type saveActivityUsecase struct {
	repository ActivityRepository
}

func NewSaveActivityUsecase(repository ActivityRepository) SaveActivityUsecase {
	return &saveActivityUsecase{
		repository: repository}
}

func (u *saveActivityUsecase) SaveActivity(ctx context.Context, data map[string]interface{}) error {
	return nil
}
