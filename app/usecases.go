package app

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/asaphin/static-page-activity-tracker/domain"
)

type SaveActivityUsecase struct {
	repository ActivityRepository
}

func NewSaveActivityUsecase(repository ActivityRepository) *SaveActivityUsecase {
	return &SaveActivityUsecase{
		repository: repository}
}

func (u *SaveActivityUsecase) SaveActivity(ctx context.Context, activity *domain.Activity) error {
	err := u.repository.Save(ctx, activity)
	if err != nil {
		log.Error().Err(err).Msg("unable to save activity to repository")
	}

	return err
}
