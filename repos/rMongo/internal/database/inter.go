package database

import (
	"context"

	"github.com/yasseldg/go-simple/repos/rIndex"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/collection"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/connection"

	"go.mongodb.org/mongo-driver/mongo"
)

type Inter interface {
	Name() string
	Collections() collection.Map
	GetCollection(ctx context.Context, coll_name string, conn *connection.Base, indexes ...rIndex.Inter) (*collection.Base, error)
}

type InterAdmin interface {
	Drop(ctx context.Context) error
	ListCollectionNames(ctx context.Context, filter interface{}) ([]string, error)
	NewCollection(name string) *mongo.Collection
}
