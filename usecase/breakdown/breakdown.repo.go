package breakdown

import (
	"context"

	"github.com/lenoobz/aws-yahoo-profile-norm-countries/entities"
)

// Reader interface
type Reader interface{}

// Writer interface
type Writer interface {
	InsertAssetCountryBreakdown(ctx context.Context, breakdown *entities.FundBreakdown) error
}

// Repo interface
type Repo interface {
	Reader
	Writer
}
