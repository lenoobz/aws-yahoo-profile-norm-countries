package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	logger "github.com/lenoobz/aws-lambda-logger"
	"github.com/lenoobz/aws-yahoo-profile-norm-countries/config"
	"github.com/lenoobz/aws-yahoo-profile-norm-countries/infrastructure/repositories/mongodb/repos"
	"github.com/lenoobz/aws-yahoo-profile-norm-countries/usecase/breakdown"
	"github.com/lenoobz/aws-yahoo-profile-norm-countries/usecase/yahoo"
)

type YahooProfileCountryRequest struct {
	Tickers []string `json:"tickers"`
}

func main() {
	lambda.Start(lambdaHandler)
}

func lambdaHandler(ctx context.Context, req YahooProfileCountryRequest) {
	log.Println("lambda handler is called")

	appConf := config.AppConf

	// create new logger
	zap, err := logger.NewZapLogger()
	if err != nil {
		log.Fatal("create app logger failed")
	}
	defer zap.Close()

	// create new repository
	yahooProfileRepo, err := repos.NewYahooProfileMongo(nil, zap, &appConf.Mongo)
	if err != nil {
		log.Fatalf("create Yahoo profile mongo failed: %v\n", err)
	}
	defer yahooProfileRepo.Close()

	// create new repository
	countryBreakdownRepo, err := repos.NewCountryBreakdownMongo(nil, zap, &appConf.Mongo)
	if err != nil {
		log.Fatal("create country breakdown mongo failed")
	}
	defer countryBreakdownRepo.Close()

	// create new service
	yahooService := yahoo.NewService(yahooProfileRepo, zap)
	breakdownService := breakdown.NewService(countryBreakdownRepo, *yahooService, zap)

	// try correlation context
	if err := breakdownService.AddAssetCountryBreakdownByTickers(ctx, req.Tickers); err != nil {
		log.Fatal("add asset country breakdown failed")
	}
}
