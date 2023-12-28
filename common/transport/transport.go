package transport

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

func getHeaders(method string) map[string]string {
	return map[string]string{
		"Content-Type":                 "application/json",
		"Access-Control-Allow-Headers": "Content-Type",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": method,
	}
}

func newResponse(statusCode int, method string) *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		Headers:    getHeaders(method),
		StatusCode: statusCode,
	}
}

func Send(statusCode int, method string, body any) (*events.APIGatewayProxyResponse, error) {
	resp := newResponse(statusCode, method)

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

func SendError(httpMethod string, err error) (*events.APIGatewayProxyResponse, error) {
	errResponseBody := map[string]any{
		"errorMessage": err.Error(),
		"errorCode":    "InternalServerError",
	}

	resp := newResponse(http.StatusInternalServerError, httpMethod)

	errResponseJSON, _ := json.Marshal(errResponseBody)
	resp.Body = string(errResponseJSON)

	return resp, nil
}
