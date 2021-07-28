package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssetProfileModel struct {
	ID         *primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt  int64               `bson:"createdAt,omitempty"`
	ModifiedAt int64               `bson:"modifiedAt,omitempty"`
	Enabled    bool                `bson:"enabled"`
	Deleted    bool                `bson:"deleted"`
	Schema     string              `bson:"schema,omitempty"`
	Ticker     string              `bson:"ticker,omitempty"`
	Sector     string              `bson:"sector,omitempty"`
	Country    string              `bson:"country,omitempty"`
}
