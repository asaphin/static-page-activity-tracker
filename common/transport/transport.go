package transport

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func getHeaders() map[string]string {
	return map[string]string{
		"Content-Type":                 "application/json",
		"Access-Control-Allow-Methods": "OPTIONS,GET,POST,PUT,DELETE",
		"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
		"Access-Control-Allow-Origin":  "*",
	}
}

func newResponse(statusCode int) *events.APIGatewayV2HTTPResponse {
	return &events.APIGatewayV2HTTPResponse{
		Headers:    getHeaders(),
		StatusCode: statusCode,
	}
}

func Send(statusCode int, body any) (*events.APIGatewayV2HTTPResponse, error) {
	resp := newResponse(statusCode)

	var stringBody string

	if s, ok := body.(string); ok {
		stringBody = s
	} else {
		jsonBody, _ := json.Marshal(body)
		stringBody = string(jsonBody)
	}

	resp.Body = stringBody

	return resp, nil
}

func SendError(err error) (*events.APIGatewayV2HTTPResponse, error) {
	errResponseBody := map[string]any{
		"errorMessage": err.Error(),
		"statusCode":   "InternalServerError",
	}

	resp := newResponse(http.StatusInternalServerError)

	errResponseJSON, _ := json.Marshal(errResponseBody)
	resp.Body = string(errResponseJSON)

	return resp, nil
}
