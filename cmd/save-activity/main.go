package main

import (
	"context"
	"encoding/json"
	"github.com/asaphin/static-page-activity-tracker/app"
	"github.com/asaphin/static-page-activity-tracker/infrastructure/repository"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type handler struct {
	usecase app.SaveActivityUsecase
}

func (h *handler) handle(ctx context.Context, request events.APIGatewayProxyRequest) error {
	body := []byte(request.Body)

	var data map[string]interface{}

	err := json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	return h.usecase.SaveActivity(ctx, data)
}

func main() {
	repo := repository.NewDynamoDBActivityRepository()

	h := handler{usecase: app.NewSaveActivityUsecase(repo)}

	lambda.Start(h.handle)
}
