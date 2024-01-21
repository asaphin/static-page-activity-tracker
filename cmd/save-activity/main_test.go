package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type SaveActivityUsecaseMock struct {
	mock.Mock
}

func (m *SaveActivityUsecaseMock) SaveActivity(_ context.Context, data map[string]interface{}) error {
	args := m.Called(data)

	return args.Error(0)
}

func TestHandler(t *testing.T) {
	testCases := []struct {
		name          string
		request       *events.APIGatewayProxyRequest
		usecase       *SaveActivityUsecaseMock
		shouldBeError bool
	}{
		{
			name: "should call SaveActivity method with specified data and not return an error",
			request: &events.APIGatewayProxyRequest{
				Body: "{\"k1\": \"v1\", \"k2\": \"v2\"}",
			},
			usecase: func() *SaveActivityUsecaseMock {
				m := &SaveActivityUsecaseMock{}

				m.On("SaveActivity", map[string]interface{}{"k1": "v1", "k2": "v2"}).
					Return(nil).Once()

				return m
			}(),
			shouldBeError: false,
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			h := handler{usecase: c.usecase}

			_, err := h.handle(context.Background(), c.request)

			if c.shouldBeError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
