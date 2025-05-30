package database

import (
	"context"
	"fmt"

	"github.com/yasseldg/go-simple/repos/rMongo/internal/collection"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/connection"

	"github.com/yasseldg/go-simple/repos/rIndex"

	"github.com/yasseldg/mgm/v4"

	"go.mongodb.org/mongo-driver/mongo"
)

type Base struct {
	database *mongo.Database

	collections collection.Map
}

func New(database *mongo.Database) *Base {
	return &Base{
		database:    database,
		collections: make(collection.Map),
	}
}

func (db *Base) Name() string {
	return db.database.Name()
}

func (db *Base) Collections() collection.Map {
	return db.collections
}

func (db *Base) GetCollection(ctx context.Context, coll_name string, conn *connection.Base, indexes ...rIndex.Inter) (*collection.Base, error) {
	coll := db.collections.Get(coll_name)
	if coll != nil {
		return coll, nil
	}

	return db.setCollection(ctx, coll_name, conn, indexes...)
}

// admin functions

func (db *Base) Drop(ctx context.Context) error {
	return db.database.Drop(ctx)
}

func (db *Base) RunCommand(ctx context.Context, command any) error {

	result := db.database.RunCommand(ctx, command)

	return result.Err()
}

func (db *Base) ListCollectionNames(ctx context.Context, filter any) ([]string, error) {
	return db.database.ListCollectionNames(ctx, filter)
}

func (db *Base) NewCollection(name string) *mongo.Collection {
	return db.database.Collection(name)
}

// private methods

func (db *Base) setCollection(ctx context.Context, coll_name string, conn *connection.Base, indexes ...rIndex.Inter) (*collection.Base, error) {

	coll := mgm.NewCollection(db.database, coll_name)
	if coll == nil {
		return nil, fmt.Errorf(" mgm.NewCollection( %s , %s ) is nil", db.database.Name(), coll_name)
	}

	base := collection.New(coll, db.database.Name(), conn)

	err := base.CreateIndexes(ctx, indexes...)
	if err != nil {
		return nil, err
	}

	db.collections.Set(base)

	return db.collections[coll.Name()], nil
}
