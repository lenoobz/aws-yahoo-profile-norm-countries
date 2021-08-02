package breakdown

import (
	"context"

	logger "github.com/lenoobz/aws-lambda-logger"
	"github.com/lenoobz/aws-yahoo-profile-norm-countries/usecase/yahoo"
)

// Service sector
type Service struct {
	breakdownRepo Repo
	yahooService  yahoo.Service
	log           logger.ContextLog
}

// NewService create new service
func NewService(breakdownRepo Repo, yahooService yahoo.Service, log logger.ContextLog) *Service {
	return &Service{
		breakdownRepo: breakdownRepo,
		yahooService:  yahooService,
		log:           log,
	}
}

// AddAssetCountryBreakdownByTickers adds new asset country breakdown for given tickers
func (s *Service) AddAssetCountryBreakdownByTickers(ctx context.Context, tickers []string) error {
	s.log.Info(ctx, "adding new asset country breakdown by tickers", "tickers", tickers)

	profiles, err := s.yahooService.FindYahooProfilesByTickers(ctx, tickers)
	if err != nil {
		s.log.Error(ctx, "find all Yahoo profiles by tickers failed", "error", err)
	}

	for _, profile := range profiles {
		countryBreakdown := profile.MapYahooProfileToAssetCountryBreakdown(ctx, s.log)

		if err := s.breakdownRepo.InsertAssetCountryBreakdown(ctx, countryBreakdown); err != nil {
			s.log.Error(ctx, "insert asset country breakdown failed", "error", err)
			return err
		}
	}

	return nil
}

// AddAssetCountryBreakdown adds new asset country breakdown
func (s *Service) AddAssetCountryBreakdown(ctx context.Context) error {
	s.log.Info(ctx, "adding new asset country breakdown")

	profiles, err := s.yahooService.FindYahooProfiles(ctx)
	if err != nil {
		s.log.Error(ctx, "find all Yahoo profiles failed", "error", err)
	}

	for _, profile := range profiles {
		countryBreakdown := profile.MapYahooProfileToAssetCountryBreakdown(ctx, s.log)

		if err := s.breakdownRepo.InsertAssetCountryBreakdown(ctx, countryBreakdown); err != nil {
			s.log.Error(ctx, "insert asset country breakdown failed", "error", err)
			return err
		}
	}

	return nil
}
