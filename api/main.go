package main

import (
	"context"

	"validator/routes"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

var chiLambdaV2 *chiadapter.ChiLambdaV2

func setupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "OPTIONS"},
		MaxAge:         300,
	}).Handler)

	r.Get("/", routes.HomeRoute)

	r.Get("/validate", routes.ValidateRoute)

	return r
}

func init() {
	r := setupRouter()
	chiLambdaV2 = chiadapter.NewV2(r)
	chiLambdaV2.StripBasePath("/prod")
}

func handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return chiLambdaV2.ProxyWithContextV2(ctx, req)
}

func main() {
	lambda.Start(handler)
}
