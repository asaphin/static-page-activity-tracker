package main

import (
	"context"
	"encoding/json"
	"github.com/asaphin/static-page-activity-tracker/app"
	"github.com/asaphin/static-page-activity-tracker/common/logging"
	"github.com/asaphin/static-page-activity-tracker/common/transport"
	"github.com/asaphin/static-page-activity-tracker/infrastructure/repository"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog/log"
)

func init() {
	logging.Setup()
}

type handler struct {
	usecase app.SaveActivityUsecase
}

func (h *handler) handle(ctx context.Context, request *events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Debug().Interface("request", request).Msg("got a request")

	body := []byte(request.Body)

	var data map[string]interface{}

	err := json.Unmarshal(body, &data)
	if err != nil {
		return transport.SendError(err)
	}

	log.Debug().Interface("data", data).Msg("request data unmarshalled")

	err = h.usecase.SaveActivity(ctx, data)
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
