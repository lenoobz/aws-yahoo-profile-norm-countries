package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	corid "github.com/lenoobz/aws-lambda-corid"
	logger "github.com/lenoobz/aws-lambda-logger"
	"github.com/lenoobz/aws-yahoo-profile-norm-countries/config"
	"github.com/lenoobz/aws-yahoo-profile-norm-countries/infrastructure/repositories/mongodb/repos"
	"github.com/lenoobz/aws-yahoo-profile-norm-countries/usecase/breakdown"
	"github.com/lenoobz/aws-yahoo-profile-norm-countries/usecase/yahoo"
)

func main() {
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
	id, _ := uuid.NewRandom()
	ctx := corid.NewContext(context.Background(), id)

	// try correlation context
	if err := breakdownService.AddAssetCountryBreakdown(ctx); err != nil {
		log.Fatal("add asset country breakdown failed")
	}
}
