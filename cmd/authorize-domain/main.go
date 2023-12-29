package main

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/rs/zerolog"
	"os"
	"strings"
)

var log = zerolog.New(os.Stderr).Level(zerolog.DebugLevel).With().Timestamp().Logger()

var (
	awsRegion                  = os.Getenv("AWS_REGION")
	authorizedDomainsTableName = os.Getenv("AUTHORIZED_DOMAINS_TABLE")
)

var db = func() *dynamo.DB {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(awsRegion)}))

	return dynamo.New(sess)
}()

var (
	pk       = "AUTHORIZED_DOMAIN"
	skPrefix = "DOMAIN#"
)

type AuthorizedDomainRecord struct {
	PK string `json:"pk"`
	SK string `json:"sk"`
}

var authorizrdDomains = func() map[string]struct{} {
	var domains []*AuthorizedDomainRecord

	err := db.Table(authorizedDomainsTableName).Get("pk", pk).Range("sk", dynamo.BeginsWith, skPrefix).All(&domains)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to load authorized domains")
	}

	domainsMap := make(map[string]struct{})

	for _, domain := range domains {
		domainsMap[strings.TrimPrefix(domain.SK, skPrefix)] = struct{}{}
	}

	return domainsMap
}()

func handle(_ context.Context, request events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	log.Debug().Interface("request", request).Msg("got a request")

	if _, ok := authorizrdDomains[request.AuthorizationToken]; !ok {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	resp := events.APIGatewayCustomAuthorizerResponse{
		PrincipalID: request.AuthorizationToken,
		PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   "Allow",
					Resource: []string{"*"},
				},
			},
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(handle)
}
