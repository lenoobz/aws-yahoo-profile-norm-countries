package entities

import (
	"context"
	"fmt"
	"strings"

	logger "github.com/lenoobz/aws-lambda-logger"
	"github.com/lenoobz/aws-yahoo-profile-norm-countries/consts"
)

// YahooProfile struct
type YahooProfile struct {
	Ticker  string `json:"ticker,omitempty"`
	Country string `json:"country,omitempty"`
}

// MapYahooProfileToAssetCountryBreakdown map Yahoo profile to asset country breakdown
func (f *YahooProfile) MapYahooProfileToAssetCountryBreakdown(ctx context.Context, log logger.ContextLog) *FundBreakdown {
	countryCode, err := getCountryCode(f.Country)
	if err != nil {
		log.Warn(ctx, "get country code failed", "error", err, "CountryName", f.Country)
	}

	countryBreakdown := &CountryBreakdown{
		CountryCode:    countryCode,
		CountryName:    f.Country,
		FundMktPercent: 100,
		FundTnaPercent: 100,
	}

	var countriesBreakdown []*CountryBreakdown
	countriesBreakdown = append(countriesBreakdown, countryBreakdown)

	fundBreakdown := &FundBreakdown{
		Ticker:     f.Ticker,
		AssetClass: consts.EQUITY,
		Countries:  countriesBreakdown,
	}

	return fundBreakdown
}

// getCountryCode gets country code of a given name
func getCountryCode(name string) (string, error) {
	for _, country := range consts.Countries {
		if strings.EqualFold(strings.ToUpper(country.Name), strings.ToUpper(name)) {
			return country.Alpha3Code, nil
		}
	}
	return "OTH", fmt.Errorf("cannot find country code for country %s", name)
}
