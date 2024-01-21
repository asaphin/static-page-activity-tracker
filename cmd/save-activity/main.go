package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog/log"

	"github.com/asaphin/static-page-activity-tracker/app"
	"github.com/asaphin/static-page-activity-tracker/common/logging"
	"github.com/asaphin/static-page-activity-tracker/common/transport"
	"github.com/asaphin/static-page-activity-tracker/domain"
	"github.com/asaphin/static-page-activity-tracker/infrastructure/repository"
)

func init() {
	logging.Setup()
}

type ActivityDTO struct {
	Page         string                 `json:"page" dynamo:"page"`
	ActivityType string                 `json:"activityType" dynamo:"activityType"`
	Data         map[string]interface{} `json:"data" dynamo:"data"`
}

type handler struct {
	usecase *app.SaveActivityUsecase
}

func (h *handler) handle(ctx context.Context, request *events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	log.Debug().Interface("request", request).Msg("got a request")

	body := []byte(request.Body)

	var activity ActivityDTO

	err := json.Unmarshal(body, &activity)
	if err != nil {
		return transport.SendError(err)
	}

	log.Debug().Interface("data", activity).Msg("request data unmarshalled")

	err = h.usecase.SaveActivity(ctx, &domain.Activity{
		Page:         activity.Page,
		Timestamp:    time.Now().UnixMilli(),
		ActivityType: activity.ActivityType,
		IpAddress:    request.Headers["X-Forwarded-For"],
		UserAgent:    request.Headers["User-Agent"],
		Data:         activity.Data,
	})
	if err != nil {
		return transport.SendError(err)
	}

	return transport.Send(http.StatusCreated, "OK")
}

func main() {
	repo := repository.NewDynamoDBActivityRepository()

	h := handler{usecase: app.NewSaveActivityUsecase(repo)}

	lambda.Start(h.handle)
}
